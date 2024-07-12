#!/bin/bash

# Check if the number of commits is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <number_of_commits>"
  exit 1
fi

# Number of commits to create
NUM_COMMITS=$1

# Loop to create the specified number of commits
for ((i=1; i<=NUM_COMMITS; i++))
do
  echo "f" >> README.md
  git add README.md
  git commit -m "Append f to README.md - commit $i"
done
