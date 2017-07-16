#!/usr/bin/env bash

set -e
echo "" > coverage.txt

for d in $(go list ./test/...); do
    go test -coverpkg ./src/... $d -coverprofile=profile.out
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
