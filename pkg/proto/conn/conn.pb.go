// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: conn.proto

// protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative conn.proto

package conn

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

type PushReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PushMessages []*msg.Message `protobuf:"bytes,1,rep,name=push_messages,json=pushMessages,proto3" json:"push_messages,omitempty"`
}

func (x *PushReq) Reset() {
	*x = PushReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushReq) ProtoMessage() {}

func (x *PushReq) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushReq.ProtoReflect.Descriptor instead.
func (*PushReq) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0}
}

func (x *PushReq) GetPushMessages() []*msg.Message {
	if x != nil {
		return x.PushMessages
	}
	return nil
}

type PushRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailedUin []uint64 `protobuf:"varint,1,rep,packed,name=failed_uin,json=failedUin,proto3" json:"failed_uin,omitempty"`
}

func (x *PushRsp) Reset() {
	*x = PushRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRsp) ProtoMessage() {}

func (x *PushRsp) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRsp.ProtoReflect.Descriptor instead.
func (*PushRsp) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{1}
}

func (x *PushRsp) GetFailedUin() []uint64 {
	if x != nil {
		return x.FailedUin
	}
	return nil
}

type HeatBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin       uint64 `protobuf:"varint,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Timestamp uint64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Contents  []byte `protobuf:"bytes,3,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *HeatBeat) Reset() {
	*x = HeatBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeatBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeatBeat) ProtoMessage() {}

func (x *HeatBeat) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeatBeat.ProtoReflect.Descriptor instead.
func (*HeatBeat) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{2}
}

func (x *HeatBeat) GetUin() uint64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *HeatBeat) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *HeatBeat) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

type MemberStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin      uint64 `protobuf:"varint,1,opt,name=uin,proto3" json:"uin,omitempty"`
	ConnAddr string `protobuf:"bytes,2,opt,name=conn_addr,json=connAddr,proto3" json:"conn_addr,omitempty"`
	IsOnline uint32 `protobuf:"varint,3,opt,name=is_online,json=isOnline,proto3" json:"is_online,omitempty"`
}

func (x *MemberStat) Reset() {
	*x = MemberStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberStat) ProtoMessage() {}

func (x *MemberStat) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberStat.ProtoReflect.Descriptor instead.
func (*MemberStat) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{3}
}

func (x *MemberStat) GetUin() uint64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *MemberStat) GetConnAddr() string {
	if x != nil {
		return x.ConnAddr
	}
	return ""
}

func (x *MemberStat) GetIsOnline() uint32 {
	if x != nil {
		return x.IsOnline
	}
	return 0
}

var File_conn_proto protoreflect.FileDescriptor

var file_conn_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f,
	0x6e, 0x6e, 0x1a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a,
	0x07, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x31, 0x0a, 0x0d, 0x70, 0x75, 0x73, 0x68,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0c, 0x70,
	0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x07, 0x50,
	0x75, 0x73, 0x68, 0x52, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x5f, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x55, 0x69, 0x6e, 0x22, 0x56, 0x0a, 0x08, 0x48, 0x65, 0x61, 0x74, 0x42, 0x65, 0x61,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x75, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x58, 0x0a,
	0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69,
	0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x32, 0x33, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x75, 0x73, 0x68, 0x12, 0x0d,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e,
	0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x73, 0x70, 0x42, 0x16, 0x5a, 0x14,
	0x73, 0x68, 0x61, 0x72, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6e, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conn_proto_rawDescOnce sync.Once
	file_conn_proto_rawDescData = file_conn_proto_rawDesc
)

func file_conn_proto_rawDescGZIP() []byte {
	file_conn_proto_rawDescOnce.Do(func() {
		file_conn_proto_rawDescData = protoimpl.X.CompressGZIP(file_conn_proto_rawDescData)
	})
	return file_conn_proto_rawDescData
}

var file_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conn_proto_goTypes = []interface{}{
	(*PushReq)(nil),     // 0: conn.PushReq
	(*PushRsp)(nil),     // 1: conn.PushRsp
	(*HeatBeat)(nil),    // 2: conn.HeatBeat
	(*MemberStat)(nil),  // 3: conn.MemberStat
	(*msg.Message)(nil), // 4: msg.Message
}
var file_conn_proto_depIdxs = []int32{
	4, // 0: conn.PushReq.push_messages:type_name -> msg.Message
	0, // 1: conn.ConnService.push:input_type -> conn.PushReq
	1, // 2: conn.ConnService.push:output_type -> conn.PushRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_conn_proto_init() }
func file_conn_proto_init() {
	if File_conn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushReq); i {
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
		file_conn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRsp); i {
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
		file_conn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeatBeat); i {
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
		file_conn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberStat); i {
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
			RawDescriptor: file_conn_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conn_proto_goTypes,
		DependencyIndexes: file_conn_proto_depIdxs,
		MessageInfos:      file_conn_proto_msgTypes,
	}.Build()
	File_conn_proto = out.File
	file_conn_proto_rawDesc = nil
	file_conn_proto_goTypes = nil
	file_conn_proto_depIdxs = nil
}
