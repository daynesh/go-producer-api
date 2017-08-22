export GOPATH:=$(CURDIR)/vendor:$(CURDIR)

GLIDE_EXE := $(GOPATH)/bin/glide
APPNAME := $(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))

default: install build

.PHONY: clean
clean:
	rm -rf vendor
	rm -f $(APPNAME)

.PHONY: build
build:
	@go build -o $(APPNAME) src/app/main.go

.PHONY: install
install: clean glide
	@env GOPATH=`pwd` vendor/bin/glide install

	@# The following ensures that imports will resolve packages located in ./vendor/src
	@mv vendor vendor.tmp
	@mkdir vendor
	@mv vendor.tmp vendor/src

.PHONY: glide
glide:
	@env GOPATH=`pwd`/vendor go get github.com/Masterminds/glide

.PHONY: tests
tests:
	@go test ./src/test/... -v

.PHONY: testcoverage
testcoverage:
	./scripts/generate-test-coverage.sh

.PHONY: lintchecks
lintchecks:
	./scripts/travis.gofmt.sh
	./scripts/travis.govet.sh
	./scripts/travis.gometalint.sh

.SILENT: clean