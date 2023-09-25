#!/bin/bash -i
set -x

# Install go
GO_VERSION="go1.21.1.linux-amd64.tar.gz"
GO_PATH="export PATH=\$PATH:/usr/local/go/bin"
GO_HOME_PATH="export PATH=\$PATH:$HOME/go/bin"

if [ ! -f "$GO_VERSION" ]; then
  curl -LO https://go.dev/dl/$GO_VERSION
fi

sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf $GO_VERSION

if ! grep -q "$GO_PATH" "$HOME/.bashrc"; then
  echo "$GO_PATH" >> ~/.bashrc
fi

if ! grep -q "$GO_HOME_PATH" "$HOME/.bashrc"; then
  echo "$GO_HOME_PATH" >> ~/.bashrc
fi

set +x

