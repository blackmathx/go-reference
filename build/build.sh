#!/usr/bin/env bash
set -euo pipefail

# build/build.sh - build helper for POSIX environments (WSL, Git Bash, macOS, Linux)
ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT_DIR"

mkdir -p bin
echo "Building go-reference into ./bin/goref"

# build with module-aware mode
GOOS=${GOOS:-"$(go env GOOS)"}
GOARCH=${GOARCH:-"$(go env GOARCH)"}

go build -o bin/goref .

echo "Build complete: ./bin/goref"
