package h

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gitlab.com/mastocred/web-app/controller"
	"gitlab.com/mastocred/web-app/h/graph"
	"gitlab.com/mastocred/web-app/h/graph/generated"
	"gitlab.com/mastocred/web-app/utility/environment"
)

func GraphqlHandler(l zerolog.Logger, c controller.Operations, env *environment.Env) gin.HandlerFunc {
	resolver := graph.New(l, c, env)
	graphqlResolver := generated.Config{Resolvers: resolver}

	h := handler.NewDefaultServer(generated.NewExecutableSchema(graphqlResolver))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}

}

// PlaygroundHandler spin up the playground /graphql-ui
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/api")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
