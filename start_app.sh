#!/bin/bash

# Assuming your Go binary name is discord-slack-notify, and it's in the same directory as the script

# Clean old binary
# Uncomment if you want to remove the previous build before a new one
# rm -f ./discord-slack-notify

# Build your Go application
# Adjust the path if your Go files are in a subdirectory
echo "Building application..."
go build -o discord-slack-notify

# Check if the application is already running
if [ -f discord-slack-notify.pid ]; then
    OLD_PID=$(cat discord-slack-notify.pid)
    if kill -0 $OLD_PID > /dev/null 2>&1; then
        echo "Stopping running application..."
        kill $OLD_PID
        # Wait for the old process to stop
        while kill -0 $OLD_PID > /dev/null 2>&1; do
            sleep 1
        done
    fi
fi

# Remove old log and pid files
rm -f discord-slack-notify.log discord-slack-notify.pid

# Print out a message indicating the application is starting
echo "Starting application..."

# Start your application in the background, redirect output to a log file,
# and store its PID in a separate file
./discord-slack-notify > discord-slack-notify.log 2>&1 & echo $! > discord-slack-notify.pid

# Print out a message indicating the application has been started
echo "Application has been started."
