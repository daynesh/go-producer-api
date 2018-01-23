#!/bin/bash
# Install metalinter and any of its referenced linters
echo -n "Installing gometalinter..."
env GOPATH=`pwd`/vendor go get -u gopkg.in/alecthomas/gometalinter.v1
env GOPATH=`pwd`/vendor:`pwd` vendor/bin/gometalinter.v1 --install 1>&/dev/null
echo "done"

# Now execute metalinter
echo -n "Executing gometalinter..."
SECONDS=0
OUTPUT="$(env GOPATH=`pwd`:`pwd`/vendor vendor/bin/gometalinter.v1 ./... --config=${DIR}/gometalinter.json 2>&1)"
if [ -n "${OUTPUT}" ]; then
    echo -e "${RED}FAILED"

    # Gather all errors into an array to iterate through
    IFS=$'\n' read -d '' -r -a errors <<< "$OUTPUT"
    
    # Now iterate through each error and print with a tab prefix
    for indivError in "${errors[@]}"
    do
        echo "    ${indivError}"
    done

    # Print execution time
    duration=$SECONDS
    echo "
    Completed in $(($duration / 60))m and $(($duration % 60))s"

    echo -e "${NC}"
    exit 1
else
    echo "PASSED"

    # Print execution time
    duration=$SECONDS
    echo "    Completed in $(($duration / 60))m and $(($duration % 60))s"
fi