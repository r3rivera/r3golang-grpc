#!/bin/bash

echo "Generating the necessary code from the .proto files..."
protoc -I=. --go_out=. greet/greetpb/*.proto

echo "Complete generating the coded from the .proto files..."