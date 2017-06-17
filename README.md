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
$ make build
```

## Runtime Instructions
To run the application after its been built, simply execute the build artifact produced from the build step:
```
$ ./go-producer-api
```
