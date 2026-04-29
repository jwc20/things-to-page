#!/bin/bash

LABEL="com.cjw.sqlite-sync"
PLIST_PATH="$HOME/Library/LaunchAgents/$LABEL.plist"

launchctl unload "$PLIST_PATH"
echo "Job $LABEL unloaded."