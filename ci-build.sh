#!/bin/bash
# THIS ONLY WORK IN OUR CI!

docker run --rm -v /data/jenkins/jobs/${JOB_NAME}/workspace:/usr/src/monitoring -w /usr/src/monitoring golang:1.4 /bin/bash -c 'go get -d && go build -v'
