#!/bin/bash

RED='\033[0;1;31m'
NC='\033[0m' # No Color

# Install metalinter and any of its referenced linters
go get -u gopkg.in/alecthomas/gometalinter.v1
gometalinter.v1 --install 1>&/dev/null

# Iterate through every directory and run gometalinter against it
# This is done instead of simply running gometalinter.v1 ./src/...
# as a hope that this would be faster
for directory in $(find ./src -type d -print);
do
    echo -e "Running gometalinter on ${directory}..."
    OUTPUT="$(gometalinter.v1 ${directory} -D gotype --deadline=600s 2>&1)"
    if [ -n "${OUTPUT}" ]; then
        echo -e "  ${RED}gometalinter detected problems:"
        echo "    ${OUTPUT}"
        echo -e "${NC}"
        exit 1
    fi
done