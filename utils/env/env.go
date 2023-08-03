package env

import "os"

func GetEnv(env, defaultValues string) string {
	environment := os.Getenv(env)

	if environment == "" {
		return defaultValues
	}

	return environment
}
