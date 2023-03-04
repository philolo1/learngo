package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"bytes"
	"encoding/json"
	"log"
)

type Request struct {
	Model    string `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	ID      string `json:"id"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason  string  `json:"finish_reason"`
}

func main() {
	// Load environment variables from .env file
	if os.Getenv("OPENAI_TOKEN") == "" {
		log.Fatal("OPENAI_TOKEN needs to be set")
	}

  instructions := "You are an AI that answers with code snippets in go. "

  if os.Getenv("OPENAI_INSTRUCTIONS") == "" {
    instructions = os.Getenv("OPENAI_INSTRUCTIONS")
  }

  args := os.Args
  var question string
  if len(args) == 2 {
    fmt.Println("The first argument is:", args[1])
    question = args[1]
  } else {
    log.Panic("Need to set question in first parameter")
  }

	// Set up HTTP client
	client := &http.Client{}

	// Set up HTTP request
	url := "https://api.openai.com/v1/chat/completions"
	requestBody, err := json.Marshal(Request{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{
				Role:    "system",
				Content: instructions,
			},
			{
				Role:    "user",
				Content: question,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_TOKEN"))

	// Send HTTP request
  log.Printf("Sending request...")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read HTTP response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal HTTP response body into struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Print HTTP response
	fmt.Printf("HTTP Response: %+v\n", response.Choices[0].Message.Content)
}
