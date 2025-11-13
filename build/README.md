This folder contains simple build helpers for the project.

Recommended commands (from repository root):

# Build for development (binary in ./bin)
# POSIX (bash / WSL / Git Bash):
./build/build.sh

# Windows (cmd.exe):
build\\build.bat

# Or use go directly:
# build and run in one step:
go run .

# build a binary (output in ./bin):
mkdir -p bin && go build -o bin/goref .

# Notes:
# - The scripts are small wrappers around `go build` with sensible defaults.
# - Use the OS-specific script on Windows or WSL as appropriate.
