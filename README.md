Golang RPC Application

#Quick start
https://grpc.io/docs/quickstart/go/ 

Steps to install the proto buffer code generator

1. Install the PROTOC
https://github.com/protocolbuffers/protobuf/releases/tag/v3.11.4 

2. Install the GO Plugin, this will be installed in $GOPATH Directory with a file protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go 

3. Add the executable protoc-gen-go to your $PATH so that protoc see can see it
export PATH="$PATH:$(go env GOPATH)/bin"

4. Generating the code use the following:
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/greet.proto 

 EX: protoc -I=. --go_out=. greet/greetpb/greet.proto  ---- I used the . to create the greet.pb.go in the same directory as my .proto

  Note: SRC_DIR is the source directory
        DST_DIR is the destination directory

#Visual Code Plugin for protocol buffer syntax highlight
https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3


#Golang Dependencies
go get -u google.golang.org/grpc

