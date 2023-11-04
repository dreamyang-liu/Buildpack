#!/bin/bash
cp ./aws-builder/builder.toml ./output/builder.toml
cd ./output
pack builder create io.buildpacks.stacks.amazonlinux.2023:drmyang --config ./builder.toml -v
