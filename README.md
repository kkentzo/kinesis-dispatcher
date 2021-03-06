[![Travis Build Status](https://travis-ci.org/e-travel/message-dispatcher.svg?branch=master)](https://travis-ci.org/e-travel/message-dispatcher)
[![Go Report
Card](https://goreportcard.com/badge/github.com/e-travel/message-dispatcher)](https://goreportcard.com/report/github.com/e-travel/message-dispatcher)

# Message Dispatcher

`message-dispatcher` is a simple message broker service that can
receive arbitrary messages through:

* unix datagram socket
* udp endpoint

and forward them to:

* AWS Kinesis
* InfluxDB (through the http API)

The service is fully agnostic as to the format and content of the
messages. Emphasis is given to low latency as far as the service's
client is concerned and ensuring non-blocking client requests.

## Instructions

### Server

Usage instructions: `./message-dispatcher -help`

### Clients

Ruby client examples can be found in directory `clients/`

### Development

0. Install `go`
1. Install the [go dep tool](https://github.com/golang/dep)
2. Download the code using `go get github.com/e-travel/message-dispatcher`
3. `dep ensure`
4. `go build -v -i`

### Tests and benchmarks

Tests and examples:

    $ go test ./... -v

Server Benchmarks

    $ go test -bench=. ./...

# License

This program is available as open source under the terms of the [MIT
License](http://opensource.org/licenses/MIT).
