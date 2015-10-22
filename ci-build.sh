#!/bin/bash
set -e

SRC_PATH=$(pwd)
PROJECT_NAME="github.com/experimental-platform/platform-monitoring"

export GO15VENDOREXPERIMENT=1
./install-glide-v1.sh
cp $HOME/bin/glide .
docker run -v "${SRC_PATH}:/go/src/$PROJECT_NAME" -w "/go/src/$PROJECT_NAME" -e GO15VENDOREXPERIMENT=1 golang:1.5 /bin/bash -c './glide up && go build -v'
