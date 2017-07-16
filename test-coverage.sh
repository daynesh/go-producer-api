#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./test/...); do
    go test -coverpkg ./src/... $d -coverprofile=indivprofile.out
    if [ -f indivprofile.out ]; then
        cat indivprofile.out >> coverage.txt
        rm indivprofile.out
    fi
done
