#!/usr/bin/env bash

set -eux

TMPFILE=$(mktemp)
TMPDIR=$(mktemp -d)

curl -L https://github.com/Masterminds/glide/releases/download/0.6.1/glide-linux-amd64.zip > "$TMPFILE"
unzip "$TMPFILE" -d "$TMPDIR"
mkdir -p "$HOME/bin"
echo 'export PATH="$PATH:$HOME/bin"' >>  "$HOME/.profile"
mv "$TMPDIR/linux-amd64/glide" "$HOME/bin/glide"
#rm -rf "$TMPFILE" "$TMPDIR"
