package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	DB_DRIVER        string
	DATA_SOURCE_NAME string
}

var Env env

func init() {
	fmt.Println("in init")
	godotenv.Load(".env")
	Env.DB_DRIVER = os.Getenv("DB_DRIVER")
	Env.DATA_SOURCE_NAME = os.Getenv("DATA_SOURCE_NAME")
}
