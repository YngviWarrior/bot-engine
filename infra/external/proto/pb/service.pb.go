// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: service.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x12, 0x74, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0c, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x61, 0x76, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x32, 0xc2, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0xf1, 0x0e, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72,
	0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a,
	0x1b, 0x47, 0x65, 0x74, 0x41, 0x76, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x50, 0x61,
	0x72, 0x69, 0x74, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x26, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x50,
	0x61, 0x72, 0x69, 0x74, 0x79, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x76, 0x67,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x45, 0x78, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x37, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x1c, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68,
	0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f,
	0x69, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d,
	0x74, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46,
	0x69, 0x72, 0x73, 0x74, 0x4d, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x77, 0x6f, 0x43,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x76, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x76, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x76, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_service_proto_goTypes = []any{
	(*CreateUserRequest)(nil),                   // 0: pb.CreateUserRequest
	(*GetUserByEmailRequest)(nil),               // 1: pb.GetUserByEmailRequest
	(*GetUserByIdRequest)(nil),                  // 2: pb.GetUserByIdRequest
	(*ListTradeConfigRequest)(nil),              // 3: pb.ListTradeConfigRequest
	(*CreateTradeConfigRequest)(nil),            // 4: pb.CreateTradeConfigRequest
	(*UpdateTradeConfigRequest)(nil),            // 5: pb.UpdateTradeConfigRequest
	(*GetTradeConfigRequest)(nil),               // 6: pb.GetTradeConfigRequest
	(*CreateUserStrategyRequest)(nil),           // 7: pb.CreateUserStrategyRequest
	(*ListUserStrategyRequest)(nil),             // 8: pb.ListUserStrategyRequest
	(*ListParityRequest)(nil),                   // 9: pb.ListParityRequest
	(*GetAvgPriceByParityExchangeRequest)(nil),  // 10: pb.GetAvgPriceByParityExchangeRequest
	(*ListCoinRequest)(nil),                     // 11: pb.ListCoinRequest
	(*GetWalletWithCoinRequest)(nil),            // 12: pb.GetWalletWithCoinRequest
	(*ListWalletWithCoinRequest)(nil),           // 13: pb.ListWalletWithCoinRequest
	(*CreateWalletRequest)(nil),                 // 14: pb.CreateWalletRequest
	(*UpdateWalletRequest)(nil),                 // 15: pb.UpdateWalletRequest
	(*GetCandleFirstMtsRequest)(nil),            // 16: pb.GetCandleFirstMtsRequest
	(*CreateCandlesRequest)(nil),                // 17: pb.CreateCandlesRequest
	(*ListCandleLimitRequest)(nil),              // 18: pb.ListCandleLimitRequest
	(*ListAvgPricesRequest)(nil),                // 19: pb.ListAvgPricesRequest
	(*GetFirstPriceRequest)(nil),                // 20: pb.GetFirstPriceRequest
	(*CreateAveragePriceRequest)(nil),           // 21: pb.CreateAveragePriceRequest
	(*ListOperationRequest)(nil),                // 22: pb.ListOperationRequest
	(*ListOperationByPeriodRequest)(nil),        // 23: pb.ListOperationByPeriodRequest
	(*ListAllOperationRequest)(nil),             // 24: pb.ListAllOperationRequest
	(*UpdateOperationRequest)(nil),              // 25: pb.UpdateOperationRequest
	(*UserResponse)(nil),                        // 26: pb.UserResponse
	(*TradeConfigResponse)(nil),                 // 27: pb.TradeConfigResponse
	(*UpdateTradeConfigResponse)(nil),           // 28: pb.UpdateTradeConfigResponse
	(*GetTradeConfigResponse)(nil),              // 29: pb.GetTradeConfigResponse
	(*UserStrategyResponse)(nil),                // 30: pb.UserStrategyResponse
	(*ListParityResponse)(nil),                  // 31: pb.ListParityResponse
	(*GetAvgPriceByParityExchangeResponse)(nil), // 32: pb.GetAvgPriceByParityExchangeResponse
	(*ListCoinResponse)(nil),                    // 33: pb.ListCoinResponse
	(*GetWalletWithCoinResponse)(nil),           // 34: pb.GetWalletWithCoinResponse
	(*ListWalletWithCoinResponse)(nil),          // 35: pb.ListWalletWithCoinResponse
	(*CreateWalletResponse)(nil),                // 36: pb.CreateWalletResponse
	(*UpdateWalletResponse)(nil),                // 37: pb.UpdateWalletResponse
	(*GetCandleFirstMtsResponse)(nil),           // 38: pb.GetCandleFirstMtsResponse
	(*CreateCandlesResponse)(nil),               // 39: pb.CreateCandlesResponse
	(*ListCandleLimitResponse)(nil),             // 40: pb.ListCandleLimitResponse
	(*ListAvgPricesResponse)(nil),               // 41: pb.ListAvgPricesResponse
	(*GetFirstPriceResponse)(nil),               // 42: pb.GetFirstPriceResponse
	(*CreateAveragePriceResponse)(nil),          // 43: pb.CreateAveragePriceResponse
	(*ListOperationResponse)(nil),               // 44: pb.ListOperationResponse
	(*ListOperationByPeriodResponse)(nil),       // 45: pb.ListOperationByPeriodResponse
	(*ListAllOperationResponse)(nil),            // 46: pb.ListAllOperationResponse
	(*UpdateOperationResponse)(nil),             // 47: pb.UpdateOperationResponse
}
var file_service_proto_depIdxs = []int32{
	0,  // 0: pb.UserService.CreateUser:input_type -> pb.CreateUserRequest
	1,  // 1: pb.UserService.GetUserByEmail:input_type -> pb.GetUserByEmailRequest
	2,  // 2: pb.UserService.GetUserById:input_type -> pb.GetUserByIdRequest
	3,  // 3: pb.ExchangeService.ListTradeConfig:input_type -> pb.ListTradeConfigRequest
	4,  // 4: pb.ExchangeService.CreateTradeConfig:input_type -> pb.CreateTradeConfigRequest
	5,  // 5: pb.ExchangeService.UpdateTradeConfig:input_type -> pb.UpdateTradeConfigRequest
	6,  // 6: pb.ExchangeService.GetTradeConfig:input_type -> pb.GetTradeConfigRequest
	7,  // 7: pb.ExchangeService.CreateUserStrategy:input_type -> pb.CreateUserStrategyRequest
	8,  // 8: pb.ExchangeService.ListUserStrategy:input_type -> pb.ListUserStrategyRequest
	9,  // 9: pb.ExchangeService.ListParity:input_type -> pb.ListParityRequest
	10, // 10: pb.ExchangeService.GetAvgPriceByParityExchange:input_type -> pb.GetAvgPriceByParityExchangeRequest
	11, // 11: pb.ExchangeService.ListCoin:input_type -> pb.ListCoinRequest
	12, // 12: pb.ExchangeService.GetWalletWithCoin:input_type -> pb.GetWalletWithCoinRequest
	13, // 13: pb.ExchangeService.ListWalletWithCoin:input_type -> pb.ListWalletWithCoinRequest
	14, // 14: pb.ExchangeService.CreateWallet:input_type -> pb.CreateWalletRequest
	15, // 15: pb.ExchangeService.UpdateWallet:input_type -> pb.UpdateWalletRequest
	16, // 16: pb.ExchangeService.GetCandleFirstMts:input_type -> pb.GetCandleFirstMtsRequest
	16, // 17: pb.ExchangeService.GetLastTwoCandles:input_type -> pb.GetCandleFirstMtsRequest
	17, // 18: pb.ExchangeService.CreateCandles:input_type -> pb.CreateCandlesRequest
	18, // 19: pb.ExchangeService.ListCandleLimit:input_type -> pb.ListCandleLimitRequest
	19, // 20: pb.ExchangeService.ListAvgPrices:input_type -> pb.ListAvgPricesRequest
	20, // 21: pb.ExchangeService.GetFirstPrice:input_type -> pb.GetFirstPriceRequest
	21, // 22: pb.ExchangeService.CreateAveragePrice:input_type -> pb.CreateAveragePriceRequest
	22, // 23: pb.ExchangeService.ListOperation:input_type -> pb.ListOperationRequest
	23, // 24: pb.ExchangeService.ListOperationByPeriod:input_type -> pb.ListOperationByPeriodRequest
	24, // 25: pb.ExchangeService.ListAllOperation:input_type -> pb.ListAllOperationRequest
	25, // 26: pb.ExchangeService.UpdateOperation:input_type -> pb.UpdateOperationRequest
	26, // 27: pb.UserService.CreateUser:output_type -> pb.UserResponse
	26, // 28: pb.UserService.GetUserByEmail:output_type -> pb.UserResponse
	26, // 29: pb.UserService.GetUserById:output_type -> pb.UserResponse
	27, // 30: pb.ExchangeService.ListTradeConfig:output_type -> pb.TradeConfigResponse
	27, // 31: pb.ExchangeService.CreateTradeConfig:output_type -> pb.TradeConfigResponse
	28, // 32: pb.ExchangeService.UpdateTradeConfig:output_type -> pb.UpdateTradeConfigResponse
	29, // 33: pb.ExchangeService.GetTradeConfig:output_type -> pb.GetTradeConfigResponse
	30, // 34: pb.ExchangeService.CreateUserStrategy:output_type -> pb.UserStrategyResponse
	30, // 35: pb.ExchangeService.ListUserStrategy:output_type -> pb.UserStrategyResponse
	31, // 36: pb.ExchangeService.ListParity:output_type -> pb.ListParityResponse
	32, // 37: pb.ExchangeService.GetAvgPriceByParityExchange:output_type -> pb.GetAvgPriceByParityExchangeResponse
	33, // 38: pb.ExchangeService.ListCoin:output_type -> pb.ListCoinResponse
	34, // 39: pb.ExchangeService.GetWalletWithCoin:output_type -> pb.GetWalletWithCoinResponse
	35, // 40: pb.ExchangeService.ListWalletWithCoin:output_type -> pb.ListWalletWithCoinResponse
	36, // 41: pb.ExchangeService.CreateWallet:output_type -> pb.CreateWalletResponse
	37, // 42: pb.ExchangeService.UpdateWallet:output_type -> pb.UpdateWalletResponse
	38, // 43: pb.ExchangeService.GetCandleFirstMts:output_type -> pb.GetCandleFirstMtsResponse
	38, // 44: pb.ExchangeService.GetLastTwoCandles:output_type -> pb.GetCandleFirstMtsResponse
	39, // 45: pb.ExchangeService.CreateCandles:output_type -> pb.CreateCandlesResponse
	40, // 46: pb.ExchangeService.ListCandleLimit:output_type -> pb.ListCandleLimitResponse
	41, // 47: pb.ExchangeService.ListAvgPrices:output_type -> pb.ListAvgPricesResponse
	42, // 48: pb.ExchangeService.GetFirstPrice:output_type -> pb.GetFirstPriceResponse
	43, // 49: pb.ExchangeService.CreateAveragePrice:output_type -> pb.CreateAveragePriceResponse
	44, // 50: pb.ExchangeService.ListOperation:output_type -> pb.ListOperationResponse
	45, // 51: pb.ExchangeService.ListOperationByPeriod:output_type -> pb.ListOperationByPeriodResponse
	46, // 52: pb.ExchangeService.ListAllOperation:output_type -> pb.ListAllOperationResponse
	47, // 53: pb.ExchangeService.UpdateOperation:output_type -> pb.UpdateOperationResponse
	27, // [27:54] is the sub-list for method output_type
	0,  // [0:27] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_trade_config_proto_init()
	file_user_proto_init()
	file_user_strategy_proto_init()
	file_wallet_proto_init()
	file_coin_proto_init()
	file_candle_proto_init()
	file_parity_proto_init()
	file_operation_proto_init()
	file_avg_price_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_service_proto_rawDesc), len(file_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
	}.Build()
	File_service_proto = out.File
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
