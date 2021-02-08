To generate this proto file as functional go program, make sure you have protoc generator and add them to environment path.

To get it execute this command

~~go get -u github.com/golang/protobuf/protoc-gen-go~~
    
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

If somehow the command above still not work, try to install the package by executing

~~go install github.com/golang/protobuf/protoc-gen-go~~
    
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

(If its still not work, try with the protoc.exe in this directory)

After that, include the binary path to the environment path (if the generated binary is protoc-gen-go, 
we can rename this file to protoc for simpler name)

To generate go file from this proto execute command like this

    protoc --go_out=plugins=grpc:chat chat.proto

To automatically generate all .proto file within working directory, replace 

    "grpc:chat" to "grpc:."
and 

    "chat.proto" to "*" or "*.proto"

incase of building with google.protobuf type like Timestamp, use include parameter to the command so it would be like this

    protoc -I=. -I="D:\GOPATH\protobuf\src" --go_out=plugins=grpc:. *.proto    

the protobuf\src is the location of src folder of protobuf repository (download it via https://github.com/protocolbuffers/protobuf) 