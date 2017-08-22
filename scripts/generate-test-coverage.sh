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

workdir=.cover
profile="$workdir/cover.out"
mode=count

generate_cover_data() {
    rm -rf "$workdir"
    mkdir "$workdir"

    for pkg in "$@"; do
        f="$workdir/$(echo $pkg | tr / -).cover"
        env GOPATH=`pwd`/vendor:`pwd` go test -covermode="$mode" -coverprofile="$f" -coverpkg=app/... "$pkg"
    done

    echo "mode: $mode" >"$profile"
    grep -h -v "^mode:" "$workdir"/*.cover >>"$profile"
}

show_cover_report() {
    env GOPATH=`pwd` go tool cover -${1}="$profile"
}

push_to_coveralls() {
    env GOPATH=`pwd`/vendor go get github.com/mattn/goveralls
    echo "Pushing coverage statistics to coveralls.io"
    env GOPATH=`pwd` vendor/bin/goveralls -coverprofile="$profile"
}

generate_cover_data $(env GOPATH=`pwd` go list ./src/...)
show_cover_report func
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
