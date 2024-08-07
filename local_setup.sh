#!/bin/bash

# Fetch secrets from GitHub and export them as environment variables

# Ensure gh is installed
if ! command -v gh &> /dev/null
then
    echo "gh (GitHub CLI) could not be found. Please install it."
    exit 1
fi

# Authenticate if not already authenticated
if ! gh auth status &> /dev/null
then
    echo "You need to log in to GitHub CLI."
    gh auth login
fi

REPO="OakLift/oaklift_webapp"

# Fetch and export secrets
export GOOGLE_CLIENT_ID=$(gh secret list -R "$REPO" --json name,value | jq -r '.[] | select(.name == "GOOGLE_CLIENT_ID") | .value' | base64 -d)
export GOOGLE_CLIENT_SECRET=$(gh secret list -R "$REPO" --json name,value | jq -r '.[] | select(.name == "GOOGLE_CLIENT_SECRET") | .value' | base64 -d)
export DB_USER=$(gh secret list -R "$REPO" --json name,value | jq -r '.[] | select(.name == "DB_USER") | .value' | base64 -d)
export DB_PASSWORD=$(gh secret list -R "$REPO" --json name,value | jq -r '.[] | select(.name == "DB_PASSWORD") | .value' | base64 -d)
export DB_NAME=$(gh secret list -R "$REPO" --json name,value | jq -r '.[] | select(.name == "DB_NAME") | .value' | base64 -d)
export DB_HOST=$(gh secret list -R "$REPO" --json name,value | jq -r '.[] | select(.name == "DB_HOST") | .value' | base64 -d)
export DB_PORT=5432

echo "Secrets have been set."
