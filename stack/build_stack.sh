#!/bin/bash
docker build . \
  --build-arg CANDIDATE_NAME=test \
  --file build.Dockerfile \
  --tag aws-buildpack-build

docker build . \
  --build-arg CANDIDATE_NAME=test \
  --file run.Dockerfile \
  --tag aws-buildpack-run