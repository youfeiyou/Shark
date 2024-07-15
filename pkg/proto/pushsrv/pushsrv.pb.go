// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: pushsrv.proto

// protoc --proto_path=../msg  -I. --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative pushsrv.proto

package pushsrv

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	msg "shark/pkg/proto/msg"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId string       `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Uins      []uint64     `protobuf:"varint,2,rep,packed,name=uins,proto3" json:"uins,omitempty"`
	Msg       *msg.Message `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pushsrv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_pushsrv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_pushsrv_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Req) GetUins() []uint64 {
	if x != nil {
		return x.Uins
	}
	return nil
}

func (x *Req) GetMsg() *msg.Message {
	if x != nil {
		return x.Msg
	}
	return nil
}

type Resp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin  uint64 `protobuf:"varint,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Code uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Resp) Reset() {
	*x = Resp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pushsrv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resp) ProtoMessage() {}

func (x *Resp) ProtoReflect() protoreflect.Message {
	mi := &file_pushsrv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resp.ProtoReflect.Descriptor instead.
func (*Resp) Descriptor() ([]byte, []int) {
	return file_pushsrv_proto_rawDescGZIP(), []int{1}
}

func (x *Resp) GetUin() uint64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *Resp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Rsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res []*Resp `protobuf:"bytes,1,rep,name=res,proto3" json:"res,omitempty"`
}

func (x *Rsp) Reset() {
	*x = Rsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pushsrv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rsp) ProtoMessage() {}

func (x *Rsp) ProtoReflect() protoreflect.Message {
	mi := &file_pushsrv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rsp.ProtoReflect.Descriptor instead.
func (*Rsp) Descriptor() ([]byte, []int) {
	return file_pushsrv_proto_rawDescGZIP(), []int{2}
}

func (x *Rsp) GetRes() []*Resp {
	if x != nil {
		return x.Res
	}
	return nil
}

var File_pushsrv_proto protoreflect.FileDescriptor

var file_pushsrv_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x75, 0x73, 0x68, 0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x75, 0x73, 0x68, 0x73, 0x72, 0x76, 0x1a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x69, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x04, 0x75, 0x69, 0x6e, 0x73, 0x12, 0x1e, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2c, 0x0a,
	0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x26, 0x0a, 0x03, 0x52,
	0x73, 0x70, 0x12, 0x1f, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x73, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x03,
	0x72, 0x65, 0x73, 0x32, 0x34, 0x0a, 0x0e, 0x50, 0x75, 0x73, 0x68, 0x53, 0x72, 0x76, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x75, 0x73, 0x68, 0x12, 0x0c, 0x2e,
	0x70, 0x75, 0x73, 0x68, 0x73, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x75,
	0x73, 0x68, 0x73, 0x72, 0x76, 0x2e, 0x52, 0x73, 0x70, 0x42, 0x19, 0x5a, 0x17, 0x73, 0x68, 0x61,
	0x72, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x75, 0x73,
	0x68, 0x73, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pushsrv_proto_rawDescOnce sync.Once
	file_pushsrv_proto_rawDescData = file_pushsrv_proto_rawDesc
)

func file_pushsrv_proto_rawDescGZIP() []byte {
	file_pushsrv_proto_rawDescOnce.Do(func() {
		file_pushsrv_proto_rawDescData = protoimpl.X.CompressGZIP(file_pushsrv_proto_rawDescData)
	})
	return file_pushsrv_proto_rawDescData
}

var file_pushsrv_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pushsrv_proto_goTypes = []interface{}{
	(*Req)(nil),         // 0: pushsrv.Req
	(*Resp)(nil),        // 1: pushsrv.Resp
	(*Rsp)(nil),         // 2: pushsrv.Rsp
	(*msg.Message)(nil), // 3: msg.Message
}
var file_pushsrv_proto_depIdxs = []int32{
	3, // 0: pushsrv.Req.msg:type_name -> msg.Message
	1, // 1: pushsrv.Rsp.res:type_name -> pushsrv.Resp
	0, // 2: pushsrv.PushSrvService.push:input_type -> pushsrv.Req
	2, // 3: pushsrv.PushSrvService.push:output_type -> pushsrv.Rsp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pushsrv_proto_init() }
func file_pushsrv_proto_init() {
	if File_pushsrv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pushsrv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pushsrv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pushsrv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pushsrv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pushsrv_proto_goTypes,
		DependencyIndexes: file_pushsrv_proto_depIdxs,
		MessageInfos:      file_pushsrv_proto_msgTypes,
	}.Build()
	File_pushsrv_proto = out.File
	file_pushsrv_proto_rawDesc = nil
	file_pushsrv_proto_goTypes = nil
	file_pushsrv_proto_depIdxs = nil
}
