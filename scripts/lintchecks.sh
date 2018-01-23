#!/bin/bash

# Get current directory for this script
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Get go files excluding ./vendor and ./.glide
GOFILES=$(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./.glide/*")

# Output text coloring
RED='\033[0;1;31m'
NC='\033[0m' # No Color

# Now execute all lint checks
source $DIR/gofmt.sh
source $DIR/govet.sh
source $DIR/gometalint.sh