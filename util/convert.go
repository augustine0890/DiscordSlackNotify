package util

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// HumanReadableTime converts a Slack timestamp string to a human-readable time format.
func HumanReadableTime(timestamp string) string {
	// Splitting the timestamp to extract the seconds part before the dot
	parts := strings.Split(timestamp, ".")
	if len(parts) > 0 {
		// Parsing the seconds part to a Unix timestamp
		sec, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			// Log the error and return a default message or handle it as needed
			log.Printf("Error parsing timestamp: %s\n", err)
			return "Invalid timestamp"
		}
		// Formatting the Unix timestamp to a human-readable date string
		// Ensure the format string is "2006-01-02 15:04:05" to match Go's reference date
		t := time.Unix(sec, 0)
		humanReadable := t.Format("2006-01-02 15:04:05")
		return humanReadable
	}
	// Return a default or error message if the timestamp format is unexpected
	return "Invalid timestamp format"
}
