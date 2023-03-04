// main.go
package main

import (
  "log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
  "encoding/json"
  "net/http"
  "strings"
  "os"
)

type Person struct {
    Message string `json:"message"`
    Age int `json:"age"`
}

func validateToken(str string) bool {
  values := strings.Split(str, " ")

  if (len(values) != 2) {
    return false
  }

  secret := values[1]

  return secret == os.Getenv("SECRET_TOKEN")
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  log.Println("Hello world")

  authHeader := request.Headers["authorization"]
  log.Printf("Headers %v", request.Headers["authorization"])

  if authHeader == "" {
    return events.APIGatewayProxyResponse{
      Headers: map[string]string{
        "Content-Type": "application/json",
      },
      Body: `{ "Error": "Authorization code required"}`,
      StatusCode: http.StatusUnauthorized,
    }, nil
  }

  if !validateToken(authHeader) {
    return events.APIGatewayProxyResponse{
      Headers: map[string]string{
        "Content-Type": "application/json",
      },
      Body: `{ "Error": "INVALID SECRET"}`,
      StatusCode: http.StatusUnauthorized,
    }, nil

  }




  person := Person{"Hello, World!", 42}
  jsonBytes, err := json.Marshal(person)
  if err != nil {
      panic(err)
  }

  bodyString := string(jsonBytes)


  response := events.APIGatewayProxyResponse{
    Headers: map[string]string{
      "Content-Type": "application/json",
    },
    Body: bodyString,
    StatusCode: 200,
  }

  return response, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}



