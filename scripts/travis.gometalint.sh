#!/bin/bash
RED='\033[0;1;31m'
NC='\033[0m' # No Color

# Install metalinter and any of its referenced linters
env GOPATH=`pwd`/vendor go get -u gopkg.in/alecthomas/gometalinter.v1
env GOPATH=`pwd`/vendor:`pwd` vendor/bin/gometalinter.v1 --install 1>&/dev/null

OUTPUT="$(env GOPATH=`pwd`:`pwd`/vendor vendor/bin/gometalinter.v1 ./src/app/... -D gotype --deadline=600s 2>&1)"
if [ -n "${OUTPUT}" ]; then
    echo -e "${RED}gometalinter detected problems:"
    echo "    ${OUTPUT}"
    echo -e "${NC}"
    exit 1
fi