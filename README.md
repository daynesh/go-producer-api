[![build status](https://travis-ci.org/daynesh/go-producer-api.svg?branch=master)](https://travis-ci.org/daynesh/go-producer-api.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/daynesh/go-producer-api)](https://goreportcard.com/report/github.com/daynesh/go-producer-api)
[![Codecov Status](https://codecov.io/gh/daynesh/go-producer-api/branch/master/graph/badge.svg)](https://codecov.io/gh/daynesh/go-producer-api)
[![Coveralls Status](https://coveralls.io/repos/github/daynesh/go-producer-api/badge.svg?branch=master)](https://coveralls.io/github/daynesh/go-producer-api?branch=master)


# Producer API
This application serves as an API allowing clients to publish data through a pub/sub messaging queue.  It utilizes the following technologies:

<a href="https://golang.org/" target="_blank" title="Go">
  <img height="50" src="https://upload.wikimedia.org/wikipedia/commons/2/23/Golang.png"/>
</a>
<a href="https://kafka.apache.org/" target="_blank" title="Kafka">
  <img height="50" src="https://kafka.apache.org/images/logo.png"/>
</a>

## Dependencies
Building and running the application requires the following:
- An installation of `go`
- An appropriately defined `$GOPATH` environment variable as per [these instructions](https://golang.org/doc/code.html#GOPATH)
- A directory hierarchy under `$GOPATH` consistent with [these instructions](https://golang.org/doc/code.html#Workspaces)

## Build Instructions
The following step(s) will install all application dependencies and create the `go-producer-api` executable:
```
$ make
```

## Runtime Instructions
To run the application after its been built, simply execute the build artifact produced from the build step:
```
$ ./go-producer-api
```
