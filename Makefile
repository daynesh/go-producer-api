export GOPATH:=$(CURDIR)/vendor:$(CURDIR)

ifndef JOB_NAME
JOB_NAME := $(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))
endif
APPNAME := $(JOB_NAME)
BUILD_DIR := $(APPNAME)-$(BUILD_NUMBER)
BUILD_TARBALL := $(APPNAME)-$(BUILD_NUMBER).tar.bz2

default: install build

.PHONY: clean
clean:
	rm -rf vendor .glide
	rm -f $(APPNAME)

.PHONY: build
build:
	@go build -o $(APPNAME) src/app/main.go
	@go build -o $(APPNAME) -ldflags "-X app/bootstrap.BuildNumber=$(BUILD_NUMBER)" src/app/main.go

.PHONY: install
install: clean glide
	@env GOPATH=`pwd` .glide/bin/glide install

	@# The following ensures that imports will resolve packages located in ./vendor/src
	@mv vendor vendor.tmp
	@mkdir vendor
	@mv vendor.tmp vendor/src

.PHONY: glide
glide:
	@env GOPATH=`pwd`/.glide go get github.com/Masterminds/glide

.PHONY: tests
tests:
	@go test ./src/test/... -v

.PHONY: testcoverage
testcoverage:
	./scripts/generate-test-coverage.sh

.PHONY: lintchecks
lintchecks:
	./scripts/lintchecks.sh

.PHONY: dist
dist: install build
	rm -f build/$(BUILD_TARBALL)
	mkdir -p build/$(BUILD_DIR) && cp $(APPNAME) build/$(BUILD_DIR)/
	tar -C ./build -cjf build/$(BUILD_TARBALL) $(BUILD_DIR)

.SILENT: clean
