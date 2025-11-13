This folder contains build helpers for the project.

Recommended commands (from repository root):

# Build for development (binary in ./bin)
# POSIX (bash / WSL / Git Bash):
./build/build.sh

# Or use go directly to run in one step:
go run .

# build a binary (output in ./bin):
mkdir -p bin && go build -o bin/goref .