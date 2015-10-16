#!/usr/bin/env bash

go get -d github.com/Masterminds/glide
cd /go/src/github.com/Masterminds/glide
git checkout 0.6.1
make install
