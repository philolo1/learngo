package main

import (
	"fmt"
	"os"
	// "time"
  "log"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

type Env struct {
  token string
  channelID string
}

func (env* Env) InitEnv() {
	// Load Env variables from .dot file
	godotenv.Load(".env")

	env.token = os.Getenv("SLACK_AUTH_TOKEN")
	env.channelID = os.Getenv("SLACK_CHANNEL_ID")
}

func main() {

  env := Env{}
  env.InitEnv()

  log.Printf("%v", env)

	client := slack.New(env.token, slack.OptionDebug(true))
	// Create the Slack attachment that we will send to the channel
	/*attachment := slack.Attachment{
		Pretext: "Super Bot Message",
		Text:    "some text",
		// Color Styles the Text, making it possible to have like Warnings etc.
		Color: "#36a64f",
		// Fields are Optional extra data!
		Fields: []slack.AttachmentField{
			{
				Title: "Date",
				Value: time.Now().String(),
			},
		},
	}
  */
	// PostMessage will send the message away.
	// First parameter is just the channelID, makes no sense to accept it
	_, timestamp, err := client.PostMessage(
		env.channelID,
		// uncomment the item below to add a extra Header to the message, try it out :)
		//slack.MsgOptionText("New message from bot", false),
    slack.MsgOptionBlocks(
      // テキストのみのセクションブロック
      &slack.SectionBlock{
        Type: slack.MBTSection,
        Text: &slack.TextBlockObject{
          Type: "mrkdwn",
          Text: "Hello, Assistant to the Regional Manager Dwight! *Michael Scott* wants to know where you'd like to take the Paper Company investors to dinner tonight.\n\n *Please select a restaurant:*",
        },
      },

      // 区切り線
      slack.NewDividerBlock(),

      // アクセサリつきのセクションブロック
      &slack.SectionBlock{
        Type: slack.MBTSection,
        Text: &slack.TextBlockObject{
          Type: "mrkdwn",
          Text: "*Farmhouse Thai Cuisine*\n:star::star::star::star: 1528 reviews\n They do have some vegan options, like the roti and curry, plus they have a ton of salad stuff and noodles can be ordered without meat!! They have something for everyone here",
        },
        Accessory: slack.NewAccessory(
          slack.NewImageBlockElement("https://s3-media2.fl.yelpcdn.com/bphoto/korel-1YjNtFtJlMTaC26A/o.jpg", "alt text for image"),
        ),
      },

      // ボタン
      &slack.ActionBlock{
        Type: slack.MBTAction,
        Elements: &slack.BlockElements{
          ElementSet: []slack.BlockElement{
            &slack.ButtonBlockElement{
              Type:  slack.METButton,
              Style: slack.StylePrimary,
              Text:  &slack.TextBlockObject{Type: "plain_text", Text: "Yes", Emoji: true},
              Value: "click_me_123",
              URL:   "https://google.com",
            },
            &slack.ButtonBlockElement{
              Type:  slack.METButton,
              Style: slack.StyleDanger,
              Text:  &slack.TextBlockObject{Type: "plain_text", Text: "No", Emoji: true},
              Value: "click_me_123",
              URL:   "https://google.com",
            },
          },
        },
      },
    ),
	)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Message sent at %s", timestamp)
}
