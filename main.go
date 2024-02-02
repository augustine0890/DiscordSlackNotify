package main

import (
	"discord-slack-notify/util"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	stage := flag.String("stage", "prod", "The environment running")
	flag.Parse()

	// Loading environment variables
	err := util.LoadEnv(*stage)
	if err != nil {
		fmt.Printf("Error loading environmnet variables: %v\n", err)
	}
	log.Printf("Running with %v environment\n", *stage)

	fmt.Println("DiscordSlackNotify")

	runSlack()
}

func runSlack() {
	SLACK_AUTH_TOKEN := os.Getenv("SLACK_AUTH_TOKEN")
	CHANNEL_ID := "C06H3GZNG00" // #discord-noti channel

	api := slack.New(SLACK_AUTH_TOKEN)
	attachment := slack.Attachment{
		Pretext: "For testing purpose",
		Text:    "Have a good day!",
	}
	channelID, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionText("Main Message", false),
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		log.Fatalf("Error posting message: %s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelID, timestamp)
}
