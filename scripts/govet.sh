#!/bin/bash

echo -n "Executing govet..."

# Gather files that failed vetting
FAILURES=()
for fileToVet in ${GOFILES}
    do
        output=`go tool vet ${fileToVet} 2>&1`
        if [ -n "${output}" ]; then
            FAILURES+=("${output}")
        fi
    done

# Output results
if [ -n "${FAILURES}" ]; then
    echo -e "${RED}FAILED${NC}"
    for indivError in "${FAILURES[@]}"
        do
            echo -e "${RED}    ${indivError}${NC}"
        done
    exit 1
else
    echo "PASSED"
fi