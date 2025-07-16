#!/bin/bash

# Build and run the Munros API server

# Build the application
echo "Building the application..."
go build -o bin/munros-api ./src/cmd/main.go

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "Build successful. Starting server..."
    echo "Server will be available at: http://localhost:8080"
    echo "Press Ctrl+C to stop the server"
    echo ""

    # Run the server
    ./bin/munros-api
else
    echo "Build failed. Please check for errors."
    exit 1
fi
