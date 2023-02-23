package env

import (
  "github.com/joho/godotenv"
  "os"
  "fmt"
)

const (
  ENV_DATABASE_CONNECTION_STR = "DATABASE_CONNECTION_STRING"
  ENV_PORT = "PORT"
)

type EnvStruct struct {
  DatabaseConnectionString string
  Port string
}

var environment *EnvStruct

func init() {
  fmt.Println("Loading environment variables")
  if err := godotenv.Load(); err != nil {
    panic("Error loading .env file")
  }

  if (os.Getenv(ENV_DATABASE_CONNECTION_STR) == "") {
    panic("environment ENV_DATABASE_CONNECTION_STR is not set")
  }

  environment = &EnvStruct{
    DatabaseConnectionString: os.Getenv(ENV_DATABASE_CONNECTION_STR),
    Port: os.Getenv(ENV_PORT),
  }

  if environment.Port == "" {
    environment.Port = "9010"
  }
  fmt.Println("Loaded env variables")
}

func GetEnv() *EnvStruct {
  return environment
}
