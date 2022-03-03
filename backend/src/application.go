package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MastoCred-Inc/web-app/controller"
	"github.com/MastoCred-Inc/web-app/database/postgres_db"
	"github.com/MastoCred-Inc/web-app/h"
	"github.com/MastoCred-Inc/web-app/middleware"
	"github.com/MastoCred-Inc/web-app/utility/environment"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	applicationLogger := logger.With().Str("server", "app").Logger()

	r := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig), gin.Recovery())
	r.Use(ginzerolog.Logger("api"))

	env, err := environment.New()
	if err != nil {
		applicationLogger.Fatal().Err(err)
	}

	postgresDB := postgres_db.New(logger, env)
	defer postgresDB.Close()

	r.Use(GinContextToContextMiddleware())
	newMiddleware, err := middleware.NewMiddleware(logger, *env, postgresDB)
	if err != nil {
		applicationLogger.Fatal().Msgf("middleware error: %v", err)
	}

	controller := controller.New(logger, postgresDB, newMiddleware)

	r.Any("/api", h.GraphqlHandler(logger, *controller)) // grpc endpoint handler
	r.GET("/graphql-ui", h.PlaygroundHandler())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"rest":    false,
			"graphql": true,
		})
	})

	srv := &http.Server{
		Addr:    ":9004",
		Handler: r,
	}

	//go func() {
	// 	// service connection
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		applicationLogger.Fatal().Msgf("listen: %s", err)
	}
	//}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		applicationLogger.Fatal().Msgf("Server Shutdown: %v", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		applicationLogger.Info().Msgf("timeout of 5 seconds.")
	default:
	}

}

// GinContextToContextMiddleware middleware for gin context
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "ctxkey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
