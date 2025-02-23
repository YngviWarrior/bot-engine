// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.1
// source: operation.proto

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

type Operation struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Operation       uint64                 `protobuf:"varint,1,opt,name=operation,proto3" json:"operation,omitempty"`
	User            uint64                 `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Parity          uint64                 `protobuf:"varint,3,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange        uint64                 `protobuf:"varint,4,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Strategy        uint64                 `protobuf:"varint,5,opt,name=strategy,proto3" json:"strategy,omitempty"`
	StrategyVariant uint64                 `protobuf:"varint,6,opt,name=strategy_variant,json=strategyVariant,proto3" json:"strategy_variant,omitempty"`
	MtsStart        uint64                 `protobuf:"varint,7,opt,name=mts_start,json=mtsStart,proto3" json:"mts_start,omitempty"`
	MtsFinish       uint64                 `protobuf:"varint,8,opt,name=mts_finish,json=mtsFinish,proto3" json:"mts_finish,omitempty"`
	OpenPrice       float64                `protobuf:"fixed64,9,opt,name=open_price,json=openPrice,proto3" json:"open_price,omitempty"`
	ClosePrice      float64                `protobuf:"fixed64,10,opt,name=close_price,json=closePrice,proto3" json:"close_price,omitempty"`
	InvestedAmount  float64                `protobuf:"fixed64,11,opt,name=invested_amount,json=investedAmount,proto3" json:"invested_amount,omitempty"`
	ProfitAmount    float64                `protobuf:"fixed64,12,opt,name=profit_amount,json=profitAmount,proto3" json:"profit_amount,omitempty"`
	Profit          float64                `protobuf:"fixed64,13,opt,name=profit,proto3" json:"profit,omitempty"`
	Closed          bool                   `protobuf:"varint,14,opt,name=closed,proto3" json:"closed,omitempty"`
	Audit           bool                   `protobuf:"varint,15,opt,name=audit,proto3" json:"audit,omitempty"`
	Enabled         bool                   `protobuf:"varint,16,opt,name=enabled,proto3" json:"enabled,omitempty"`
	TimesCanceled   float64                `protobuf:"fixed64,17,opt,name=times_canceled,json=timesCanceled,proto3" json:"times_canceled,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Operation) Reset() {
	*x = Operation{}
	mi := &file_operation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_operation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetOperation() uint64 {
	if x != nil {
		return x.Operation
	}
	return 0
}

func (x *Operation) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *Operation) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *Operation) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *Operation) GetStrategy() uint64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *Operation) GetStrategyVariant() uint64 {
	if x != nil {
		return x.StrategyVariant
	}
	return 0
}

func (x *Operation) GetMtsStart() uint64 {
	if x != nil {
		return x.MtsStart
	}
	return 0
}

func (x *Operation) GetMtsFinish() uint64 {
	if x != nil {
		return x.MtsFinish
	}
	return 0
}

func (x *Operation) GetOpenPrice() float64 {
	if x != nil {
		return x.OpenPrice
	}
	return 0
}

func (x *Operation) GetClosePrice() float64 {
	if x != nil {
		return x.ClosePrice
	}
	return 0
}

func (x *Operation) GetInvestedAmount() float64 {
	if x != nil {
		return x.InvestedAmount
	}
	return 0
}

func (x *Operation) GetProfitAmount() float64 {
	if x != nil {
		return x.ProfitAmount
	}
	return 0
}

func (x *Operation) GetProfit() float64 {
	if x != nil {
		return x.Profit
	}
	return 0
}

func (x *Operation) GetClosed() bool {
	if x != nil {
		return x.Closed
	}
	return false
}

func (x *Operation) GetAudit() bool {
	if x != nil {
		return x.Audit
	}
	return false
}

func (x *Operation) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Operation) GetTimesCanceled() float64 {
	if x != nil {
		return x.TimesCanceled
	}
	return 0
}

type ListOperationRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	User            uint64                 `protobuf:"varint,1,opt,name=user,proto3" json:"user,omitempty"`
	Strategy        uint64                 `protobuf:"varint,2,opt,name=strategy,proto3" json:"strategy,omitempty"`
	StrategyVariant uint64                 `protobuf:"varint,3,opt,name=strategy_variant,json=strategyVariant,proto3" json:"strategy_variant,omitempty"`
	Parity          uint64                 `protobuf:"varint,4,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange        uint64                 `protobuf:"varint,5,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Closed          bool                   `protobuf:"varint,6,opt,name=closed,proto3" json:"closed,omitempty"`
	Enabled         bool                   `protobuf:"varint,7,opt,name=enabled,proto3" json:"enabled,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ListOperationRequest) Reset() {
	*x = ListOperationRequest{}
	mi := &file_operation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationRequest) ProtoMessage() {}

func (x *ListOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationRequest.ProtoReflect.Descriptor instead.
func (*ListOperationRequest) Descriptor() ([]byte, []int) {
	return file_operation_proto_rawDescGZIP(), []int{1}
}

func (x *ListOperationRequest) GetUser() uint64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *ListOperationRequest) GetStrategy() uint64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *ListOperationRequest) GetStrategyVariant() uint64 {
	if x != nil {
		return x.StrategyVariant
	}
	return 0
}

func (x *ListOperationRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *ListOperationRequest) GetExchange() uint64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *ListOperationRequest) GetClosed() bool {
	if x != nil {
		return x.Closed
	}
	return false
}

func (x *ListOperationRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type ListOperationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operations    []*Operation           `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperationResponse) Reset() {
	*x = ListOperationResponse{}
	mi := &file_operation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationResponse) ProtoMessage() {}

func (x *ListOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationResponse.ProtoReflect.Descriptor instead.
func (*ListOperationResponse) Descriptor() ([]byte, []int) {
	return file_operation_proto_rawDescGZIP(), []int{2}
}

func (x *ListOperationResponse) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

type ListOperationByPeriodRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MtsStart      uint64                 `protobuf:"varint,1,opt,name=mts_start,json=mtsStart,proto3" json:"mts_start,omitempty"`
	MtsEnd        uint64                 `protobuf:"varint,2,opt,name=mts_end,json=mtsEnd,proto3" json:"mts_end,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperationByPeriodRequest) Reset() {
	*x = ListOperationByPeriodRequest{}
	mi := &file_operation_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperationByPeriodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationByPeriodRequest) ProtoMessage() {}

func (x *ListOperationByPeriodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_operation_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationByPeriodRequest.ProtoReflect.Descriptor instead.
func (*ListOperationByPeriodRequest) Descriptor() ([]byte, []int) {
	return file_operation_proto_rawDescGZIP(), []int{3}
}

func (x *ListOperationByPeriodRequest) GetMtsStart() uint64 {
	if x != nil {
		return x.MtsStart
	}
	return 0
}

func (x *ListOperationByPeriodRequest) GetMtsEnd() uint64 {
	if x != nil {
		return x.MtsEnd
	}
	return 0
}

type ListOperationByPeriodResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operations    []*Operation           `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListOperationByPeriodResponse) Reset() {
	*x = ListOperationByPeriodResponse{}
	mi := &file_operation_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOperationByPeriodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOperationByPeriodResponse) ProtoMessage() {}

func (x *ListOperationByPeriodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_operation_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOperationByPeriodResponse.ProtoReflect.Descriptor instead.
func (*ListOperationByPeriodResponse) Descriptor() ([]byte, []int) {
	return file_operation_proto_rawDescGZIP(), []int{4}
}

func (x *ListOperationByPeriodResponse) GetOperations() []*Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

var File_operation_proto protoreflect.FileDescriptor

var file_operation_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x89, 0x04, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x74, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x74, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x74, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x6d, 0x74, 0x73, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a,
	0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0a, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x69, 0x6e, 0x76, 0x65, 0x73, 0x74, 0x65, 0x64, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65, 0x64, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x65,
	0x64, 0x22, 0xd7, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x54, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x74, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x74, 0x73, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x6d, 0x74, 0x73, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6d, 0x74, 0x73, 0x45, 0x6e, 0x64, 0x22, 0x4e, 0x0a, 0x1d, 0x4c, 0x69, 0x73,
	0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_operation_proto_rawDescOnce sync.Once
	file_operation_proto_rawDescData []byte
)

func file_operation_proto_rawDescGZIP() []byte {
	file_operation_proto_rawDescOnce.Do(func() {
		file_operation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_operation_proto_rawDesc), len(file_operation_proto_rawDesc)))
	})
	return file_operation_proto_rawDescData
}

var file_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_operation_proto_goTypes = []any{
	(*Operation)(nil),                     // 0: pb.Operation
	(*ListOperationRequest)(nil),          // 1: pb.ListOperationRequest
	(*ListOperationResponse)(nil),         // 2: pb.ListOperationResponse
	(*ListOperationByPeriodRequest)(nil),  // 3: pb.ListOperationByPeriodRequest
	(*ListOperationByPeriodResponse)(nil), // 4: pb.ListOperationByPeriodResponse
}
var file_operation_proto_depIdxs = []int32{
	0, // 0: pb.ListOperationResponse.operations:type_name -> pb.Operation
	0, // 1: pb.ListOperationByPeriodResponse.operations:type_name -> pb.Operation
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_operation_proto_init() }
func file_operation_proto_init() {
	if File_operation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_operation_proto_rawDesc), len(file_operation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_operation_proto_goTypes,
		DependencyIndexes: file_operation_proto_depIdxs,
		MessageInfos:      file_operation_proto_msgTypes,
	}.Build()
	File_operation_proto = out.File
	file_operation_proto_goTypes = nil
	file_operation_proto_depIdxs = nil
}
