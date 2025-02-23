// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.1
// source: parity.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Parity struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=Parity,proto3" json:"Parity,omitempty"`
	Symbol        string                 `protobuf:"bytes,2,opt,name=Symbol,proto3" json:"Symbol,omitempty"`
	Active        bool                   `protobuf:"varint,3,opt,name=Active,proto3" json:"Active,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Parity) Reset() {
	*x = Parity{}
	mi := &file_parity_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Parity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Parity) ProtoMessage() {}

func (x *Parity) ProtoReflect() protoreflect.Message {
	mi := &file_parity_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Parity.ProtoReflect.Descriptor instead.
func (*Parity) Descriptor() ([]byte, []int) {
	return file_parity_proto_rawDescGZIP(), []int{0}
}

func (x *Parity) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *Parity) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Parity) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

type ListParityRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListParityRequest) Reset() {
	*x = ListParityRequest{}
	mi := &file_parity_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListParityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListParityRequest) ProtoMessage() {}

func (x *ListParityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parity_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListParityRequest.ProtoReflect.Descriptor instead.
func (*ListParityRequest) Descriptor() ([]byte, []int) {
	return file_parity_proto_rawDescGZIP(), []int{1}
}

type ListParityResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parities      []*Parity              `protobuf:"bytes,1,rep,name=parities,proto3" json:"parities,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListParityResponse) Reset() {
	*x = ListParityResponse{}
	mi := &file_parity_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListParityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListParityResponse) ProtoMessage() {}

func (x *ListParityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parity_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListParityResponse.ProtoReflect.Descriptor instead.
func (*ListParityResponse) Descriptor() ([]byte, []int) {
	return file_parity_proto_rawDescGZIP(), []int{2}
}

func (x *ListParityResponse) GetParities() []*Parity {
	if x != nil {
		return x.Parities
	}
	return nil
}

var File_parity_proto protoreflect.FileDescriptor

var file_parity_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x50, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x50, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x26, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_parity_proto_rawDescOnce sync.Once
	file_parity_proto_rawDescData []byte
)

func file_parity_proto_rawDescGZIP() []byte {
	file_parity_proto_rawDescOnce.Do(func() {
		file_parity_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_parity_proto_rawDesc), len(file_parity_proto_rawDesc)))
	})
	return file_parity_proto_rawDescData
}

var file_parity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_parity_proto_goTypes = []any{
	(*Parity)(nil),             // 0: pb.Parity
	(*ListParityRequest)(nil),  // 1: pb.ListParityRequest
	(*ListParityResponse)(nil), // 2: pb.ListParityResponse
}
var file_parity_proto_depIdxs = []int32{
	0, // 0: pb.ListParityResponse.parities:type_name -> pb.Parity
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_parity_proto_init() }
func file_parity_proto_init() {
	if File_parity_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_parity_proto_rawDesc), len(file_parity_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_parity_proto_goTypes,
		DependencyIndexes: file_parity_proto_depIdxs,
		MessageInfos:      file_parity_proto_msgTypes,
	}.Build()
	File_parity_proto = out.File
	file_parity_proto_goTypes = nil
	file_parity_proto_depIdxs = nil
}
