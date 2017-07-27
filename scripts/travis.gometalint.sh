#!/bin/bash
set -e

RED='\033[0;1;31m'
NC='\033[0m' # No Color

go get gopkg.in/alecthomas/gometalinter.v1
gometalinter.v1 --install

if [ -n "$(gometalinter.v1 ./src/... -D gotype --deadline=180s 2>&1)" ]; then
    echo -e "${RED}gometalinter detected problems:"
    gometalinter.v1 ./src/... -D gotype --deadline=180s
    echo -e "${NC}"
    exit 1
fi