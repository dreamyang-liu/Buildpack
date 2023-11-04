#!/bin/bash
dependency=$1

if [ -z "$dependency" ]; then
  echo "Usage: $0 <dependency>"
  exit 1
fi
rm -rf ./output/$dependency
mkdir -p ./output/$dependency/bin
go build -ldflags="-s -w" -o ./output/$dependency/bin/main ./cmd/$dependency/main.go
cp ./output/$dependency/bin/main ./output/$dependency/bin/detect
cp ./output/$dependency/bin/main ./output/$dependency/bin/build
cp ./cmd/$dependency/buildpack.toml ./output/$dependency/buildpack.toml