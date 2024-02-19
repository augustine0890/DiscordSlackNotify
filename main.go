package main

import (
	"discord-slack-notify/slack"
	"discord-slack-notify/util"
	"flag"
	"fmt"
	"log"
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

	log.Println("DiscordSlackNotify")
	// Initialize Slack client once
	slack.InitializeSlackClient()

	// Use sendSlackMessage whenever needed
	slack.SendSlackMessage("", "New message from Discord!")
}
