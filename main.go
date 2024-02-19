package main

import (
	"discord-slack-notify/discord"
	"discord-slack-notify/slack"
	"discord-slack-notify/util"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
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
	// slack.SendSlackMessage("", "New message from Discord!")

	// Create a new Discord session using the provided bot token
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Println("Error creating Discord session: ", err)
	}

	// Register message handler
	dg.AddHandler(discord.MessageHandler)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening Discord session: %s\n", err)
	}
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
