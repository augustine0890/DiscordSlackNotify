package slack

import (
	"log"
	"os"

	"github.com/slack-go/slack"
)

const CHANNEL_ID = "C06H3GZNG00" // Default channel ID
var api *slack.Client            // Global Slack client variable

func InitializeSlackClient() {
	SLACK_AUTH_TOKEN := os.Getenv("SLACK_AUTH_TOKEN")
	api = slack.New(SLACK_AUTH_TOKEN)
}

func SendSlackMessage(channelID string, message string) {
	// Use default CHANNEL_ID if channelID is not specified
	if channelID == "" {
		channelID = CHANNEL_ID
	}

	attachment := slack.Attachment{
		Pretext: "For Testing Purpose",
		Text:    message,
	}

	_, _, err := api.PostMessage(
		channelID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		log.Fatalf("Error posting message: %s\n", err)
	}

	// log.Printf("Message successfully sent to Channel %s at %s\n", channelID, timestamp)
}
