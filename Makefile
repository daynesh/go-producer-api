GLIDE_EXE=$(GOPATH)/bin/glide

# Identify current project name
APPNAME=$(notdir $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST))))))

default: install build

clean:
	rm -rf vendor
	rm -f $(APPNAME)
	rm -f $(GOPATH)/bin/glide

build:
	go build -o $(APPNAME) src/*.go

install: clean glide
	$(GLIDE_EXE) install

glide:
	go get github.com/Masterminds/glide

.PHONY: clean build install glide
.SILENT: clean