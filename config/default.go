package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func GetEnvVariablesFromFile(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value

}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func GetEnvVariables(k string) (string, error) {
	v := os.Getenv(k)
	if v == "" {
		return "", fmt.Errorf("%s environment variable not set", k)
	}
	return v, nil
}

func GetEnvBool(key string) (bool, error) {
	value := strings.ToLower(os.Getenv(key))
	if value == "true" {
		return true, nil
	} else if value == "false" {
		return false, nil
	}
	return false, fmt.Errorf("invalid boolean value for environment variable %s: %s", key, value)
}
