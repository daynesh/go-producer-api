# Producer API
This application serves as an API allowing clients to publish data through a pub/sub messaging queue.  It utilizes the following technologies:

<a href="https://golang.org/" target="_blank" title="Go">
  <img height="50" src="https://upload.wikimedia.org/wikipedia/commons/2/23/Golang.png"/>
</a>
<a href="https://kafka.apache.org/" target="_blank" title="Kafka">
  <img height="50" src="https://kafka.apache.org/images/logo.png"/>
</a>

## Build Instructions
Building the application requires `go` and the appropriately defined `$GOPATH` env variable with its appropriate directory structure. See [here](https://golang.org/doc/code.html#GOPATH) for more details.

Assuming the above prerequisites exists, the following will install all dependencies and create the `go-producer-api` executable:
```
$ make build
```

## Runtime Instructions
To run the application after its been built, simply execute the build artifact produced from the build step:
```
$ ./go-producer-api
```
