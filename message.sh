#!/bin/sh
# pre-commit automatically passes the commit message file path as $1
COMMIT_MSG_FILE="$1"

if [ ! -f "$COMMIT_MSG_FILE" ]; then
    echo "Error: Commit message file not found at $COMMIT_MSG_FILE"
    exit 1
fi

# Read and modify the commit message
echo "$(cat "$COMMIT_MSG_FILE") (modified by pre-commit)" > "$COMMIT_MSG_FILE"