package config

import (
	"errors"
	"flag"
	"os"
	"strings"
)

func parseAddress(configField *string, defaultValue string) func(string) error {
	return func(value string) error {
		hp := strings.Split(value, ":")
		if len(hp) != 2 {
			return errors.New("need address in a form host:port")
		}

		if *configField == defaultValue {
			*configField = value
		}

		return nil
	}
}

func parseDBConnection(configField *string, defaultValue string) func(string) error {
	return func(value string) error {
		if *configField == defaultValue {
			*configField = value
		}

		return nil
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

type Config struct {
	// Адрес запуска HTTP-сервера
	Address string
	//Строка коннекта к базе данных
	DatabaseAddress string
}

// TODO: Тут переписать
func GetConfig() *Config {
	var cfg Config
	defaultAddress := "localhost:8080"
	defaultDatabase := "postgres://postgres:postgres@localhost:5432/gophermart?sslmode=disable"
	// defaultDatabase := ""

	flag.Func("a", "Адрес запуска HTTP-сервера", parseAddress(&cfg.Address, defaultAddress))
	flag.Func("d", "Строка коннекта к БД", parseDBConnection(&cfg.DatabaseAddress, defaultDatabase))

	cfg.Address = getEnv("RUN_ADDRESS", defaultAddress)
	cfg.DatabaseAddress = getEnv("DATABASE_URI", defaultDatabase)

	flag.Parse()

	return &cfg
}
