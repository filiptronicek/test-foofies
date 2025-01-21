#!/bin/sh
COMMIT_MSG_FILE=$1

# Append the text to the commit message
echo "$(cat $COMMIT_MSG_FILE) (modified by pre-commit)" > $COMMIT_MSG_FILE