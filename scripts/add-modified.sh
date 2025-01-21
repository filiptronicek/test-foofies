#!/bin/sh
COMMIT_MSG_FILE="$1"

if [ ! -f "$COMMIT_MSG_FILE" ]; then
    echo "Error: Commit message file not found at $COMMIT_MSG_FILE"
    exit 1
fi

echo "$(cat "$COMMIT_MSG_FILE") (modified by lefthook)" > "$COMMIT_MSG_FILE"