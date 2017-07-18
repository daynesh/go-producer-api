#!/bin/bash
set -e

if [ -n "$(gofmt -l src test)" ]; then
    echo "Go code is not formatted:"
    gofmt -l src test
    exit 1
fi