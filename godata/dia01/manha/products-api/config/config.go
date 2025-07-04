package config

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

var apiConfig *ApiConfig

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Database string
}

type ApiConfig struct {
	ApiToken string
	Database DatabaseConfig
}

func Init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	apiConfig = &ApiConfig{
		ApiToken: os.Getenv("TOKEN"),
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
	}
}

func GetConfig() *ApiConfig {
	if apiConfig == nil {
		Init()
	}
	return apiConfig
}

func GetEnv[T any](key string) T {
	var fallback T

	if apiConfig == nil {
		return fallback
	}

	val := reflect.ValueOf(*apiConfig)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == key && field.IsExported() {
			fieldVal := val.Field(i)
			if value, ok := fieldVal.Interface().(T); ok {
				return value
			}
		}
	}

	return fallback
}

func GetEnvWithFallback[T any](key string, fallback T) T {
	if apiConfig == nil {
		return fallback
	}

	val := reflect.ValueOf(*apiConfig)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == key && field.IsExported() {
			fieldVal := val.Field(i)
			if value, ok := fieldVal.Interface().(T); ok {
				return value
			}
		}
	}

	return fallback
}
