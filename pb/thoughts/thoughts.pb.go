// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: thoughts/thoughts.proto

package thoughts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Thought struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created,proto3" json:"created,omitempty"`
	Text    string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Uuid    []byte                 `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *Thought) Reset() {
	*x = Thought{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thoughts_thoughts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thought) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thought) ProtoMessage() {}

func (x *Thought) ProtoReflect() protoreflect.Message {
	mi := &file_thoughts_thoughts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thought.ProtoReflect.Descriptor instead.
func (*Thought) Descriptor() ([]byte, []int) {
	return file_thoughts_thoughts_proto_rawDescGZIP(), []int{0}
}

func (x *Thought) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *Thought) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Thought) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

type ThoughtPad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thoughts []*Thought `protobuf:"bytes,1,rep,name=thoughts,proto3" json:"thoughts,omitempty"`
}

func (x *ThoughtPad) Reset() {
	*x = ThoughtPad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thoughts_thoughts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThoughtPad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThoughtPad) ProtoMessage() {}

func (x *ThoughtPad) ProtoReflect() protoreflect.Message {
	mi := &file_thoughts_thoughts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThoughtPad.ProtoReflect.Descriptor instead.
func (*ThoughtPad) Descriptor() ([]byte, []int) {
	return file_thoughts_thoughts_proto_rawDescGZIP(), []int{1}
}

func (x *ThoughtPad) GetThoughts() []*Thought {
	if x != nil {
		return x.Thoughts
	}
	return nil
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thought string `protobuf:"bytes,1,opt,name=thought,proto3" json:"thought,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thoughts_thoughts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thoughts_thoughts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_thoughts_thoughts_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetThought() string {
	if x != nil {
		return x.Thought
	}
	return ""
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid []byte `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thoughts_thoughts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thoughts_thoughts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_thoughts_thoughts_proto_rawDescGZIP(), []int{3}
}

func (x *AddResponse) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thoughts_thoughts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_thoughts_thoughts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_thoughts_thoughts_proto_rawDescGZIP(), []int{4}
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Thoughts []*Thought `protobuf:"bytes,1,rep,name=thoughts,proto3" json:"thoughts,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thoughts_thoughts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_thoughts_thoughts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_thoughts_thoughts_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetThoughts() []*Thought {
	if x != nil {
		return x.Thoughts
	}
	return nil
}

var File_thoughts_thoughts_proto protoreflect.FileDescriptor

var file_thoughts_thoughts_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x74, 0x68, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6d, 0x61, 0x63, 0x68, 0x73,
	0x68, 0x65, 0x76, 0x2e, 0x6d, 0x61, 0x73, 0x61, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x68, 0x6f, 0x75,
	0x67, 0x68, 0x74, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x07, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74,
	0x12, 0x34, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x4c,
	0x0a, 0x0a, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x50, 0x61, 0x64, 0x12, 0x3e, 0x0a, 0x08,
	0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6d, 0x61, 0x63, 0x68, 0x73, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x61, 0x73, 0x61, 0x2e, 0x70,
	0x62, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x54, 0x68, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x52, 0x08, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x22, 0x26, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x68,
	0x6f, 0x75, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x68, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x22, 0x21, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x73,
	0x68, 0x65, 0x76, 0x2e, 0x6d, 0x61, 0x73, 0x61, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x68, 0x6f, 0x75,
	0x67, 0x68, 0x74, 0x73, 0x2e, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x52, 0x08, 0x74, 0x68,
	0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x32, 0xc3, 0x01, 0x0a, 0x0e, 0x54, 0x68, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x03, 0x41, 0x64, 0x64,
	0x12, 0x25, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x73, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x61, 0x73, 0x61,
	0x2e, 0x70, 0x62, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x73, 0x68,
	0x65, 0x76, 0x2e, 0x6d, 0x61, 0x73, 0x61, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67,
	0x68, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x59, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x6d, 0x61, 0x63, 0x68,
	0x73, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x61, 0x73, 0x61, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x68, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x73, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x61, 0x73,
	0x61, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xdc, 0x01, 0x0a,
	0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x73, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x61,
	0x73, 0x61, 0x2e, 0x70, 0x62, 0x2e, 0x74, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x42, 0x0d,
	0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x63, 0x68,
	0x73, 0x68, 0x65, 0x76, 0x2f, 0x6d, 0x61, 0x73, 0x61, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x68, 0x6f,
	0x75, 0x67, 0x68, 0x74, 0x73, 0xa2, 0x02, 0x04, 0x4d, 0x4d, 0x50, 0x54, 0xaa, 0x02, 0x19, 0x4d,
	0x61, 0x63, 0x68, 0x73, 0x68, 0x65, 0x76, 0x2e, 0x4d, 0x61, 0x73, 0x61, 0x2e, 0x50, 0x62, 0x2e,
	0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0xca, 0x02, 0x19, 0x4d, 0x61, 0x63, 0x68, 0x73,
	0x68, 0x65, 0x76, 0x5c, 0x4d, 0x61, 0x73, 0x61, 0x5c, 0x50, 0x62, 0x5c, 0x54, 0x68, 0x6f, 0x75,
	0x67, 0x68, 0x74, 0x73, 0xe2, 0x02, 0x25, 0x4d, 0x61, 0x63, 0x68, 0x73, 0x68, 0x65, 0x76, 0x5c,
	0x4d, 0x61, 0x73, 0x61, 0x5c, 0x50, 0x62, 0x5c, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x4d,
	0x61, 0x63, 0x68, 0x73, 0x68, 0x65, 0x76, 0x3a, 0x3a, 0x4d, 0x61, 0x73, 0x61, 0x3a, 0x3a, 0x50,
	0x62, 0x3a, 0x3a, 0x54, 0x68, 0x6f, 0x75, 0x67, 0x68, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_thoughts_thoughts_proto_rawDescOnce sync.Once
	file_thoughts_thoughts_proto_rawDescData = file_thoughts_thoughts_proto_rawDesc
)

func file_thoughts_thoughts_proto_rawDescGZIP() []byte {
	file_thoughts_thoughts_proto_rawDescOnce.Do(func() {
		file_thoughts_thoughts_proto_rawDescData = protoimpl.X.CompressGZIP(file_thoughts_thoughts_proto_rawDescData)
	})
	return file_thoughts_thoughts_proto_rawDescData
}

var file_thoughts_thoughts_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_thoughts_thoughts_proto_goTypes = []interface{}{
	(*Thought)(nil),               // 0: machshev.masa.pb.thoughts.Thought
	(*ThoughtPad)(nil),            // 1: machshev.masa.pb.thoughts.ThoughtPad
	(*AddRequest)(nil),            // 2: machshev.masa.pb.thoughts.AddRequest
	(*AddResponse)(nil),           // 3: machshev.masa.pb.thoughts.AddResponse
	(*ListRequest)(nil),           // 4: machshev.masa.pb.thoughts.ListRequest
	(*ListResponse)(nil),          // 5: machshev.masa.pb.thoughts.ListResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_thoughts_thoughts_proto_depIdxs = []int32{
	6, // 0: machshev.masa.pb.thoughts.Thought.created:type_name -> google.protobuf.Timestamp
	0, // 1: machshev.masa.pb.thoughts.ThoughtPad.thoughts:type_name -> machshev.masa.pb.thoughts.Thought
	0, // 2: machshev.masa.pb.thoughts.ListResponse.thoughts:type_name -> machshev.masa.pb.thoughts.Thought
	2, // 3: machshev.masa.pb.thoughts.ThoughtService.Add:input_type -> machshev.masa.pb.thoughts.AddRequest
	4, // 4: machshev.masa.pb.thoughts.ThoughtService.List:input_type -> machshev.masa.pb.thoughts.ListRequest
	3, // 5: machshev.masa.pb.thoughts.ThoughtService.Add:output_type -> machshev.masa.pb.thoughts.AddResponse
	5, // 6: machshev.masa.pb.thoughts.ThoughtService.List:output_type -> machshev.masa.pb.thoughts.ListResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_thoughts_thoughts_proto_init() }
func file_thoughts_thoughts_proto_init() {
	if File_thoughts_thoughts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thoughts_thoughts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Thought); i {
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
		file_thoughts_thoughts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThoughtPad); i {
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
		file_thoughts_thoughts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_thoughts_thoughts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_thoughts_thoughts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_thoughts_thoughts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_thoughts_thoughts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thoughts_thoughts_proto_goTypes,
		DependencyIndexes: file_thoughts_thoughts_proto_depIdxs,
		MessageInfos:      file_thoughts_thoughts_proto_msgTypes,
	}.Build()
	File_thoughts_thoughts_proto = out.File
	file_thoughts_thoughts_proto_rawDesc = nil
	file_thoughts_thoughts_proto_goTypes = nil
	file_thoughts_thoughts_proto_depIdxs = nil
}
