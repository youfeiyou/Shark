// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: msg.proto

// protoc --proto_path=../conn  -I. --go_out=. --go_opt=paths=source_relative  --go-grpc_out=. --go-grpc_opt=paths=source_relative msg.proto

package msg

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

type MessageType int32

const (
	MessageType_INVALID_MSG_TYPE      MessageType = 0
	MessageType_PERSONAL_MESSAGE_TYPE MessageType = 1
	MessageType_GROUP_MESSAGE_TYPE    MessageType = 2
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "INVALID_MSG_TYPE",
		1: "PERSONAL_MESSAGE_TYPE",
		2: "GROUP_MESSAGE_TYPE",
	}
	MessageType_value = map[string]int32{
		"INVALID_MSG_TYPE":      0,
		"PERSONAL_MESSAGE_TYPE": 1,
		"GROUP_MESSAGE_TYPE":    2,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seq      *SessionSeq `protobuf:"bytes,1,opt,name=seq,proto3" json:"seq,omitempty"`
	IsDelete uint32      `protobuf:"varint,100,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *Session) GetSeq() *SessionSeq {
	if x != nil {
		return x.Seq
	}
	return nil
}

func (x *Session) GetIsDelete() uint32 {
	if x != nil {
		return x.IsDelete
	}
	return 0
}

type AppFetchSeq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IosSeq uint64 `protobuf:"varint,1,opt,name=ios_seq,json=iosSeq,proto3" json:"ios_seq,omitempty"`
	AndSeq uint64 `protobuf:"varint,2,opt,name=and_seq,json=andSeq,proto3" json:"and_seq,omitempty"`
	PcSeq  uint64 `protobuf:"varint,3,opt,name=pc_seq,json=pcSeq,proto3" json:"pc_seq,omitempty"`
}

func (x *AppFetchSeq) Reset() {
	*x = AppFetchSeq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppFetchSeq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppFetchSeq) ProtoMessage() {}

func (x *AppFetchSeq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppFetchSeq.ProtoReflect.Descriptor instead.
func (*AppFetchSeq) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *AppFetchSeq) GetIosSeq() uint64 {
	if x != nil {
		return x.IosSeq
	}
	return 0
}

func (x *AppFetchSeq) GetAndSeq() uint64 {
	if x != nil {
		return x.AndSeq
	}
	return 0
}

func (x *AppFetchSeq) GetPcSeq() uint64 {
	if x != nil {
		return x.PcSeq
	}
	return 0
}

type SessionMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgReadSeq uint64       `protobuf:"varint,1,opt,name=msg_read_seq,json=msgReadSeq,proto3" json:"msg_read_seq,omitempty"`
	ClientSeq  *AppFetchSeq `protobuf:"bytes,2,opt,name=client_seq,json=clientSeq,proto3" json:"client_seq,omitempty"`
}

func (x *SessionMember) Reset() {
	*x = SessionMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionMember) ProtoMessage() {}

func (x *SessionMember) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionMember.ProtoReflect.Descriptor instead.
func (*SessionMember) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *SessionMember) GetMsgReadSeq() uint64 {
	if x != nil {
		return x.MsgReadSeq
	}
	return 0
}

func (x *SessionMember) GetClientSeq() *AppFetchSeq {
	if x != nil {
		return x.ClientSeq
	}
	return nil
}

type SessionSeq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxMsgSeq   uint64 `protobuf:"varint,1,opt,name=max_msg_seq,json=maxMsgSeq,proto3" json:"max_msg_seq,omitempty"`
	MaxMsgAtSeq uint64 `protobuf:"varint,2,opt,name=max_msg_at_seq,json=maxMsgAtSeq,proto3" json:"max_msg_at_seq,omitempty"`
	PacketSeq   uint64 `protobuf:"varint,3,opt,name=packet_seq,json=packetSeq,proto3" json:"packet_seq,omitempty"`
}

func (x *SessionSeq) Reset() {
	*x = SessionSeq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionSeq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionSeq) ProtoMessage() {}

func (x *SessionSeq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionSeq.ProtoReflect.Descriptor instead.
func (*SessionSeq) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *SessionSeq) GetMaxMsgSeq() uint64 {
	if x != nil {
		return x.MaxMsgSeq
	}
	return 0
}

func (x *SessionSeq) GetMaxMsgAtSeq() uint64 {
	if x != nil {
		return x.MaxMsgAtSeq
	}
	return 0
}

func (x *SessionSeq) GetPacketSeq() uint64 {
	if x != nil {
		return x.PacketSeq
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uin       uint64      `protobuf:"varint,1,opt,name=uin,proto3" json:"uin,omitempty"`
	Type      MessageType `protobuf:"varint,2,opt,name=type,proto3,enum=msg.MessageType" json:"type,omitempty"`
	Content   []byte      `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp uint64      `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SessionId string      `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	Seq       *SessionSeq `protobuf:"bytes,100,opt,name=seq,proto3" json:"seq,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetUin() uint64 {
	if x != nil {
		return x.Uin
	}
	return 0
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_INVALID_MSG_TYPE
}

func (x *Message) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Message) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Message) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *Message) GetSeq() *SessionSeq {
	if x != nil {
		return x.Seq
	}
	return nil
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x49, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x03, 0x73,
	0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x71, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x56, 0x0a, 0x0b, 0x41,
	0x70, 0x70, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6f,
	0x73, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x69, 0x6f, 0x73,
	0x53, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x71, 0x12, 0x15, 0x0a, 0x06,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x70, 0x63,
	0x53, 0x65, 0x71, 0x22, 0x62, 0x0a, 0x0d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x73, 0x67, 0x5f, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x52,
	0x65, 0x61, 0x64, 0x53, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x41, 0x70, 0x70, 0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x71, 0x52, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x71, 0x22, 0x70, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x71, 0x12, 0x1e, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x73, 0x67,
	0x5f, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x4d,
	0x73, 0x67, 0x53, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x73, 0x67,
	0x5f, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d,
	0x61, 0x78, 0x4d, 0x73, 0x67, 0x41, 0x74, 0x53, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x71, 0x22, 0xbb, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x75, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x71, 0x52, 0x03, 0x73, 0x65, 0x71, 0x2a, 0x56, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15,
	0x50, 0x45, 0x52, 0x53, 0x4f, 0x4e, 0x41, 0x4c, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x52, 0x4f, 0x55, 0x50,
	0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x02, 0x42,
	0x15, 0x5a, 0x13, 0x73, 0x68, 0x61, 0x72, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_msg_proto_goTypes = []interface{}{
	(MessageType)(0),      // 0: msg.MessageType
	(*Session)(nil),       // 1: msg.Session
	(*AppFetchSeq)(nil),   // 2: msg.AppFetchSeq
	(*SessionMember)(nil), // 3: msg.SessionMember
	(*SessionSeq)(nil),    // 4: msg.SessionSeq
	(*Message)(nil),       // 5: msg.Message
}
var file_msg_proto_depIdxs = []int32{
	4, // 0: msg.Session.seq:type_name -> msg.SessionSeq
	2, // 1: msg.SessionMember.client_seq:type_name -> msg.AppFetchSeq
	0, // 2: msg.Message.type:type_name -> msg.MessageType
	4, // 3: msg.Message.seq:type_name -> msg.SessionSeq
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppFetchSeq); i {
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
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionMember); i {
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
		file_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionSeq); i {
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
		file_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
