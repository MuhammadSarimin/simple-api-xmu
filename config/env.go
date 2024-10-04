package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/muhammadsarimin/simple-api-xmu/types"
	"github.com/pkg/errors"
)

var ENV types.Config

func Init() {

	if err := load(); err != nil {
		log.Fatal(err)
	}

	InitLog()
}

func load() error {

	if err := godotenv.Load(); err != nil {
		return errors.Wrap(err, "config/env: load .env file")
	}

	if err := env.Parse(&ENV); err != nil {
		return errors.Wrap(err, "config/env: parse config .env file")
	}

	if err := env.Parse(&ENV.DB); err != nil {
		return errors.Wrap(err, "config/env: parse redis .env file")
	}

	return nil
}
