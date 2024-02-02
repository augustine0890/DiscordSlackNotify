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
	message := `
	"Super Chicken TD 🐔🐔 " Feedback Contest

	@everyone GM fam! 🎮 
	👏🏻 Have you visited our new channel, 'Spotlight,'( ⁠spotlight ) where we'll be showcasing exciting games from GameOVEN every week? 
	To share the excitement of games in GameOVEN, we're hosting a game feedback contest 🌟 

	Congrats @starG !!!
	We sent the 1000 points as a last prize!

	🌈 Event Period:
	From now until February 7th, ending at 23:59 (UTC+0).

	🐾 How to Participate:
	Read the game spotlight (⁠spotlight⁠Super Chicken TD 🐔🐔 : My crop…)
	Dive into the "Super Chicken TD" game.
	Play, conquer levels, and share your feedback as a comment under the post (⁠spotlight⁠Super Chicken TD 🐔🐔 : My crop…).

	🎁 Win Big!
	Three lucky winners will be chosen and each receive 1000 Discord Points as a reward for sharing their experiences and feedback! Your thoughts are valuable, and we want to celebrate them!

	🏆 Prize Announcement:
	Winners will be revealed on 14th Feb 2024!  Stay tuned to discover if you're one of the lucky victors! 🎉 
	`

	api := slack.New(SLACK_AUTH_TOKEN)
	attachment := slack.Attachment{
		Pretext: "For testing purpose",
		Text:    message,
	}
	channelID, timestamp, err := api.PostMessage(
		CHANNEL_ID,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		log.Fatalf("Error posting message: %s\n", err)
	}

	log.Printf("Message successfully sent to Channel %s at %s\n", channelID, timestamp)
}
