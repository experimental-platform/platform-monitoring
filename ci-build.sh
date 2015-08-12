#!/bin/bash

SRC_PATH=$(pwd)

docker run -v ${SRC_PATH}:/usr/src/monitoring -w /usr/src/monitoring golang:1.4 /bin/bash -c 'go get -d && go build -v'
