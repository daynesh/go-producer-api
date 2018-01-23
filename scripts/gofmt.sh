#!/bin/bash
set -e

echo -n "Executing gofmt..."
if [ -n "$(gofmt -l ${GOFILES})" ]; then
    echo -e "${RED}FAILED${NC}"
    for i in `gofmt -l ${GOFILES}`; do echo -e "${RED}    $i${NC}"; done
    exit 1
else
    echo "PASSED"
fi

set +e