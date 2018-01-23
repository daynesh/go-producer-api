#!/bin/sh
# Generate test coverage statistics for Go packages.
#
# Works around the fact that `go test -coverprofile` currently does not work
# with multiple packages, see https://code.google.com/p/go/issues/detail?id=6909
#
# Usage: scripts/generate-test-coverage.sh [--html|--coveralls|--codecov]
#
#     --html      Additionally create HTML report and open it in browser
#     --coveralls Push coverage statistics to coveralls.io
#

set -e

# Variable Definitions
RED="\033[0;31m"             # ANSI color code Red - Used to color encode error text
NOCOLOR="\033[0m"            # ANSI color code No color - Used to unset previously set color encoding
workdir=.cover               # directory to store all coverage output into (include intermediate files)
profile="$workdir/cover.out" # file to store aggregated coverate output
mode=count
PROJ_TYPE="UNDEFINED"        # Can either be APP or LIB
COVERPKG_OPTIONS=""          # Used to specify value for -coverpkg option for "go test" (set differently depending on APP or LIB)
GOPATH=""                    # Used for executing "go test" and "go tool cover" (set differently depending on APP or LIB)

# Checks if project to be covered is an application or a library
# Log is simple: check if ./src/app exists
#    - projects with a ./src directory allows 'go test -coverprofile' to
#      generate coverage details that can be easily parsed by 'go tool cover'
#    - also ensures ./src/app exists to conform to our application directory
#      structure standards
check_project_type() {
    if [ -d "${PWD}/src" ]; then
        if [ -d "${PWD}/src/app" ]; then
            PROJ_TYPE="APP"
        else
            echo "${RED}    ./src/app must exist as this project is assumed to be an application and not a library${NC}"
            exit 1
        fi
    else
        PROJ_TYPE="LIB"
    fi
}

# Construct a temp GOPATH outside of our application's directory aka pwd
# This is needed to significantly ease coverage profile generation and summary for libraries
define_custom_environment() {
    # Name the directory (to be defined in /tmp)
    baseworkdir=${PWD##*/}

    # Remove any existing directory with the same name
    rm -rf /tmp/${baseworkdir}-go-workspace

    # Create go-friendly directory structure
    mkdir -p /tmp/${baseworkdir}-go-workspace/src/github.com/discovery-digital

    # Link to our existing project
    ln -s $PWD /tmp/${baseworkdir}-go-workspace/src/github.com/discovery-digital/

    # Finally, set GOPATH to /tmp and change directory
    export GOPATH=/tmp/${baseworkdir}-go-workspace
    cd /tmp/${baseworkdir}-go-workspace/src/github.com/discovery-digital/${baseworkdir}
}

# Generate coverage file (e.g., .cover/cover.out)
# Func Parameter #1: list of go packages that may include tests
generate_coverage_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    # Iterate through each specified package and generate an individual coverage file
    for pkg in "$@"; do
        f="$workdir/$(echo $pkg | tr / -).cover"
        env GOPATH=$GOPATH go test -covermode="$mode" -coverprofile="$f" $COVERPKG_OPTIONS $pkg 2>/dev/null
    done

    # Now aggregate all individual coverage file into one master file to be used by "go tool cover"
    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
    echo "Generated cover profile: $profile"
}

# Output cover profile details
show_cover_report() {
    env GOPATH=$GOPATH go tool cover -${1}="$profile"
}

# Send coverage details to coverals.io
push_to_coveralls() {
    env GOPATH=`pwd`/vendor go get github.com/mattn/goveralls
    echo "Pushing coverage statistics to coveralls.io"

    if [ "$COVERALLS_TOKEN" ]; then
        env GOPATH=$GOPATH vendor/bin/goveralls -coverprofile="$profile" -service=travis-ci -repotoken $COVERALLS_TOKEN
    else
        echo >&2 "$RED No COVERALLS_TOKEN defined!$NOCOLOR";
        exit 1;
    fi
}

echo "Generating test coverage details..."

# Check if project being tested is a library or application
# (Libraries need to run 'go test cover' differently then Applications)
check_project_type
echo "  Projected type detected: $PROJ_TYPE"

if [ $PROJ_TYPE = "APP" ]; then
    # Set variables that differ depending on whether project type
    COVERPKG_OPTIONS="-coverpkg=app/..."
    GOPATH="$PWD/vendor:$PWD"

    # Call generate_cover_data() passing it a list of all packages being evaluated
    generate_coverage_data $(GOPATH=`pwd` go list ./... | grep -v vendor)
elif [ $PROJ_TYPE = "LIB" ]; then
    # Use custom environment
    define_custom_environment

    # Update GOPATH for running "go test" and "go tool cover"
    GOPATH="$GOPATH:$PWD/vendor"

    # Call generate_coverage_data
    generate_coverage_data $(go list ./... | grep -v vendor)
fi

# Output cover profile details for each function
show_cover_report func

# Finally, do additional things as needed
case "$1" in
"")
    ;;
--html)
    show_cover_report html ;;
--coveralls)
    push_to_coveralls ;;
*)
    echo >&2 "error: invalid option: $1"; exit 1 ;;
esac
