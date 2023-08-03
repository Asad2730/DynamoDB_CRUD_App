package config

import (
	"strconv"

	"github.com/Asad2730/DynamoDB_CRUD_App/utils/env"
)

type Config struct {
	Port        int
	TimeOut     int
	Dialect     string
	DataBaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        parseEnvToInt("PORT", "8080"),
		TimeOut:     parseEnvToInt("TimeOut", "30"),
		Dialect:     env.GetEnv("DIALECT", "sqlite3"),
		DataBaseURI: env.GetEnv("DATABASE_URI", ":memory:"),
	}
}

func parseEnvToInt(envName, defaultValue string) int {

	num, err := strconv.Atoi(env.GetEnv(envName, defaultValue))

	if err != nil {
		return 0
	}

	return num
}
