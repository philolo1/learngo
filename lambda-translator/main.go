// main.go
package main

import (
  "log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/joho/godotenv"
  "encoding/json"
  "net/http"
  "github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
  "context"
  "os"
  "fmt"
	openai "github.com/sashabaranov/go-openai"
)

type Env struct {
  slackToken string
  openaiApiKey string
}

var (
    ErrorLog   *log.Logger
    InfoLog   *log.Logger
    aiClient *openai.Client
    slackClient *slack.Client
    env Env
)

func (env* Env) InitEnv() {
	// Load Env variables from .dot file
	godotenv.Load(".env")

	env.slackToken = os.Getenv("SLACK_AUTH_TOKEN")
	env.openaiApiKey = os.Getenv("OPENAI_API_KEY")

  ErrorLog = log.New(os.Stdout, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
  InfoLog = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
  aiClient = openai.NewClient(env.openaiApiKey)
}

func createErrorResponse(message string, statusCode int)  (events.APIGatewayProxyResponse, error) {
  return events.APIGatewayProxyResponse{
        Headers: map[string]string{
          "Content-Type": "application/json",
        },
        Body: fmt.Sprintf(`{ "Error": "%v"}`, message),
        StatusCode: statusCode,
      }, nil

}

func translate(message string) string {
	resp, err := aiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
        {
          Role: openai.ChatMessageRoleSystem,
          Content: "You are a translator who translates from Japanese to English and the other way around. You only answer with the translation. You do not include pronounciation.",
        },
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello",
				},
        {
          Role:    openai.ChatMessageRoleAssistant,
					Content: "こんにちは",
        },
        {
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},

			},
		},
	)

  if err != nil {
    ErrorLog.Printf("Could not translate %v", err)
    return ""
  }

  InfoLog.Printf("Translate %v\n", resp.Choices[0].Message.Content)

	return resp.Choices[0].Message.Content
}

func sendMessage(channelID string, message string, threadTimestamp string) error {
  params := slack.PostMessageParameters{
    ThreadTimestamp: threadTimestamp,
  }

  resultMessage := translate(message)

  _, _, err := slackClient.PostMessage(channelID, slack.MsgOptionText(resultMessage, false), slack.MsgOptionPostMessageParameters(params))
  if err != nil {
    return err
  }
  return nil
}


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

  env.InitEnv()
  InfoLog.Println("Hello world")

  slackClient = slack.New(env.slackToken, slack.OptionDebug(true))

	body := request.Body

	event, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
		if err != nil {
			ErrorLog.Printf("Error parsing event: %v", err)
      return createErrorResponse(fmt.Sprintf("Error parsing event: %v", err), http.StatusInternalServerError)
		}


		switch event.Type {
		case slackevents.URLVerification:
			var challengeResponse slackevents.ChallengeResponse
			err = json.Unmarshal([]byte(body), &challengeResponse)
			if err != nil {
				ErrorLog.Printf("Error unmarshalling URL verification response: %s", err)
        return createErrorResponse("Error unmarshalling URL verification response:", http.StatusInternalServerError)
			}

      return events.APIGatewayProxyResponse{
        Headers: map[string]string{
          "Content-Type": "text/plain",
        },
        Body: challengeResponse.Challenge,
        StatusCode: http.StatusOK,
      }, nil

		case slackevents.CallbackEvent:
			// handle the incoming event

      msgEvt := event.InnerEvent.Data.(*slackevents.MessageEvent)

      InfoLog.Printf("Event %v\n", msgEvt)

      if msgEvt.Type == "message"  {
        // process the message
        text := msgEvt.Text
        InfoLog.Printf("text %v", text)

        if len(msgEvt.ThreadTimeStamp) == 0 {
          slackClient.AddReaction("thumbsup", slack.NewRefToMessage(
            msgEvt.Channel,
            msgEvt.EventTimeStamp,
          ))
          InfoLog.Printf("time %v", msgEvt.EventTimeStamp)

          sendMessage(msgEvt.Channel, text, msgEvt.EventTimeStamp)
          return events.APIGatewayProxyResponse{
            Headers: map[string]string{
              "Content-Type": "application/json",
            },
            Body: `{ "Status": "Message was handled"}`,
            StatusCode: http.StatusOK,
          }, nil

        }
      } else {
        InfoLog.Printf("Undhandled message %v", msgEvt.Type)
      }
		default:
			// ignore unrecognized events
		}

    return events.APIGatewayProxyResponse{
      Headers: map[string]string{
        "Content-Type": "application/json",
      },
      Body: `{ "Status": "Nothing needs to be done"}`,
      StatusCode: http.StatusOK,
    }, nil










}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}



