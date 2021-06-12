package service

//import (
//	proto "github.com/golang/protobuf/proto"
//	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
//	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
//	reflect "reflect"
//	sync "sync"
//)
//
//
//type HelloRequest struct {
//	state         protoimpl.MessageState
//	sizeCache     protoimpl.SizeCache
//	unknownFields protoimpl.UnknownFields
//
//	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
//}
//
//func (x *HelloRequest) Reset() {
//	*x = HelloRequest{}
//	if protoimpl.UnsafeEnabled {
//		mi := &file_examples_helloworld_helloworld_helloworld_proto_msgTypes[0]
//		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
//		ms.StoreMessageInfo(mi)
//	}
//}