#!/bin/bash

set -e

# Detect OS and arch
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case $ARCH in
    x86_64) ARCH=amd64 ;;
    aarch64) ARCH=arm64 ;;
esac

# Download binary
URL="https://github.com/you/yourproject/releases/latest/download/mytool-${OS}-${ARCH}"
curl -L "$URL" -o mytool
chmod +x mytool
mv mytool /usr/local/bin/

echo "✅ mytool installed!"
