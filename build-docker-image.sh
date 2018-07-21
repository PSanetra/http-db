#!/bin/sh

docker build -t http-db:local .

docker image rm $(docker image ls -q --filter="label=BUILD_IMAGE=true")
