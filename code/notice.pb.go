// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: proto/notice.proto

package code

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ServerReq) Reset() {
	*x = ServerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerReq) ProtoMessage() {}

func (x *ServerReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerReq.ProtoReflect.Descriptor instead.
func (*ServerReq) Descriptor() ([]byte, []int) {
	return file_proto_notice_proto_rawDescGZIP(), []int{0}
}

func (x *ServerReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ClientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *ServerReq           `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Id       string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Metadata map[string]*Metadata `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 元数据
}

func (x *ClientReq) Reset() {
	*x = ClientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReq) ProtoMessage() {}

func (x *ClientReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReq.ProtoReflect.Descriptor instead.
func (*ClientReq) Descriptor() ([]byte, []int) {
	return file_proto_notice_proto_rawDescGZIP(), []int{1}
}

func (x *ClientReq) GetServer() *ServerReq {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *ClientReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ClientReq) GetMetadata() map[string]*Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Metadata_Int
	//	*Metadata_Uint
	//	*Metadata_Float
	//	*Metadata_String_
	//	*Metadata_Bool
	Value isMetadata_Value `protobuf_oneof:"value"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_proto_notice_proto_rawDescGZIP(), []int{2}
}

func (m *Metadata) GetValue() isMetadata_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Metadata) GetInt() int64 {
	if x, ok := x.GetValue().(*Metadata_Int); ok {
		return x.Int
	}
	return 0
}

func (x *Metadata) GetUint() uint64 {
	if x, ok := x.GetValue().(*Metadata_Uint); ok {
		return x.Uint
	}
	return 0
}

func (x *Metadata) GetFloat() float64 {
	if x, ok := x.GetValue().(*Metadata_Float); ok {
		return x.Float
	}
	return 0
}

func (x *Metadata) GetString_() string {
	if x, ok := x.GetValue().(*Metadata_String_); ok {
		return x.String_
	}
	return ""
}

func (x *Metadata) GetBool() bool {
	if x, ok := x.GetValue().(*Metadata_Bool); ok {
		return x.Bool
	}
	return false
}

type isMetadata_Value interface {
	isMetadata_Value()
}

type Metadata_Int struct {
	Int int64 `protobuf:"varint,1,opt,name=int,proto3,oneof"`
}

type Metadata_Uint struct {
	Uint uint64 `protobuf:"varint,2,opt,name=uint,proto3,oneof"`
}

type Metadata_Float struct {
	Float float64 `protobuf:"fixed64,3,opt,name=float,proto3,oneof"`
}

type Metadata_String_ struct {
	String_ string `protobuf:"bytes,4,opt,name=string,proto3,oneof"`
}

type Metadata_Bool struct {
	Bool bool `protobuf:"varint,5,opt,name=bool,proto3,oneof"`
}

func (*Metadata_Int) isMetadata_Value() {}

func (*Metadata_Uint) isMetadata_Value() {}

func (*Metadata_Float) isMetadata_Value() {}

func (*Metadata_String_) isMetadata_Value() {}

func (*Metadata_Bool) isMetadata_Value() {}

type SendReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server    *ServerReq `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Message   []byte     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`     // 发送消息
	IdList    []string   `protobuf:"bytes,3,rep,name=idList,proto3" json:"idList,omitempty"`       // 发送id集合
	Condition []byte     `protobuf:"bytes,4,opt,name=condition,proto3" json:"condition,omitempty"` // 条件
}

func (x *SendReq) Reset() {
	*x = SendReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReq) ProtoMessage() {}

func (x *SendReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReq.ProtoReflect.Descriptor instead.
func (*SendReq) Descriptor() ([]byte, []int) {
	return file_proto_notice_proto_rawDescGZIP(), []int{3}
}

func (x *SendReq) GetServer() *ServerReq {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *SendReq) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SendReq) GetIdList() []string {
	if x != nil {
		return x.IdList
	}
	return nil
}

func (x *SendReq) GetCondition() []byte {
	if x != nil {
		return x.Condition
	}
	return nil
}

type RecvResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID  string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	Message   []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Heartbeat bool   `protobuf:"varint,3,opt,name=heartbeat,proto3" json:"heartbeat,omitempty"`
}

func (x *RecvResp) Reset() {
	*x = RecvResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_notice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecvResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecvResp) ProtoMessage() {}

func (x *RecvResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_notice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecvResp.ProtoReflect.Descriptor instead.
func (*RecvResp) Descriptor() ([]byte, []int) {
	return file_proto_notice_proto_rawDescGZIP(), []int{4}
}

func (x *RecvResp) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *RecvResp) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *RecvResp) GetHeartbeat() bool {
	if x != nil {
		return x.Heartbeat
	}
	return false
}

var File_proto_notice_proto protoreflect.FileDescriptor

var file_proto_notice_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xcc, 0x01, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x4b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x85, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x12, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x03, 0x69, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x75, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x48, 0x00, 0x52, 0x04, 0x75, 0x69, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x05, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f,
	0x6f, 0x6c, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x07,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x5e, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x76, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x32, 0x9b, 0x02, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b, 0x52, 0x65,
	0x63, 0x76, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x63, 0x76, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_notice_proto_rawDescOnce sync.Once
	file_proto_notice_proto_rawDescData = file_proto_notice_proto_rawDesc
)

func file_proto_notice_proto_rawDescGZIP() []byte {
	file_proto_notice_proto_rawDescOnce.Do(func() {
		file_proto_notice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_notice_proto_rawDescData)
	})
	return file_proto_notice_proto_rawDescData
}

var file_proto_notice_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_notice_proto_goTypes = []any{
	(*ServerReq)(nil),     // 0: grpc.ServerReq
	(*ClientReq)(nil),     // 1: grpc.ClientReq
	(*Metadata)(nil),      // 2: grpc.Metadata
	(*SendReq)(nil),       // 3: grpc.SendReq
	(*RecvResp)(nil),      // 4: grpc.RecvResp
	nil,                   // 5: grpc.ClientReq.MetadataEntry
	(*emptypb.Empty)(nil), // 6: google.protobuf.Empty
}
var file_proto_notice_proto_depIdxs = []int32{
	0, // 0: grpc.ClientReq.server:type_name -> grpc.ServerReq
	5, // 1: grpc.ClientReq.metadata:type_name -> grpc.ClientReq.MetadataEntry
	0, // 2: grpc.SendReq.server:type_name -> grpc.ServerReq
	2, // 3: grpc.ClientReq.MetadataEntry.value:type_name -> grpc.Metadata
	0, // 4: grpc.Notice.Register:input_type -> grpc.ServerReq
	1, // 5: grpc.Notice.AddClient:input_type -> grpc.ClientReq
	1, // 6: grpc.Notice.DelClient:input_type -> grpc.ClientReq
	3, // 7: grpc.Notice.SendMessage:input_type -> grpc.SendReq
	0, // 8: grpc.Notice.RecvMessage:input_type -> grpc.ServerReq
	6, // 9: grpc.Notice.Register:output_type -> google.protobuf.Empty
	6, // 10: grpc.Notice.AddClient:output_type -> google.protobuf.Empty
	6, // 11: grpc.Notice.DelClient:output_type -> google.protobuf.Empty
	6, // 12: grpc.Notice.SendMessage:output_type -> google.protobuf.Empty
	4, // 13: grpc.Notice.RecvMessage:output_type -> grpc.RecvResp
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_notice_proto_init() }
func file_proto_notice_proto_init() {
	if File_proto_notice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_notice_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ServerReq); i {
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
		file_proto_notice_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ClientReq); i {
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
		file_proto_notice_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Metadata); i {
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
		file_proto_notice_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SendReq); i {
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
		file_proto_notice_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*RecvResp); i {
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
	file_proto_notice_proto_msgTypes[2].OneofWrappers = []any{
		(*Metadata_Int)(nil),
		(*Metadata_Uint)(nil),
		(*Metadata_Float)(nil),
		(*Metadata_String_)(nil),
		(*Metadata_Bool)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_notice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_notice_proto_goTypes,
		DependencyIndexes: file_proto_notice_proto_depIdxs,
		MessageInfos:      file_proto_notice_proto_msgTypes,
	}.Build()
	File_proto_notice_proto = out.File
	file_proto_notice_proto_rawDesc = nil
	file_proto_notice_proto_goTypes = nil
	file_proto_notice_proto_depIdxs = nil
}
