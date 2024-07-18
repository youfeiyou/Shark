// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: groupmember.proto

// protoc --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative groupmember.proto

package groupmember

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DbGroupMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group    uint64 `protobuf:"varint,1,opt,name=group,proto3" json:"group,omitempty"`
	Uin      uint64 `protobuf:"varint,2,opt,name=uin,proto3" json:"uin,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Comments string `protobuf:"bytes,4,opt,name=comments,proto3" json:"comments,omitempty"`
	Notify   uint32 `protobuf:"varint,5,opt,name=notify,proto3" json:"notify,omitempty"`
	Status   uint32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DbGroupMember) Reset() {
	*x = DbGroupMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupmember_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DbGroupMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DbGroupMember) ProtoMessage() {}

func (x *DbGroupMember) ProtoReflect() protoreflect.Message {
	mi := &file_groupmember_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DbGroupMember.ProtoReflect.Descriptor instead.
func (*DbGroupMember) Descriptor() ([]byte, []int) {
	return file_groupmember_proto_rawDescGZIP(), []int{0}
}

func (x *DbGroupMember) GetGroup() uint64 {
	if x != nil {
		return x.Group
	}
	return 0
}

func (x *DbGroupMember) GetUin() uint64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *DbGroupMember) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DbGroupMember) GetComments() string {
	if x != nil {
		return x.Comments
	}
	return ""
}

func (x *DbGroupMember) GetNotify() uint32 {
	if x != nil {
		return x.Notify
	}
	return 0
}

func (x *DbGroupMember) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemberNum uint64 `protobuf:"varint,1,opt,name=member_num,json=memberNum,proto3" json:"member_num,omitempty"`
	Index     uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	TableNum  uint32 `protobuf:"varint,3,opt,name=table_num,json=tableNum,proto3" json:"table_num,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groupmember_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_groupmember_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_groupmember_proto_rawDescGZIP(), []int{1}
}

func (x *Meta) GetMemberNum() uint64 {
	if x != nil {
		return x.MemberNum
	}
	return 0
}

func (x *Meta) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Meta) GetTableNum() uint32 {
	if x != nil {
		return x.TableNum
	}
	return 0
}

var File_groupmember_proto protoreflect.FileDescriptor

var file_groupmember_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x97, 0x01, 0x0a, 0x0d, 0x64, 0x62, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x58, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4e, 0x75,
	0x6d, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x75, 0x6d, 0x42, 0x1d, 0x5a, 0x1b, 0x73, 0x68, 0x61, 0x72, 0x6b, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groupmember_proto_rawDescOnce sync.Once
	file_groupmember_proto_rawDescData = file_groupmember_proto_rawDesc
)

func file_groupmember_proto_rawDescGZIP() []byte {
	file_groupmember_proto_rawDescOnce.Do(func() {
		file_groupmember_proto_rawDescData = protoimpl.X.CompressGZIP(file_groupmember_proto_rawDescData)
	})
	return file_groupmember_proto_rawDescData
}

var file_groupmember_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_groupmember_proto_goTypes = []interface{}{
	(*DbGroupMember)(nil), // 0: groupmember.dbGroupMember
	(*Meta)(nil),          // 1: groupmember.meta
}
var file_groupmember_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_groupmember_proto_init() }
func file_groupmember_proto_init() {
	if File_groupmember_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_groupmember_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DbGroupMember); i {
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
		file_groupmember_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
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
			RawDescriptor: file_groupmember_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_groupmember_proto_goTypes,
		DependencyIndexes: file_groupmember_proto_depIdxs,
		MessageInfos:      file_groupmember_proto_msgTypes,
	}.Build()
	File_groupmember_proto = out.File
	file_groupmember_proto_rawDesc = nil
	file_groupmember_proto_goTypes = nil
	file_groupmember_proto_depIdxs = nil
}