package env

import (
  "github.com/joho/godotenv"
  "os"
  "fmt"
)

const (
  ENV_DATABASE_CONNECTION_STR = "DATABASE_CONNECTION_STRING"
  ENV_PORT = "PORT"
  ENV_REDIS_PASSWORD = "REDIS_PASSWORD"
  ENV_REDIS_HOST = "REDIS_HOST"
)

type EnvStruct struct {
  DatabaseConnectionString string
  Port string
  RedisPassword string
  RedisHost string
}

var environment *EnvStruct

func init() {
  fmt.Println("Loading environment variables")
  if err := godotenv.Load(); err != nil {
    panic("Error loading .env file")
  }

  if (os.Getenv(ENV_DATABASE_CONNECTION_STR) == "") {
    panic("environment DATABASE_CONNECTION_STR is not set")
  }


  if (os.Getenv(ENV_REDIS_PASSWORD) == "") {
    panic("environment REDIS_PASSWORD is not set")
  }

  if (os.Getenv(ENV_REDIS_HOST) == "") {
    panic("environment REDIS_HOST is not set")
  }

  environment = &EnvStruct{
    DatabaseConnectionString: os.Getenv(ENV_DATABASE_CONNECTION_STR),
    Port: os.Getenv(ENV_PORT),
    RedisPassword: os.Getenv(ENV_REDIS_PASSWORD),
    RedisHost: os.Getenv(ENV_REDIS_HOST),
  }

  if environment.Port == "" {
    environment.Port = "9010"
  }
  fmt.Println("Loaded env variables")
}

func GetEnv() *EnvStruct {
  return environment
}
