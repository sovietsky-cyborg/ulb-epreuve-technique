package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

type StatusError struct {
	Code int
	Err  error
}

func GetEnv(key string) string {
	env := os.Getenv(key)
	if env == "" {
		envFile, err := godotenv.Read(".env")
		if err != nil {
			panic(err)
		}
		return envFile[key]
	} else {
		return env
	}
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}
