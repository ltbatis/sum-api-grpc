# Sum API gRPC

This is a sample of an API using gRPC unary framework, that uses a Client that requests a result of a sum operation, and passes as arguments two integers. And a Server, that calculates the Sum and gives a response.
#
### To execute
First initialize the Server 
```
 go run sum/sum_server/server.go
```
After, execute the client passing two arguments as below
```
go run sum/sum_client/client.go 10 20
```
#
### Requirements
[![Stable release](https://img.shields.io/badge/libprotoc-3.17.3-green.svg)](https://github.com/protocolbuffers/protobuf/releases/tag/v3.17.3)
[![Stable release](https://img.shields.io/badge/golang-v1.17.5-blue.svg)](https://go.dev/doc/go1.17)
