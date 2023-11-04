#!/bin/bash
cp ./builders/builder.toml ./builders/builder.toml
cd ./output
pack builder create drmyang-builder:drmyang --config ./builder.toml -v
