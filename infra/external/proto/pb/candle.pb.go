// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: candle.proto

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

type Candle struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=Parity,proto3" json:"Parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=Exchange,proto3" json:"Exchange,omitempty"`
	Mts           uint64                 `protobuf:"varint,3,opt,name=Mts,proto3" json:"Mts,omitempty"`
	Open          float64                `protobuf:"fixed64,4,opt,name=Open,proto3" json:"Open,omitempty"`
	Close         float64                `protobuf:"fixed64,5,opt,name=Close,proto3" json:"Close,omitempty"`
	High          float64                `protobuf:"fixed64,6,opt,name=High,proto3" json:"High,omitempty"`
	Low           float64                `protobuf:"fixed64,7,opt,name=Low,proto3" json:"Low,omitempty"`
	Volume        float64                `protobuf:"fixed64,8,opt,name=Volume,proto3" json:"Volume,omitempty"`
	Roc           float64                `protobuf:"fixed64,9,opt,name=Roc,proto3" json:"Roc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Candle) Reset() {
	*x = Candle{}
	mi := &file_candle_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Candle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candle) ProtoMessage() {}

func (x *Candle) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candle.ProtoReflect.Descriptor instead.
func (*Candle) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{0}
}

func (x *Candle) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *Candle) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *Candle) GetMts() uint64 {
	if x != nil {
		return x.Mts
	}
	return 0
}

func (x *Candle) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Candle) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Candle) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Candle) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Candle) GetVolume() float64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *Candle) GetRoc() float64 {
	if x != nil {
		return x.Roc
	}
	return 0
}

type GetCandleFirstMtsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCandleFirstMtsRequest) Reset() {
	*x = GetCandleFirstMtsRequest{}
	mi := &file_candle_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCandleFirstMtsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandleFirstMtsRequest) ProtoMessage() {}

func (x *GetCandleFirstMtsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandleFirstMtsRequest.ProtoReflect.Descriptor instead.
func (*GetCandleFirstMtsRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{1}
}

func (x *GetCandleFirstMtsRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *GetCandleFirstMtsRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

type GetCandleFirstMtsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCandleFirstMtsResponse) Reset() {
	*x = GetCandleFirstMtsResponse{}
	mi := &file_candle_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCandleFirstMtsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCandleFirstMtsResponse) ProtoMessage() {}

func (x *GetCandleFirstMtsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCandleFirstMtsResponse.ProtoReflect.Descriptor instead.
func (*GetCandleFirstMtsResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{2}
}

func (x *GetCandleFirstMtsResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type GetLastTwoCandlesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLastTwoCandlesRequest) Reset() {
	*x = GetLastTwoCandlesRequest{}
	mi := &file_candle_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLastTwoCandlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastTwoCandlesRequest) ProtoMessage() {}

func (x *GetLastTwoCandlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastTwoCandlesRequest.ProtoReflect.Descriptor instead.
func (*GetLastTwoCandlesRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{3}
}

func (x *GetLastTwoCandlesRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *GetLastTwoCandlesRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

type GetLastTwoCandlesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLastTwoCandlesResponse) Reset() {
	*x = GetLastTwoCandlesResponse{}
	mi := &file_candle_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLastTwoCandlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastTwoCandlesResponse) ProtoMessage() {}

func (x *GetLastTwoCandlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastTwoCandlesResponse.ProtoReflect.Descriptor instead.
func (*GetLastTwoCandlesResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{4}
}

func (x *GetLastTwoCandlesResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type CreateCandlesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCandlesRequest) Reset() {
	*x = CreateCandlesRequest{}
	mi := &file_candle_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCandlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCandlesRequest) ProtoMessage() {}

func (x *CreateCandlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCandlesRequest.ProtoReflect.Descriptor instead.
func (*CreateCandlesRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCandlesRequest) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type CreateCandlesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCandlesResponse) Reset() {
	*x = CreateCandlesResponse{}
	mi := &file_candle_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCandlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCandlesResponse) ProtoMessage() {}

func (x *CreateCandlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCandlesResponse.ProtoReflect.Descriptor instead.
func (*CreateCandlesResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCandlesResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type ListCandleLimitRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Parity        uint64                 `protobuf:"varint,1,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange      uint64                 `protobuf:"varint,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Limit         uint64                 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCandleLimitRequest) Reset() {
	*x = ListCandleLimitRequest{}
	mi := &file_candle_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCandleLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCandleLimitRequest) ProtoMessage() {}

func (x *ListCandleLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCandleLimitRequest.ProtoReflect.Descriptor instead.
func (*ListCandleLimitRequest) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{7}
}

func (x *ListCandleLimitRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *ListCandleLimitRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *ListCandleLimitRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListCandleLimitResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Candles       []*Candle              `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCandleLimitResponse) Reset() {
	*x = ListCandleLimitResponse{}
	mi := &file_candle_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCandleLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCandleLimitResponse) ProtoMessage() {}

func (x *ListCandleLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_candle_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCandleLimitResponse.ProtoReflect.Descriptor instead.
func (*ListCandleLimitResponse) Descriptor() ([]byte, []int) {
	return file_candle_proto_rawDescGZIP(), []int{8}
}

func (x *ListCandleLimitResponse) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

var File_candle_proto protoreflect.FileDescriptor

var file_candle_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0xc8, 0x01, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x50,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x4d, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x48, 0x69, 0x67, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x48, 0x69, 0x67,
	0x68, 0x12, 0x10, 0x0a, 0x03, 0x4c, 0x6f, 0x77, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x4c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x52,
	0x6f, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x52, 0x6f, 0x63, 0x22, 0x4e, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x41, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x22, 0x4e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x77, 0x6f, 0x43, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x22, 0x41, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x77, 0x6f, 0x43, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a,
	0x07, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x22, 0x3c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x63,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x22, 0x3d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x22, 0x62, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x3f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x07, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x07, 0x63, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_candle_proto_rawDescOnce sync.Once
	file_candle_proto_rawDescData []byte
)

func file_candle_proto_rawDescGZIP() []byte {
	file_candle_proto_rawDescOnce.Do(func() {
		file_candle_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_candle_proto_rawDesc), len(file_candle_proto_rawDesc)))
	})
	return file_candle_proto_rawDescData
}

var file_candle_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_candle_proto_goTypes = []any{
	(*Candle)(nil),                    // 0: pb.Candle
	(*GetCandleFirstMtsRequest)(nil),  // 1: pb.GetCandleFirstMtsRequest
	(*GetCandleFirstMtsResponse)(nil), // 2: pb.GetCandleFirstMtsResponse
	(*GetLastTwoCandlesRequest)(nil),  // 3: pb.GetLastTwoCandlesRequest
	(*GetLastTwoCandlesResponse)(nil), // 4: pb.GetLastTwoCandlesResponse
	(*CreateCandlesRequest)(nil),      // 5: pb.CreateCandlesRequest
	(*CreateCandlesResponse)(nil),     // 6: pb.CreateCandlesResponse
	(*ListCandleLimitRequest)(nil),    // 7: pb.ListCandleLimitRequest
	(*ListCandleLimitResponse)(nil),   // 8: pb.ListCandleLimitResponse
}
var file_candle_proto_depIdxs = []int32{
	0, // 0: pb.GetCandleFirstMtsResponse.candles:type_name -> pb.Candle
	0, // 1: pb.GetLastTwoCandlesResponse.candles:type_name -> pb.Candle
	0, // 2: pb.CreateCandlesRequest.candles:type_name -> pb.Candle
	0, // 3: pb.CreateCandlesResponse.candles:type_name -> pb.Candle
	0, // 4: pb.ListCandleLimitResponse.candles:type_name -> pb.Candle
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_candle_proto_init() }
func file_candle_proto_init() {
	if File_candle_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_candle_proto_rawDesc), len(file_candle_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_candle_proto_goTypes,
		DependencyIndexes: file_candle_proto_depIdxs,
		MessageInfos:      file_candle_proto_msgTypes,
	}.Build()
	File_candle_proto = out.File
	file_candle_proto_goTypes = nil
	file_candle_proto_depIdxs = nil
}
