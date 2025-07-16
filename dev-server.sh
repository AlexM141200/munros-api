#!/bin/bash

# Development server with auto-reload for Munros API

echo "Starting development server with auto-reload..."
echo "This will watch for changes and restart the server automatically."
echo "Press Ctrl+C to stop."
echo ""

# Function to build and run the server
build_and_run() {
    echo "Building application..."

    # Generate templ files
    templ generate
    if [ $? -ne 0 ]; then
        echo "Failed to generate templates"
        return 1
    fi

    # Build the application
    go build -o bin/munros-api ./src/cmd/main.go
    if [ $? -ne 0 ]; then
        echo "Build failed"
        return 1
    fi

    echo "Starting server at http://localhost:8080"
    ./bin/munros-api &
    SERVER_PID=$!
    return 0
}

# Function to kill the server
kill_server() {
    if [ ! -z "$SERVER_PID" ]; then
        kill $SERVER_PID 2>/dev/null
        wait $SERVER_PID 2>/dev/null
        SERVER_PID=""
    fi
}

# Initial build and run
build_and_run
if [ $? -ne 0 ]; then
    echo "Initial build failed. Exiting."
    exit 1
fi

# Watch for changes in Go files and templ files
while true; do
    # Use find to watch for changes (basic file watching)
    if command -v fswatch >/dev/null 2>&1; then
        # Use fswatch if available (better option)
        fswatch -1 src/ | while read file; do
            echo "Changes detected in: $file"
            kill_server
            sleep 1
            build_and_run
        done
    else
        # Fallback: simple polling method
        echo "Note: Install 'fswatch' for better file watching (brew install fswatch)"
        echo "Using basic polling method..."

        LAST_MODIFIED=$(find src/ -name "*.go" -o -name "*.templ" | xargs stat -f %m 2>/dev/null | sort -n | tail -1)

        while true; do
            sleep 2
            CURRENT_MODIFIED=$(find src/ -name "*.go" -o -name "*.templ" | xargs stat -f %m 2>/dev/null | sort -n | tail -1)

            if [ "$CURRENT_MODIFIED" != "$LAST_MODIFIED" ]; then
                echo "Changes detected, restarting server..."
                kill_server
                sleep 1
                build_and_run
                LAST_MODIFIED=$CURRENT_MODIFIED
            fi
        done
    fi
done

# Cleanup on exit
trap 'kill_server; exit 0' INT TERM
