#!/bin/bash

# Define variables
OUTPUT_NAME="submission"
SCRIPT_DIR=$(dirname "$0")
SOURCE_FILE="$SCRIPT_DIR/cmd/main.go"

# Ensure the Go environment is set correctly
export PATH=$PATH:/usr/local/go/bin
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

# Remove current file
rm "$OUTPUT_NAME"

# Get latest artifact from workflows
# This script will need to be produced by host due to github PAT
bash artifact.sh

# Verify the binary format
echo "Verifying the binary..."
file "$OUTPUT_NAME"

docker stop genvoice-backend
docker rm genvoice-backend
docker rmi genvoice-backend

docker build -t genvoice-backend .
docker create --name genvoice-backend -p 8080:8080 genvoice-backend
docker start genvoice-backend