#!/bin/bash

# Configuration
LABEL="com.cjw.sqlite-sync"
PLIST_PATH="$HOME/Library/LaunchAgents/$LABEL.plist"
WORKING_DIR=$(pwd)
BINARY_PATH="$WORKING_DIR/bin/sync"

cat <<EOF > "$PLIST_PATH"
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>$LABEL</string>
    <key>ProgramArguments</key>
    <array>
        <string>$BINARY_PATH</string>
    </array>
    <key>StartCalendarInterval</key>
    <array>
        <dict>
            <key>Hour</key>
            <integer>0</integer>
            <key>Minute</key>
            <integer>0</integer>
        </dict>
        <dict>
            <key>Hour</key>
            <integer>12</integer>
            <key>Minute</key>
            <integer>0</integer>
        </dict>
    </array>
    <key>WorkingDirectory</key>
    <string>$WORKING_DIR</string>
    <key>StandardErrorPath</key>
    <string>/tmp/sqlite-sync.err</string>
    <key>StandardOutPath</key>
    <string>/tmp/sqlite-sync.out</string>
</dict>
</plist>
EOF

echo "Created $PLIST_PATH"