#!/bin/bash
set -e

RED='\033[0;1;31m'
NC='\033[0m' # No Color

if [ -n "$(gofmt -l src)" ]; then
    echo -e "${RED}Go code is not formatted...failing files:${NC}"
    for i in `gofmt -l src`; do echo -e "${RED}$i${NC}"; done
    exit 1
fi