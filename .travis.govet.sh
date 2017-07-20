#!/bin/bash
set -e

RED='\033[0;1;31m'
NC='\033[0m' # No Color

if [ -n "$(go tool vet src 2>&1)" ]; then
    echo -e "${RED}Vet errors found:"
    go tool vet src
    echo -e "${NC}"
    exit 1
fi