package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	AppName       string `env:"APP_NAME,required"`
	AppPort       string `env:"APP_PORT" envDefault:":5000"`
	DbHost        string `env:"DB_HOST,required"`
	DbPort        string `env:"DB_PORT,required"`
	DbUsername    string `env:"DB_USERNAME,required"`
	DbName        string `env:"DB_NAME,required"`
	DbPassword    string `env:"DB_PASSWORD,required"`
	RedisHost     string `env:"REDIS_HOST,required"`
	RedisPort     string `env:"REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`
}

func NewConfig(file ...string) *Configuration {
	err := godotenv.Load(file...)
	if err != nil {
		log.Printf("File .env not found %q\n", file)
	}

	cfg := Configuration{}

	err = env.Parse(&cfg)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	return &cfg
}
