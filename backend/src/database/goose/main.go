package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"github.com/rs/zerolog"
	"gitlab.com/mastocred/web-app/utility/environment"
	"gitlab.com/mastocred/web-app/utility/helper"
)

const dir = "migration"

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	gooseLogger := logger.With().Str(helper.LogStrKeyModule, "goose").Logger()
	_ = flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 2 {
		flags.Usage()
		return
	}

	command := args[0]

	// get the environment
	env, err := environment.NewLoadFromFile("../../.env")
	if err != nil {
		gooseLogger.Panic().Err(err)
	}

	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			env.Get("PG_USER"),
			env.Get("PG_PASSWORD"),
			env.Get("PG_HOST"),
			env.Get("PG_PORT"),
			env.Get("PG_DATABASE")))

	if err != nil {
		gooseLogger.Fatal().Err(err).Msgf("goose %v: %v", command, err)
	}

	var arguments []string
	for _, val := range args[1:] {
		if len(val) > 0 {
			arguments = append(arguments, val)
		}
	}

	gooseLogger.Info().Msgf("running goose %s %v : args=%d", command, arguments, len(arguments))
	if err := goose.Run(command, db, dir, arguments...); err != nil {
		gooseLogger.Fatal().Err(err).Msgf("goose %v: %v", command, err)
	}
}
