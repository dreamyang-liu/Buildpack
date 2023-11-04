#!/bin/bash
runtime=$1

if [ -z "$runtime" ]; then
  echo "Usage: $0 <runtime>"
  exit 1
fi
rm -rf ./output/$runtime
mkdir -p ./output/$runtime/bin
go build -ldflags="-s -w" -o ./output/$runtime/bin/detect ./cmd/$runtime/main.go
go build -ldflags="-s -w" -o ./output/$runtime/bin/build ./cmd/$runtime/main.go
cp ./cmd/$runtime/buildpack.toml ./output/$runtime/buildpack.toml