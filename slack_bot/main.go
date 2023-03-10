package main

import (
	"os"
	// "time"
  "log"
  "context"

	"github.com/joho/godotenv"
  "github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
  "net/http"
  "bytes"
   "encoding/json"
	openai "github.com/sashabaranov/go-openai"
)

type Env struct {
  slackToken string
  channelID string
  openaiApiKey string
}

func (env* Env) InitEnv() {
	// Load Env variables from .dot file
	godotenv.Load(".env")

	env.slackToken = os.Getenv("SLACK_AUTH_TOKEN")
	env.channelID = os.Getenv("SLACK_CHANNEL_ID")
	env.openaiApiKey = os.Getenv("OPENAI_API_KEY")
}

var (
    ErrorLog   *log.Logger
    InfoLog   *log.Logger
    aiClient *openai.Client
    slackClient *slack.Client
    env Env
)

func init() {
  env.InitEnv()
  ErrorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
  InfoLog = log.New(os.Stdout, "[INFO] ", 0)
  aiClient = openai.NewClient(env.openaiApiKey)
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


func main() {

  ErrorLog.Println("Error")

  // log.Printf("%v", env)
  slackClient = slack.New(env.slackToken, slack.OptionDebug(true))

  if (slackClient != nil) {
    log.Println("Start")
    log.Println("Start")
  }

	http.HandleFunc("/slack/events", func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()

		event, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
		if err != nil {
			log.Printf("Error parsing event: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}


		switch event.Type {
		case slackevents.URLVerification:
			var challengeResponse slackevents.ChallengeResponse
			err = json.Unmarshal([]byte(body), &challengeResponse)
			if err != nil {
				log.Printf("Error unmarshalling URL verification response: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "text")
			w.Write([]byte(challengeResponse.Challenge))

		case slackevents.CallbackEvent:
			// handle the incoming event

      msgEvt := event.InnerEvent.Data.(*slackevents.MessageEvent)

      log.Printf("Event %v\n", msgEvt)

      if msgEvt.Type == "message"  {
        // process the message
        text := msgEvt.Text
        log.Printf("text %v", text)

        if len(msgEvt.ThreadTimeStamp) == 0 {
          slackClient.AddReaction("thumbsup", slack.NewRefToMessage(
            msgEvt.Channel,
            msgEvt.EventTimeStamp,
          ))
          log.Printf("time %v", msgEvt.EventTimeStamp)

          go sendMessage(msgEvt.Channel, text, msgEvt.EventTimeStamp)
        }
      }
		default:
			// ignore unrecognized events
		}
	})

	InfoLog.Println("Listening for events on /slack/events:3000")
	http.ListenAndServe(":3000", nil)

}
