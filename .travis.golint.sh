#!/bin/bash
set -e

RED='\033[0;1;31m'
NC='\033[0m' # No Color

if [ -n "$(golint ./src/... 2>&1)" ]; then
    echo -e "${RED}golint detected problems:"
    for i in `golint ./src/...`; do echo $i; done
    echo -e "${NC}"
    exit 1
fi