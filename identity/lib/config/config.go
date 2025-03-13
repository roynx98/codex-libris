package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Port        string `validate:"required,numeric"`
	Env         string `validate:"required,oneof=development test production"`
	ServiceName string `validate:"required"`
}

var Default Config

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func loadConfig() {
	env := getEnv("IDENTITY_ENV", "development")

	godotenv.Load(".env." + env)

	Default = Config{
		Port:        getEnv("PORT", "3000"),
		Env:         getEnv("IDENTITY_ENV", ""),
		ServiceName: getEnv("SERVICE_NAME", ""),
	}

	validate := validator.New()
	err := validate.Struct(Default)

	if err != nil {
		panic(err)
	}

}

func init() {
	loadConfig()
}
