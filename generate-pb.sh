#!/bin/bash

echo "Generating the necessary code from the .proto files..."
echo "Updating GO111MODULE..."
export GO111MODULE=on
echo "PROTOC Version..."
protoc --version
echo "Generating proto codes..."
protoc -I . --go_out=plugins=grpc:. ./greet/greetpb/*.proto
echo "Complete generating the code from the .proto files..."