#!/usr/bin/env bash
set -e

REPO="dd3vahmad/plicense"
BINARY="plicense"
INSTALL_DIR="/usr/local/bin"
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Translate architecture
case $ARCH in
  x86_64) ARCH="x86_64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Get latest version
VERSION=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep tag_name | cut -d '"' -f 4)

echo "Installing $BINARY $VERSION for $OS-$ARCH..."

# Download and extract
URL="https://github.com/$REPO/releases/download/$VERSION/${BINARY}_${OS}_${ARCH}.tar.gz"
TMP_DIR=$(mktemp -d)
curl -L "$URL" -o "$TMP_DIR/${BINARY}.tar.gz"
tar -xzf "$TMP_DIR/${BINARY}.tar.gz" -C "$TMP_DIR"

# Move to PATH
sudo mv "$TMP_DIR/$BINARY" "$INSTALL_DIR/$BINARY"
sudo chmod +x "$INSTALL_DIR/$BINARY"

echo "$BINARY installed successfully!"
$BINARY --version
