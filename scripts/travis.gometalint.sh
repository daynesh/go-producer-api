#!/bin/bash

RED='\033[0;1;31m'
NC='\033[0m' # No Color

# Install metalinter and any of its referenced linters
go get gopkg.in/alecthomas/gometalinter.v1
gometalinter.v1 --install

OUTPUT="$(gometalinter.v1 ./src/... -D gotype --deadline=180s 2>&1)"
if [ -n "${OUTPUT}" ]; then
    echo -e "${RED}gometalinter detected problems:"
    echo $OUTPUT
    echo -e "${NC}"
    exit 1
fi