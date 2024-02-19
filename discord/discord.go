package discord

import (
	"discord-slack-notify/slack" // Import the custom slack package
	"os"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Delay the retrieval of ANNOUNCEMENTS_CHANNEL until needed
	announcementsChannel := os.Getenv("ANNOUNCEMENTS_CHANNEL")

	if m.ChannelID == announcementsChannel { // Replace with the target channel ID
		messageContent := m.Content
		slack.SendSlackMessage("C06H3GZNG00", messageContent) // Forward to Slack
	}
}
