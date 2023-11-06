#!/bin/bash
cp ./builders/builder.toml ./output/builder.toml
cd ./output
pack builder create drmyang-builder:drmyang --config ./builder.toml -v
