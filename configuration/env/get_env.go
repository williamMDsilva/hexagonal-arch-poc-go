package env

import "os"

func GetEnv() string {
	return os.Getenv("DATABASE_URL")
}
