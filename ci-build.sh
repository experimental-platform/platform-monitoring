#!/bin/bash
set -e

SRC_PATH=$(pwd)
PROJECT_NAME="github.com/$TRAVIS_REPO_SLUG"

export GO15VENDOREXPERIMENT=1
curl -L https://raw.githubusercontent.com/experimental-platform/misc/master/install-glide.sh | sh
cp $HOME/bin/glide .
docker run -v "${SRC_PATH}:/go/src/$PROJECT_NAME" -w "/go/src/$PROJECT_NAME" -e GO15VENDOREXPERIMENT=1 golang:1.5 /bin/bash -c "./glide up && go build -v"
