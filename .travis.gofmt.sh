#!/bin/bash
set -e

RED='\033[0;31m'
NC='\033[0m' # No Color

if [ -n "$(gofmt -l src test)" ]; then
    echo -e "${RED}Go code is not formatted:"
    gofmt -l src test
    echo -e "${NC}"
    exit 1
fi