// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.1
// source: trade_config.proto

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

type TradeConfig struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	TradeConfig             int64                  `protobuf:"varint,1,opt,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	User                    int64                  `protobuf:"varint,2,opt,name=user,proto3" json:"user,omitempty"`
	Modality                int64                  `protobuf:"varint,3,opt,name=modality,proto3" json:"modality,omitempty"`
	Strategy                int64                  `protobuf:"varint,4,opt,name=strategy,proto3" json:"strategy,omitempty"`
	StrategyEnabled         bool                   `protobuf:"varint,5,opt,name=strategy_enabled,json=strategyEnabled,proto3" json:"strategy_enabled,omitempty"`
	StrategyVariant         int64                  `protobuf:"varint,6,opt,name=strategy_variant,json=strategyVariant,proto3" json:"strategy_variant,omitempty"`
	Parity                  int64                  `protobuf:"varint,7,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange                int64                  `protobuf:"varint,8,opt,name=exchange,proto3" json:"exchange,omitempty"`
	OperationQuantity       int64                  `protobuf:"varint,9,opt,name=operation_quantity,json=operationQuantity,proto3" json:"operation_quantity,omitempty"`
	OperationAmount         float64                `protobuf:"fixed64,10,opt,name=operation_amount,json=operationAmount,proto3" json:"operation_amount,omitempty"`
	Enabled                 bool                   `protobuf:"varint,11,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DefaultProfitPercentage float64                `protobuf:"fixed64,12,opt,name=default_profit_percentage,json=defaultProfitPercentage,proto3" json:"default_profit_percentage,omitempty"`
	WalletValueLimit        float64                `protobuf:"fixed64,13,opt,name=wallet_value_limit,json=walletValueLimit,proto3" json:"wallet_value_limit,omitempty"`
	UserName                string                 `protobuf:"bytes,14,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	ModalityName            string                 `protobuf:"bytes,15,opt,name=modality_name,json=modalityName,proto3" json:"modality_name,omitempty"`
	StrategyName            string                 `protobuf:"bytes,16,opt,name=strategy_name,json=strategyName,proto3" json:"strategy_name,omitempty"`
	StrategyVariantName     string                 `protobuf:"bytes,17,opt,name=strategy_variant_name,json=strategyVariantName,proto3" json:"strategy_variant_name,omitempty"`
	StrategyVariantEnabled  bool                   `protobuf:"varint,18,opt,name=strategy_variant_enabled,json=strategyVariantEnabled,proto3" json:"strategy_variant_enabled,omitempty"`
	SymbolName              string                 `protobuf:"bytes,19,opt,name=symbol_name,json=symbolName,proto3" json:"symbol_name,omitempty"`
	ExchangeName            string                 `protobuf:"bytes,20,opt,name=exchange_name,json=exchangeName,proto3" json:"exchange_name,omitempty"`
	ParitySymbol            string                 `protobuf:"bytes,21,opt,name=parity_symbol,json=paritySymbol,proto3" json:"parity_symbol,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *TradeConfig) Reset() {
	*x = TradeConfig{}
	mi := &file_trade_config_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TradeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeConfig) ProtoMessage() {}

func (x *TradeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeConfig.ProtoReflect.Descriptor instead.
func (*TradeConfig) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{0}
}

func (x *TradeConfig) GetTradeConfig() int64 {
	if x != nil {
		return x.TradeConfig
	}
	return 0
}

func (x *TradeConfig) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

func (x *TradeConfig) GetModality() int64 {
	if x != nil {
		return x.Modality
	}
	return 0
}

func (x *TradeConfig) GetStrategy() int64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *TradeConfig) GetStrategyEnabled() bool {
	if x != nil {
		return x.StrategyEnabled
	}
	return false
}

func (x *TradeConfig) GetStrategyVariant() int64 {
	if x != nil {
		return x.StrategyVariant
	}
	return 0
}

func (x *TradeConfig) GetParity() int64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *TradeConfig) GetExchange() int64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *TradeConfig) GetOperationQuantity() int64 {
	if x != nil {
		return x.OperationQuantity
	}
	return 0
}

func (x *TradeConfig) GetOperationAmount() float64 {
	if x != nil {
		return x.OperationAmount
	}
	return 0
}

func (x *TradeConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *TradeConfig) GetDefaultProfitPercentage() float64 {
	if x != nil {
		return x.DefaultProfitPercentage
	}
	return 0
}

func (x *TradeConfig) GetWalletValueLimit() float64 {
	if x != nil {
		return x.WalletValueLimit
	}
	return 0
}

func (x *TradeConfig) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *TradeConfig) GetModalityName() string {
	if x != nil {
		return x.ModalityName
	}
	return ""
}

func (x *TradeConfig) GetStrategyName() string {
	if x != nil {
		return x.StrategyName
	}
	return ""
}

func (x *TradeConfig) GetStrategyVariantName() string {
	if x != nil {
		return x.StrategyVariantName
	}
	return ""
}

func (x *TradeConfig) GetStrategyVariantEnabled() bool {
	if x != nil {
		return x.StrategyVariantEnabled
	}
	return false
}

func (x *TradeConfig) GetSymbolName() string {
	if x != nil {
		return x.SymbolName
	}
	return ""
}

func (x *TradeConfig) GetExchangeName() string {
	if x != nil {
		return x.ExchangeName
	}
	return ""
}

func (x *TradeConfig) GetParitySymbol() string {
	if x != nil {
		return x.ParitySymbol
	}
	return ""
}

type TradeConfigResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TradeConfig   []*TradeConfig         `protobuf:"bytes,1,rep,name=trade_config,json=tradeConfig,proto3" json:"trade_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TradeConfigResponse) Reset() {
	*x = TradeConfigResponse{}
	mi := &file_trade_config_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TradeConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeConfigResponse) ProtoMessage() {}

func (x *TradeConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeConfigResponse.ProtoReflect.Descriptor instead.
func (*TradeConfigResponse) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{1}
}

func (x *TradeConfigResponse) GetTradeConfig() []*TradeConfig {
	if x != nil {
		return x.TradeConfig
	}
	return nil
}

type CreateTradeConfigRequest struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	Modality                int64                  `protobuf:"varint,1,opt,name=modality,proto3" json:"modality,omitempty"`
	Strategy                int64                  `protobuf:"varint,2,opt,name=strategy,proto3" json:"strategy,omitempty"`
	StrategyVariant         int64                  `protobuf:"varint,3,opt,name=strategy_variant,json=strategyVariant,proto3" json:"strategy_variant,omitempty"`
	Parity                  int64                  `protobuf:"varint,4,opt,name=parity,proto3" json:"parity,omitempty"`
	Exchange                int64                  `protobuf:"varint,5,opt,name=exchange,proto3" json:"exchange,omitempty"`
	OperationQuantity       int64                  `protobuf:"varint,6,opt,name=operation_quantity,json=operationQuantity,proto3" json:"operation_quantity,omitempty"`
	OperationAmount         float64                `protobuf:"fixed64,7,opt,name=operation_amount,json=operationAmount,proto3" json:"operation_amount,omitempty"`
	DefaultProfitPercentage float64                `protobuf:"fixed64,8,opt,name=default_profit_percentage,json=defaultProfitPercentage,proto3" json:"default_profit_percentage,omitempty"`
	WalletValueLimit        float64                `protobuf:"fixed64,9,opt,name=wallet_value_limit,json=walletValueLimit,proto3" json:"wallet_value_limit,omitempty"`
	Enabled                 bool                   `protobuf:"varint,10,opt,name=enabled,proto3" json:"enabled,omitempty"`
	User                    int64                  `protobuf:"varint,11,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *CreateTradeConfigRequest) Reset() {
	*x = CreateTradeConfigRequest{}
	mi := &file_trade_config_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTradeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTradeConfigRequest) ProtoMessage() {}

func (x *CreateTradeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTradeConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateTradeConfigRequest) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTradeConfigRequest) GetModality() int64 {
	if x != nil {
		return x.Modality
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetStrategy() int64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetStrategyVariant() int64 {
	if x != nil {
		return x.StrategyVariant
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetParity() int64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetExchange() int64 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetOperationQuantity() int64 {
	if x != nil {
		return x.OperationQuantity
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetOperationAmount() float64 {
	if x != nil {
		return x.OperationAmount
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetDefaultProfitPercentage() float64 {
	if x != nil {
		return x.DefaultProfitPercentage
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetWalletValueLimit() float64 {
	if x != nil {
		return x.WalletValueLimit
	}
	return 0
}

func (x *CreateTradeConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CreateTradeConfigRequest) GetUser() int64 {
	if x != nil {
		return x.User
	}
	return 0
}

type ListTradeConfigRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTradeConfigRequest) Reset() {
	*x = ListTradeConfigRequest{}
	mi := &file_trade_config_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTradeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTradeConfigRequest) ProtoMessage() {}

func (x *ListTradeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTradeConfigRequest.ProtoReflect.Descriptor instead.
func (*ListTradeConfigRequest) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{3}
}

type UpdateTradeConfigRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enabled       bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Strategy      uint64                 `protobuf:"varint,2,opt,name=strategy,proto3" json:"strategy,omitempty"`
	Parity        uint64                 `protobuf:"varint,3,opt,name=parity,proto3" json:"parity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTradeConfigRequest) Reset() {
	*x = UpdateTradeConfigRequest{}
	mi := &file_trade_config_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTradeConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTradeConfigRequest) ProtoMessage() {}

func (x *UpdateTradeConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTradeConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateTradeConfigRequest) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTradeConfigRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UpdateTradeConfigRequest) GetStrategy() uint64 {
	if x != nil {
		return x.Strategy
	}
	return 0
}

func (x *UpdateTradeConfigRequest) GetParity() uint64 {
	if x != nil {
		return x.Parity
	}
	return 0
}

type UpdateTradeConfigResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tradeconfig   *TradeConfig           `protobuf:"bytes,1,opt,name=tradeconfig,proto3" json:"tradeconfig,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTradeConfigResponse) Reset() {
	*x = UpdateTradeConfigResponse{}
	mi := &file_trade_config_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTradeConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTradeConfigResponse) ProtoMessage() {}

func (x *UpdateTradeConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trade_config_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTradeConfigResponse.ProtoReflect.Descriptor instead.
func (*UpdateTradeConfigResponse) Descriptor() ([]byte, []int) {
	return file_trade_config_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTradeConfigResponse) GetTradeconfig() *TradeConfig {
	if x != nil {
		return x.Tradeconfig
	}
	return nil
}

var File_trade_config_proto protoreflect.FileDescriptor

var file_trade_config_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xa4, 0x06, 0x0a, 0x0b, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x29, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x19, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x2c, 0x0a, 0x12, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x56,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22,
	0x49, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xa3, 0x03, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x29, 0x0a, 0x10, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2d,
	0x0a, 0x12, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a,
	0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x19, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x17, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x10, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x68, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x22, 0x4e, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x64, 0x65, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_trade_config_proto_rawDescOnce sync.Once
	file_trade_config_proto_rawDescData []byte
)

func file_trade_config_proto_rawDescGZIP() []byte {
	file_trade_config_proto_rawDescOnce.Do(func() {
		file_trade_config_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_trade_config_proto_rawDesc), len(file_trade_config_proto_rawDesc)))
	})
	return file_trade_config_proto_rawDescData
}

var file_trade_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_trade_config_proto_goTypes = []any{
	(*TradeConfig)(nil),               // 0: pb.TradeConfig
	(*TradeConfigResponse)(nil),       // 1: pb.TradeConfigResponse
	(*CreateTradeConfigRequest)(nil),  // 2: pb.CreateTradeConfigRequest
	(*ListTradeConfigRequest)(nil),    // 3: pb.ListTradeConfigRequest
	(*UpdateTradeConfigRequest)(nil),  // 4: pb.UpdateTradeConfigRequest
	(*UpdateTradeConfigResponse)(nil), // 5: pb.UpdateTradeConfigResponse
}
var file_trade_config_proto_depIdxs = []int32{
	0, // 0: pb.TradeConfigResponse.trade_config:type_name -> pb.TradeConfig
	0, // 1: pb.UpdateTradeConfigResponse.tradeconfig:type_name -> pb.TradeConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trade_config_proto_init() }
func file_trade_config_proto_init() {
	if File_trade_config_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_trade_config_proto_rawDesc), len(file_trade_config_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trade_config_proto_goTypes,
		DependencyIndexes: file_trade_config_proto_depIdxs,
		MessageInfos:      file_trade_config_proto_msgTypes,
	}.Build()
	File_trade_config_proto = out.File
	file_trade_config_proto_goTypes = nil
	file_trade_config_proto_depIdxs = nil
}
