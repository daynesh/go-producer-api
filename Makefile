ifeq ("$(GOPATH)","")
$(error GOPATH must be set)
endif

GLIDE_EXE := ${GOPATH}/bin/glide
GOREPO := ${GOPATH}/src/github.com/daynesh/go-producer-api
APPNAME := $(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))

default: install build

.PHONY: clean
clean:
	rm -rf vendor
	rm -f ${APPNAME}
	rm -f ${GOPATH}/bin/glide
	rm -f ${GOPATH}/bin/gometalinter.v1

.PHONY: build
build:
	@cd ${GOREPO}/src && go build -o ${GOREPO}/$(APPNAME) main.go

.PHONY: install
install: clean glide
	$(GLIDE_EXE) install

.PHONY: glide
glide:
	go get github.com/Masterminds/glide

.PHONY: tests
tests:
	go test ./test/... -v

.PHONY: testcoverage
testcoverage:
	./scripts/generate-test-coverage.sh

.SILENT: clean