#!/usr/bin/env bash
set -euo pipefail

GO_VERSION="1.23.3"
NODE_VERSION="20"

# Ensure Homebrew
if ! command -v brew >/dev/null 2>&1; then
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
fi

brew update

# Docker Desktop (GUI app)
brew install --cask docker

# Go (pin exact version)
brew install go@${GO_VERSION%.*}
brew link --force go@${GO_VERSION%.*}

# Node/npm (LTS)
brew install node@${NODE_VERSION}
brew link --force node@${NODE_VERSION}

echo "Launch Docker Desktop once to finish setup."