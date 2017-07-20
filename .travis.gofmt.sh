#!/bin/bash
set -e

RED='\033[0;1;31m'
NC='\033[0m' # No Color

if [ -n "$(gofmt -l src test)" ]; then
    echo -e "${RED}Go code is not formatted...failing files:"
    for i in `gofmt -l src test`; do echo -e $i; done
    echo -e "${NC}"
    exit 1
fi