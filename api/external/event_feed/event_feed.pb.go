// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/event_feed/event_feed.proto

package event_feed

import (
	context "context"
	_ "github.com/chef/automate/api/external/common"
	request "github.com/chef/automate/api/external/event_feed/request"
	response "github.com/chef/automate/api/external/event_feed/response"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-policy/iam"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_external_event_feed_event_feed_proto protoreflect.FileDescriptor

var file_external_event_feed_event_feed_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x66, 0x65, 0x65, 0x64, 0x1a, 0x27, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa2, 0x08,
	0x0a, 0x10, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xca, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46,
	0x65, 0x65, 0x64, 0x12, 0x39, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46,
	0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x42, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x66, 0x65, 0x65, 0x64, 0x8a, 0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0xe4, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3f, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18,
	0x0e, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x8a,
	0xb5, 0x18, 0x13, 0x12, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xe4, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3f, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x73,
	0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x41,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x61, 0x73, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x4a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x12, 0xe8, 0x01,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x42, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x45, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x30, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x8a,
	0xb5, 0x18, 0x0e, 0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x87, 0x01, 0x0a, 0x0b, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x38, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_external_event_feed_event_feed_proto_goTypes = []interface{}{
	(*request.GetEventFeedRequest)(nil),            // 0: chef.automate.api.event_feed.request.GetEventFeedRequest
	(*request.GetEventTypeCountsRequest)(nil),      // 1: chef.automate.api.event_feed.request.GetEventTypeCountsRequest
	(*request.GetEventTaskCountsRequest)(nil),      // 2: chef.automate.api.event_feed.request.GetEventTaskCountsRequest
	(*request.GetEventStringBucketsRequest)(nil),   // 3: chef.automate.api.event_feed.request.GetEventStringBucketsRequest
	(*request.EventExportRequest)(nil),             // 4: chef.automate.api.event_feed.request.EventExportRequest
	(*response.GetEventFeedResponse)(nil),          // 5: chef.automate.api.event_feed.response.GetEventFeedResponse
	(*response.GetEventTypeCountsResponse)(nil),    // 6: chef.automate.api.event_feed.response.GetEventTypeCountsResponse
	(*response.GetEventTaskCountsResponse)(nil),    // 7: chef.automate.api.event_feed.response.GetEventTaskCountsResponse
	(*response.GetEventStringBucketsResponse)(nil), // 8: chef.automate.api.event_feed.response.GetEventStringBucketsResponse
	(*response.EventExportResponse)(nil),           // 9: chef.automate.api.event_feed.response.EventExportResponse
}
var file_external_event_feed_event_feed_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.event_feed.EventFeedService.GetEventFeed:input_type -> chef.automate.api.event_feed.request.GetEventFeedRequest
	1, // 1: chef.automate.api.event_feed.EventFeedService.GetEventTypeCounts:input_type -> chef.automate.api.event_feed.request.GetEventTypeCountsRequest
	2, // 2: chef.automate.api.event_feed.EventFeedService.GetEventTaskCounts:input_type -> chef.automate.api.event_feed.request.GetEventTaskCountsRequest
	3, // 3: chef.automate.api.event_feed.EventFeedService.GetEventStringBuckets:input_type -> chef.automate.api.event_feed.request.GetEventStringBucketsRequest
	4, // 4: chef.automate.api.event_feed.EventFeedService.EventExport:input_type -> chef.automate.api.event_feed.request.EventExportRequest
	5, // 5: chef.automate.api.event_feed.EventFeedService.GetEventFeed:output_type -> chef.automate.api.event_feed.response.GetEventFeedResponse
	6, // 6: chef.automate.api.event_feed.EventFeedService.GetEventTypeCounts:output_type -> chef.automate.api.event_feed.response.GetEventTypeCountsResponse
	7, // 7: chef.automate.api.event_feed.EventFeedService.GetEventTaskCounts:output_type -> chef.automate.api.event_feed.response.GetEventTaskCountsResponse
	8, // 8: chef.automate.api.event_feed.EventFeedService.GetEventStringBuckets:output_type -> chef.automate.api.event_feed.response.GetEventStringBucketsResponse
	9, // 9: chef.automate.api.event_feed.EventFeedService.EventExport:output_type -> chef.automate.api.event_feed.response.EventExportResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_event_feed_event_feed_proto_init() }
func file_external_event_feed_event_feed_proto_init() {
	if File_external_event_feed_event_feed_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_external_event_feed_event_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_event_feed_event_feed_proto_goTypes,
		DependencyIndexes: file_external_event_feed_event_feed_proto_depIdxs,
	}.Build()
	File_external_event_feed_event_feed_proto = out.File
	file_external_event_feed_event_feed_proto_rawDesc = nil
	file_external_event_feed_event_feed_proto_goTypes = nil
	file_external_event_feed_event_feed_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventFeedServiceClient is the client API for EventFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventFeedServiceClient interface {
	//
	//List Events
	//
	//Returns a list of recorded events in Chef Automate, such as Infra Server events (cookbook creation, policyfile updates, and node creation) and Chef Automate internal events (profile installs and scan job creation).
	//Adding a filter makes a list of all events that meet the filter criteria.
	//
	//Example:
	//```
	//eventfeed?collapse=true&filter=organization:The%2520Watchmen&page_size=100&start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventFeed(ctx context.Context, in *request.GetEventFeedRequest, opts ...grpc.CallOption) (*response.GetEventFeedResponse, error)
	//
	//List Count of Event Type Occurrence
	//
	//Returns totals for role, cookbook, and organization event types.
	//
	//Example:
	//```
	//event_type_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTypeCounts(ctx context.Context, in *request.GetEventTypeCountsRequest, opts ...grpc.CallOption) (*response.GetEventTypeCountsResponse, error)
	//
	//List Counts Per Event Task Occurrence
	//
	//Returns totals for update, create, and delete event tasks, which are the actions taken on the event.
	//
	//Example:
	//```
	//event_task_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTaskCounts(ctx context.Context, in *request.GetEventTaskCountsRequest, opts ...grpc.CallOption) (*response.GetEventTaskCountsResponse, error)
	//
	//List Summary Stats About Events
	//
	//Returns data that populates the guitar strings visual on the top of the event feed.
	//
	//Example:
	//```
	//eventstrings?timezone=America/Denver&hours_between=1&start=2020-06-19&end=2020-06-25
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventStringBuckets(ctx context.Context, in *request.GetEventStringBucketsRequest, opts ...grpc.CallOption) (*response.GetEventStringBucketsResponse, error)
	EventExport(ctx context.Context, in *request.EventExportRequest, opts ...grpc.CallOption) (EventFeedService_EventExportClient, error)
}

type eventFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventFeedServiceClient(cc grpc.ClientConnInterface) EventFeedServiceClient {
	return &eventFeedServiceClient{cc}
}

func (c *eventFeedServiceClient) GetEventFeed(ctx context.Context, in *request.GetEventFeedRequest, opts ...grpc.CallOption) (*response.GetEventFeedResponse, error) {
	out := new(response.GetEventFeedResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeedService/GetEventFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedServiceClient) GetEventTypeCounts(ctx context.Context, in *request.GetEventTypeCountsRequest, opts ...grpc.CallOption) (*response.GetEventTypeCountsResponse, error) {
	out := new(response.GetEventTypeCountsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeedService/GetEventTypeCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedServiceClient) GetEventTaskCounts(ctx context.Context, in *request.GetEventTaskCountsRequest, opts ...grpc.CallOption) (*response.GetEventTaskCountsResponse, error) {
	out := new(response.GetEventTaskCountsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeedService/GetEventTaskCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedServiceClient) GetEventStringBuckets(ctx context.Context, in *request.GetEventStringBucketsRequest, opts ...grpc.CallOption) (*response.GetEventStringBucketsResponse, error) {
	out := new(response.GetEventStringBucketsResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.event_feed.EventFeedService/GetEventStringBuckets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventFeedServiceClient) EventExport(ctx context.Context, in *request.EventExportRequest, opts ...grpc.CallOption) (EventFeedService_EventExportClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventFeedService_serviceDesc.Streams[0], "/chef.automate.api.event_feed.EventFeedService/EventExport", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventFeedServiceEventExportClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventFeedService_EventExportClient interface {
	Recv() (*response.EventExportResponse, error)
	grpc.ClientStream
}

type eventFeedServiceEventExportClient struct {
	grpc.ClientStream
}

func (x *eventFeedServiceEventExportClient) Recv() (*response.EventExportResponse, error) {
	m := new(response.EventExportResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventFeedServiceServer is the server API for EventFeedService service.
type EventFeedServiceServer interface {
	//
	//List Events
	//
	//Returns a list of recorded events in Chef Automate, such as Infra Server events (cookbook creation, policyfile updates, and node creation) and Chef Automate internal events (profile installs and scan job creation).
	//Adding a filter makes a list of all events that meet the filter criteria.
	//
	//Example:
	//```
	//eventfeed?collapse=true&filter=organization:The%2520Watchmen&page_size=100&start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventFeed(context.Context, *request.GetEventFeedRequest) (*response.GetEventFeedResponse, error)
	//
	//List Count of Event Type Occurrence
	//
	//Returns totals for role, cookbook, and organization event types.
	//
	//Example:
	//```
	//event_type_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTypeCounts(context.Context, *request.GetEventTypeCountsRequest) (*response.GetEventTypeCountsResponse, error)
	//
	//List Counts Per Event Task Occurrence
	//
	//Returns totals for update, create, and delete event tasks, which are the actions taken on the event.
	//
	//Example:
	//```
	//event_task_counts?start=1592546400000&end=1593151199999
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventTaskCounts(context.Context, *request.GetEventTaskCountsRequest) (*response.GetEventTaskCountsResponse, error)
	//
	//List Summary Stats About Events
	//
	//Returns data that populates the guitar strings visual on the top of the event feed.
	//
	//Example:
	//```
	//eventstrings?timezone=America/Denver&hours_between=1&start=2020-06-19&end=2020-06-25
	//```
	//
	//Authorization Action:
	//```
	//event:events:list
	//```
	GetEventStringBuckets(context.Context, *request.GetEventStringBucketsRequest) (*response.GetEventStringBucketsResponse, error)
	EventExport(*request.EventExportRequest, EventFeedService_EventExportServer) error
}

// UnimplementedEventFeedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEventFeedServiceServer struct {
}

func (*UnimplementedEventFeedServiceServer) GetEventFeed(context.Context, *request.GetEventFeedRequest) (*response.GetEventFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventFeed not implemented")
}
func (*UnimplementedEventFeedServiceServer) GetEventTypeCounts(context.Context, *request.GetEventTypeCountsRequest) (*response.GetEventTypeCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTypeCounts not implemented")
}
func (*UnimplementedEventFeedServiceServer) GetEventTaskCounts(context.Context, *request.GetEventTaskCountsRequest) (*response.GetEventTaskCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventTaskCounts not implemented")
}
func (*UnimplementedEventFeedServiceServer) GetEventStringBuckets(context.Context, *request.GetEventStringBucketsRequest) (*response.GetEventStringBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventStringBuckets not implemented")
}
func (*UnimplementedEventFeedServiceServer) EventExport(*request.EventExportRequest, EventFeedService_EventExportServer) error {
	return status.Errorf(codes.Unimplemented, "method EventExport not implemented")
}

func RegisterEventFeedServiceServer(s *grpc.Server, srv EventFeedServiceServer) {
	s.RegisterService(&_EventFeedService_serviceDesc, srv)
}

func _EventFeedService_GetEventFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetEventFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServiceServer).GetEventFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeedService/GetEventFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServiceServer).GetEventFeed(ctx, req.(*request.GetEventFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeedService_GetEventTypeCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetEventTypeCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServiceServer).GetEventTypeCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeedService/GetEventTypeCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServiceServer).GetEventTypeCounts(ctx, req.(*request.GetEventTypeCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeedService_GetEventTaskCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetEventTaskCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServiceServer).GetEventTaskCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeedService/GetEventTaskCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServiceServer).GetEventTaskCounts(ctx, req.(*request.GetEventTaskCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeedService_GetEventStringBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetEventStringBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventFeedServiceServer).GetEventStringBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.event_feed.EventFeedService/GetEventStringBuckets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventFeedServiceServer).GetEventStringBuckets(ctx, req.(*request.GetEventStringBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventFeedService_EventExport_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(request.EventExportRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventFeedServiceServer).EventExport(m, &eventFeedServiceEventExportServer{stream})
}

type EventFeedService_EventExportServer interface {
	Send(*response.EventExportResponse) error
	grpc.ServerStream
}

type eventFeedServiceEventExportServer struct {
	grpc.ServerStream
}

func (x *eventFeedServiceEventExportServer) Send(m *response.EventExportResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _EventFeedService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.event_feed.EventFeedService",
	HandlerType: (*EventFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEventFeed",
			Handler:    _EventFeedService_GetEventFeed_Handler,
		},
		{
			MethodName: "GetEventTypeCounts",
			Handler:    _EventFeedService_GetEventTypeCounts_Handler,
		},
		{
			MethodName: "GetEventTaskCounts",
			Handler:    _EventFeedService_GetEventTaskCounts_Handler,
		},
		{
			MethodName: "GetEventStringBuckets",
			Handler:    _EventFeedService_GetEventStringBuckets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EventExport",
			Handler:       _EventFeedService_EventExport_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "external/event_feed/event_feed.proto",
}
