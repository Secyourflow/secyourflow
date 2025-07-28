#!/bin/bash

# Set the flags
FLAGS=""
if [ -n "$SECYOURFLOW_FLAGS" ]; then
    FLAGS="$SECYOURFLOW_FLAGS"
fi

# Set the app url
APP_URL=""
if [ -n "$SECYOURFLOW_APP_URL" ]; then
    APP_URL="$SECYOURFLOW_APP_URL"
fi

/usr/local/bin/secyourflow serve --http 0.0.0.0:8080 --flags "$FLAGS" --app-url "$APP_URL"