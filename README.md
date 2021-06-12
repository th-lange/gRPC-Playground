# Communications Module

The idea was to play around with gRPC and process communication.
So this project does not follow best coding practices!

It was just to see, how gRPC fares with a small go project to play with
inner process communication in kubernetes.

## Required:
- go1.16.5 or higher
- go get google.golang.org/protobuf/reflect/protoreflect@v1.26.0
- go get google.golang.org/protobuf/runtime/protoimpl@v1.26.0
- go get google.golang.org/protobuf/reflect/protoreflect@v1.26.0
- https://grpc.io/docs/protoc-installation/


## To Recompile the .proto files:

Example:
```shell
protoc --go_out=plugins=grpc:. data/message.proto
```


## How to use

- Build Images, tag and push
    - or use: thomaslangesmfhq/com-package:0.0.1
- Adjust the ingress file: `kubernetes/ingress.yaml`
  - or remove it if not needed
- kubectl apply -f kubernetes