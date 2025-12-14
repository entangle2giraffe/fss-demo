#!/usr/bin/env bash
set -euo pipefail

# Versions
GO_VERSION="1.23.3"
NODE_VERSION="20"
DOCKER_CHANNEL="stable"

# Update
sudo apt-get update

# Docker
sudo apt-get install -y ca-certificates curl gnupg lsb-release
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/$(. /etc/os-release && echo "$ID")/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/$(. /etc/os-release && echo "$ID") \
  $(lsb_release -cs) ${DOCKER_CHANNEL}" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
sudo usermod -aG docker "$USER"

# Go
GO_TARBALL="go${GO_VERSION}.linux-amd64.tar.gz"
curl -LO "https://go.dev/dl/${GO_TARBALL}"
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf "${GO_TARBALL}"
rm "${GO_TARBALL}"
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc

# Node/npm (NodeSource)
curl -fsSL "https://deb.nodesource.com/setup_${NODE_VERSION}.x" | sudo -E bash -
sudo apt-get install -y nodejs

echo "Done. Re-login for docker group + PATH changes."