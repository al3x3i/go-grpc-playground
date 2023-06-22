# Description
Build simple Go application client/server which communicates each other by gRPC(Google Remote Procedure Call)

# Generate Go code from .proto file
## 1
protoc --go_out=plugins=grpc:. ./protos/calculator.proto
## 2
protoc --go_out=plugins=grpc:. -I=./protos ./protos/calculator.proto

