GLIDE_EXE=$(GOPATH)/bin/glide

default: build

clean:
	rm -rf vendor
	rm -f go-producer-api
	rm -f $(GOPATH)/bin/glide

build: install
	go build

install: glide
	$(GLIDE_EXE) install

glide:
	go get github.com/Masterminds/glide

.PHONY: clean build install glide