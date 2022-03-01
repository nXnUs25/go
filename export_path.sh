#!/usr/bin/env bash

### execute this script by doing source script call
### as example:
# source export_path.sh
### this will append GOPATH variable with your go project ROOT Path
### so that src pkg and bin floders will be visible in you projects

PWD=$(pwd)

export GOPATH="$GOPATH:$PWD"

#source "${PWD}/$(basename $0)"