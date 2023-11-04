#!/bin/bash
docker build .\
  --file ./stack/build.Dockerfile \
  --tag aws-buildpack-build 

docker build . \
  --file ./stack/run.Dockerfile \
  --tag aws-buildpack-run 