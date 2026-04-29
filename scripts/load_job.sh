#!/bin/bash

LABEL="com.cjw.sqlite-sync"
PLIST_PATH="$HOME/Library/LaunchAgents/$LABEL.plist"

if [ -f "$PLIST_PATH" ]; then
    launchctl load "$PLIST_PATH"
    echo "Job $LABEL loaded successfully."
else
    echo "Error: Plist file not found at $PLIST_PATH. Run setup script first."
fi