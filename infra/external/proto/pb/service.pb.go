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
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xc2, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x98, 0x0c, 0x0a, 0x0f, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x4c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x69, 0x6e, 0x12,
	0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57,
	0x69, 0x74, 0x68, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x57, 0x69,
	0x74, 0x68, 0x43, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x74, 0x73, 0x12,
	0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x4d, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x4d, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x54, 0x77, 0x6f, 0x43, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x4d, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5e, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x79, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var file_service_proto_goTypes = []any{
	(*CreateUserRequest)(nil),             // 0: pb.CreateUserRequest
	(*GetUserByEmailRequest)(nil),         // 1: pb.GetUserByEmailRequest
	(*GetUserByIdRequest)(nil),            // 2: pb.GetUserByIdRequest
	(*ListTradeConfigRequest)(nil),        // 3: pb.ListTradeConfigRequest
	(*CreateTradeConfigRequest)(nil),      // 4: pb.CreateTradeConfigRequest
	(*UpdateTradeConfigRequest)(nil),      // 5: pb.UpdateTradeConfigRequest
	(*GetTradeConfigRequest)(nil),         // 6: pb.GetTradeConfigRequest
	(*CreateUserStrategyRequest)(nil),     // 7: pb.CreateUserStrategyRequest
	(*ListUserStrategyRequest)(nil),       // 8: pb.ListUserStrategyRequest
	(*ListParityRequest)(nil),             // 9: pb.ListParityRequest
	(*ListCoinRequest)(nil),               // 10: pb.ListCoinRequest
	(*GetWalletWithCoinRequest)(nil),      // 11: pb.GetWalletWithCoinRequest
	(*ListWalletWithCoinRequest)(nil),     // 12: pb.ListWalletWithCoinRequest
	(*CreateWalletRequest)(nil),           // 13: pb.CreateWalletRequest
	(*UpdateWalletRequest)(nil),           // 14: pb.UpdateWalletRequest
	(*GetCandleFirstMtsRequest)(nil),      // 15: pb.GetCandleFirstMtsRequest
	(*CreateCandlesRequest)(nil),          // 16: pb.CreateCandlesRequest
	(*ListCandleLimitRequest)(nil),        // 17: pb.ListCandleLimitRequest
	(*ListOperationRequest)(nil),          // 18: pb.ListOperationRequest
	(*ListOperationByPeriodRequest)(nil),  // 19: pb.ListOperationByPeriodRequest
	(*ListAllOperationRequest)(nil),       // 20: pb.ListAllOperationRequest
	(*UpdateOperationRequest)(nil),        // 21: pb.UpdateOperationRequest
	(*UserResponse)(nil),                  // 22: pb.UserResponse
	(*TradeConfigResponse)(nil),           // 23: pb.TradeConfigResponse
	(*UpdateTradeConfigResponse)(nil),     // 24: pb.UpdateTradeConfigResponse
	(*GetTradeConfigResponse)(nil),        // 25: pb.GetTradeConfigResponse
	(*UserStrategyResponse)(nil),          // 26: pb.UserStrategyResponse
	(*ListParityResponse)(nil),            // 27: pb.ListParityResponse
	(*ListCoinResponse)(nil),              // 28: pb.ListCoinResponse
	(*GetWalletWithCoinResponse)(nil),     // 29: pb.GetWalletWithCoinResponse
	(*ListWalletWithCoinResponse)(nil),    // 30: pb.ListWalletWithCoinResponse
	(*CreateWalletResponse)(nil),          // 31: pb.CreateWalletResponse
	(*UpdateWalletResponse)(nil),          // 32: pb.UpdateWalletResponse
	(*GetCandleFirstMtsResponse)(nil),     // 33: pb.GetCandleFirstMtsResponse
	(*CreateCandlesResponse)(nil),         // 34: pb.CreateCandlesResponse
	(*ListCandleLimitResponse)(nil),       // 35: pb.ListCandleLimitResponse
	(*ListOperationResponse)(nil),         // 36: pb.ListOperationResponse
	(*ListOperationByPeriodResponse)(nil), // 37: pb.ListOperationByPeriodResponse
	(*ListAllOperationResponse)(nil),      // 38: pb.ListAllOperationResponse
	(*UpdateOperationResponse)(nil),       // 39: pb.UpdateOperationResponse
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
	10, // 10: pb.ExchangeService.ListCoin:input_type -> pb.ListCoinRequest
	11, // 11: pb.ExchangeService.GetWalletWithCoin:input_type -> pb.GetWalletWithCoinRequest
	12, // 12: pb.ExchangeService.ListWalletWithCoin:input_type -> pb.ListWalletWithCoinRequest
	13, // 13: pb.ExchangeService.CreateWallet:input_type -> pb.CreateWalletRequest
	14, // 14: pb.ExchangeService.UpdateWallet:input_type -> pb.UpdateWalletRequest
	15, // 15: pb.ExchangeService.GetCandleFirstMts:input_type -> pb.GetCandleFirstMtsRequest
	15, // 16: pb.ExchangeService.GetLastTwoCandles:input_type -> pb.GetCandleFirstMtsRequest
	16, // 17: pb.ExchangeService.CreateCandles:input_type -> pb.CreateCandlesRequest
	17, // 18: pb.ExchangeService.ListCandleLimit:input_type -> pb.ListCandleLimitRequest
	18, // 19: pb.ExchangeService.ListOperation:input_type -> pb.ListOperationRequest
	19, // 20: pb.ExchangeService.ListOperationByPeriod:input_type -> pb.ListOperationByPeriodRequest
	20, // 21: pb.ExchangeService.ListAllOperation:input_type -> pb.ListAllOperationRequest
	21, // 22: pb.ExchangeService.UpdateOperation:input_type -> pb.UpdateOperationRequest
	22, // 23: pb.UserService.CreateUser:output_type -> pb.UserResponse
	22, // 24: pb.UserService.GetUserByEmail:output_type -> pb.UserResponse
	22, // 25: pb.UserService.GetUserById:output_type -> pb.UserResponse
	23, // 26: pb.ExchangeService.ListTradeConfig:output_type -> pb.TradeConfigResponse
	23, // 27: pb.ExchangeService.CreateTradeConfig:output_type -> pb.TradeConfigResponse
	24, // 28: pb.ExchangeService.UpdateTradeConfig:output_type -> pb.UpdateTradeConfigResponse
	25, // 29: pb.ExchangeService.GetTradeConfig:output_type -> pb.GetTradeConfigResponse
	26, // 30: pb.ExchangeService.CreateUserStrategy:output_type -> pb.UserStrategyResponse
	26, // 31: pb.ExchangeService.ListUserStrategy:output_type -> pb.UserStrategyResponse
	27, // 32: pb.ExchangeService.ListParity:output_type -> pb.ListParityResponse
	28, // 33: pb.ExchangeService.ListCoin:output_type -> pb.ListCoinResponse
	29, // 34: pb.ExchangeService.GetWalletWithCoin:output_type -> pb.GetWalletWithCoinResponse
	30, // 35: pb.ExchangeService.ListWalletWithCoin:output_type -> pb.ListWalletWithCoinResponse
	31, // 36: pb.ExchangeService.CreateWallet:output_type -> pb.CreateWalletResponse
	32, // 37: pb.ExchangeService.UpdateWallet:output_type -> pb.UpdateWalletResponse
	33, // 38: pb.ExchangeService.GetCandleFirstMts:output_type -> pb.GetCandleFirstMtsResponse
	33, // 39: pb.ExchangeService.GetLastTwoCandles:output_type -> pb.GetCandleFirstMtsResponse
	34, // 40: pb.ExchangeService.CreateCandles:output_type -> pb.CreateCandlesResponse
	35, // 41: pb.ExchangeService.ListCandleLimit:output_type -> pb.ListCandleLimitResponse
	36, // 42: pb.ExchangeService.ListOperation:output_type -> pb.ListOperationResponse
	37, // 43: pb.ExchangeService.ListOperationByPeriod:output_type -> pb.ListOperationByPeriodResponse
	38, // 44: pb.ExchangeService.ListAllOperation:output_type -> pb.ListAllOperationResponse
	39, // 45: pb.ExchangeService.UpdateOperation:output_type -> pb.UpdateOperationResponse
	23, // [23:46] is the sub-list for method output_type
	0,  // [0:23] is the sub-list for method input_type
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
