// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: external/cfgmgmt/request/rollouts.proto

package request

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SCMType int32

const (
	SCMType_SCM_TYPE_UNSPECIFIED SCMType = 0
	SCMType_SCM_TYPE_UNKNOWN_SCM SCMType = 1
	SCMType_SCM_TYPE_GIT         SCMType = 2
)

// Enum value maps for SCMType.
var (
	SCMType_name = map[int32]string{
		0: "SCM_TYPE_UNSPECIFIED",
		1: "SCM_TYPE_UNKNOWN_SCM",
		2: "SCM_TYPE_GIT",
	}
	SCMType_value = map[string]int32{
		"SCM_TYPE_UNSPECIFIED": 0,
		"SCM_TYPE_UNKNOWN_SCM": 1,
		"SCM_TYPE_GIT":         2,
	}
)

func (x SCMType) Enum() *SCMType {
	p := new(SCMType)
	*p = x
	return p
}

func (x SCMType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SCMType) Descriptor() protoreflect.EnumDescriptor {
	return file_external_cfgmgmt_request_rollouts_proto_enumTypes[0].Descriptor()
}

func (SCMType) Type() protoreflect.EnumType {
	return &file_external_cfgmgmt_request_rollouts_proto_enumTypes[0]
}

func (x SCMType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SCMType.Descriptor instead.
func (SCMType) EnumDescriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP(), []int{0}
}

type SCMWebType int32

const (
	SCMWebType_SCM_WEB_TYPE_UNSPECIFIED     SCMWebType = 0
	SCMWebType_SCM_WEB_TYPE_UNKNOWN_SCM_WEB SCMWebType = 1
	SCMWebType_SCM_WEB_TYPE_GITHUB          SCMWebType = 2
)

// Enum value maps for SCMWebType.
var (
	SCMWebType_name = map[int32]string{
		0: "SCM_WEB_TYPE_UNSPECIFIED",
		1: "SCM_WEB_TYPE_UNKNOWN_SCM_WEB",
		2: "SCM_WEB_TYPE_GITHUB",
	}
	SCMWebType_value = map[string]int32{
		"SCM_WEB_TYPE_UNSPECIFIED":     0,
		"SCM_WEB_TYPE_UNKNOWN_SCM_WEB": 1,
		"SCM_WEB_TYPE_GITHUB":          2,
	}
)

func (x SCMWebType) Enum() *SCMWebType {
	p := new(SCMWebType)
	*p = x
	return p
}

func (x SCMWebType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SCMWebType) Descriptor() protoreflect.EnumDescriptor {
	return file_external_cfgmgmt_request_rollouts_proto_enumTypes[1].Descriptor()
}

func (SCMWebType) Type() protoreflect.EnumType {
	return &file_external_cfgmgmt_request_rollouts_proto_enumTypes[1]
}

func (x SCMWebType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SCMWebType.Descriptor instead.
func (SCMWebType) EnumDescriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP(), []int{1}
}

// CreateRollout is a request to create a new Rollout. All
// fields have the same meaning as with the response Rollout
// type.
type CreateRollout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyName           string     `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyNodeGroup      string     `protobuf:"bytes,2,opt,name=policy_node_group,json=policyNodeGroup,proto3" json:"policy_node_group,omitempty"`
	PolicyRevisionId     string     `protobuf:"bytes,3,opt,name=policy_revision_id,json=policyRevisionId,proto3" json:"policy_revision_id,omitempty"`
	PolicyDomainUrl      string     `protobuf:"bytes,4,opt,name=policy_domain_url,json=policyDomainUrl,proto3" json:"policy_domain_url,omitempty"`
	ScmType              SCMType    `protobuf:"varint,5,opt,name=scm_type,json=scmType,proto3,enum=chef.automate.api.cfgmgmt.request.SCMType" json:"scm_type,omitempty"`
	ScmWebType           SCMWebType `protobuf:"varint,6,opt,name=scm_web_type,json=scmWebType,proto3,enum=chef.automate.api.cfgmgmt.request.SCMWebType" json:"scm_web_type,omitempty"`
	PolicyScmUrl         string     `protobuf:"bytes,7,opt,name=policy_scm_url,json=policyScmUrl,proto3" json:"policy_scm_url,omitempty"`
	PolicyScmWebUrl      string     `protobuf:"bytes,8,opt,name=policy_scm_web_url,json=policyScmWebUrl,proto3" json:"policy_scm_web_url,omitempty"`
	PolicyScmCommit      string     `protobuf:"bytes,9,opt,name=policy_scm_commit,json=policyScmCommit,proto3" json:"policy_scm_commit,omitempty"`
	Description          string     `protobuf:"bytes,10,opt,name=description,proto3" json:"description,omitempty"`
	CiJobUrl             string     `protobuf:"bytes,11,opt,name=ci_job_url,json=ciJobUrl,proto3" json:"ci_job_url,omitempty"`
	CiJobId              string     `protobuf:"bytes,12,opt,name=ci_job_id,json=ciJobId,proto3" json:"ci_job_id,omitempty"`
	ScmAuthorName        string     `protobuf:"bytes,13,opt,name=scm_author_name,json=scmAuthorName,proto3" json:"scm_author_name,omitempty"`
	ScmAuthorEmail       string     `protobuf:"bytes,14,opt,name=scm_author_email,json=scmAuthorEmail,proto3" json:"scm_author_email,omitempty"`
	PolicyDomainUsername string     `protobuf:"bytes,15,opt,name=policy_domain_username,json=policyDomainUsername,proto3" json:"policy_domain_username,omitempty"`
}

func (x *CreateRollout) Reset() {
	*x = CreateRollout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRollout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRollout) ProtoMessage() {}

func (x *CreateRollout) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRollout.ProtoReflect.Descriptor instead.
func (*CreateRollout) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRollout) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *CreateRollout) GetPolicyNodeGroup() string {
	if x != nil {
		return x.PolicyNodeGroup
	}
	return ""
}

func (x *CreateRollout) GetPolicyRevisionId() string {
	if x != nil {
		return x.PolicyRevisionId
	}
	return ""
}

func (x *CreateRollout) GetPolicyDomainUrl() string {
	if x != nil {
		return x.PolicyDomainUrl
	}
	return ""
}

func (x *CreateRollout) GetScmType() SCMType {
	if x != nil {
		return x.ScmType
	}
	return SCMType_SCM_TYPE_UNSPECIFIED
}

func (x *CreateRollout) GetScmWebType() SCMWebType {
	if x != nil {
		return x.ScmWebType
	}
	return SCMWebType_SCM_WEB_TYPE_UNSPECIFIED
}

func (x *CreateRollout) GetPolicyScmUrl() string {
	if x != nil {
		return x.PolicyScmUrl
	}
	return ""
}

func (x *CreateRollout) GetPolicyScmWebUrl() string {
	if x != nil {
		return x.PolicyScmWebUrl
	}
	return ""
}

func (x *CreateRollout) GetPolicyScmCommit() string {
	if x != nil {
		return x.PolicyScmCommit
	}
	return ""
}

func (x *CreateRollout) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateRollout) GetCiJobUrl() string {
	if x != nil {
		return x.CiJobUrl
	}
	return ""
}

func (x *CreateRollout) GetCiJobId() string {
	if x != nil {
		return x.CiJobId
	}
	return ""
}

func (x *CreateRollout) GetScmAuthorName() string {
	if x != nil {
		return x.ScmAuthorName
	}
	return ""
}

func (x *CreateRollout) GetScmAuthorEmail() string {
	if x != nil {
		return x.ScmAuthorEmail
	}
	return ""
}

func (x *CreateRollout) GetPolicyDomainUsername() string {
	if x != nil {
		return x.PolicyDomainUsername
	}
	return ""
}

type Rollouts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filters to apply to the request for the rollouts list.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
}

func (x *Rollouts) Reset() {
	*x = Rollouts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rollouts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rollouts) ProtoMessage() {}

func (x *Rollouts) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rollouts.ProtoReflect.Descriptor instead.
func (*Rollouts) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP(), []int{1}
}

func (x *Rollouts) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

type RolloutById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RolloutId string `protobuf:"bytes,1,opt,name=rollout_id,json=rolloutId,proto3" json:"rollout_id,omitempty"`
}

func (x *RolloutById) Reset() {
	*x = RolloutById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolloutById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolloutById) ProtoMessage() {}

func (x *RolloutById) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolloutById.ProtoReflect.Descriptor instead.
func (*RolloutById) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP(), []int{2}
}

func (x *RolloutById) GetRolloutId() string {
	if x != nil {
		return x.RolloutId
	}
	return ""
}

type RolloutForChefRun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PolicyName       string `protobuf:"bytes,1,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	PolicyGroup      string `protobuf:"bytes,2,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	PolicyRevisionId string `protobuf:"bytes,3,opt,name=policy_revision_id,json=policyRevisionId,proto3" json:"policy_revision_id,omitempty"`
}

func (x *RolloutForChefRun) Reset() {
	*x = RolloutForChefRun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RolloutForChefRun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RolloutForChefRun) ProtoMessage() {}

func (x *RolloutForChefRun) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RolloutForChefRun.ProtoReflect.Descriptor instead.
func (*RolloutForChefRun) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP(), []int{3}
}

func (x *RolloutForChefRun) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *RolloutForChefRun) GetPolicyGroup() string {
	if x != nil {
		return x.PolicyGroup
	}
	return ""
}

func (x *RolloutForChefRun) GetPolicyRevisionId() string {
	if x != nil {
		return x.PolicyRevisionId
	}
	return ""
}

type ListNodeSegmentsWithRolloutProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filters to apply to the request for the node segments list.
	Filter []string `protobuf:"bytes,1,rep,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListNodeSegmentsWithRolloutProgress) Reset() {
	*x = ListNodeSegmentsWithRolloutProgress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeSegmentsWithRolloutProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeSegmentsWithRolloutProgress) ProtoMessage() {}

func (x *ListNodeSegmentsWithRolloutProgress) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeSegmentsWithRolloutProgress.ProtoReflect.Descriptor instead.
func (*ListNodeSegmentsWithRolloutProgress) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP(), []int{4}
}

func (x *ListNodeSegmentsWithRolloutProgress) GetFilter() []string {
	if x != nil {
		return x.Filter
	}
	return nil
}

type CreateRolloutTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateRolloutTest) Reset() {
	*x = CreateRolloutTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRolloutTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRolloutTest) ProtoMessage() {}

func (x *CreateRolloutTest) ProtoReflect() protoreflect.Message {
	mi := &file_external_cfgmgmt_request_rollouts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRolloutTest.ProtoReflect.Descriptor instead.
func (*CreateRolloutTest) Descriptor() ([]byte, []int) {
	return file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP(), []int{5}
}

var File_external_cfgmgmt_request_rollouts_proto protoreflect.FileDescriptor

var file_external_cfgmgmt_request_rollouts_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67,
	0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x72, 0x6f, 0x6c, 0x6c, 0x6f,
	0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66, 0x67,
	0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb1, 0x05, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x12, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x45, 0x0a, 0x08, 0x73, 0x63, 0x6d, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66, 0x67, 0x6d,
	0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x43, 0x4d, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x07, 0x73, 0x63, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4f, 0x0a, 0x0c,
	0x73, 0x63, 0x6d, 0x5f, 0x77, 0x65, 0x62, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x43, 0x4d, 0x57, 0x65, 0x62, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x73, 0x63, 0x6d, 0x57, 0x65, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x63, 0x6d, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x63, 0x6d,
	0x55, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x12, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x63,
	0x6d, 0x5f, 0x77, 0x65, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x63, 0x6d, 0x57, 0x65, 0x62, 0x55, 0x72, 0x6c,
	0x12, 0x2a, 0x0a, 0x11, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x73, 0x63, 0x6d, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x53, 0x63, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x0a, 0x63, 0x69, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x4a, 0x6f, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x09,
	0x63, 0x69, 0x5f, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x69, 0x4a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x63, 0x6d, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x63, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x10, 0x73, 0x63, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x63, 0x6d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x34, 0x0a, 0x16, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x22, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74,
	0x49, 0x64, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x46, 0x6f,
	0x72, 0x43, 0x68, 0x65, 0x66, 0x52, 0x75, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x2c, 0x0a, 0x12,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x23, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x57, 0x69,
	0x74, 0x68, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x6c, 0x6f, 0x75, 0x74, 0x54, 0x65, 0x73, 0x74, 0x2a, 0x4f,
	0x0a, 0x07, 0x53, 0x43, 0x4d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x43, 0x4d,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x43, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x43, 0x4d, 0x10, 0x01, 0x12, 0x10, 0x0a,
	0x0c, 0x53, 0x43, 0x4d, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x49, 0x54, 0x10, 0x02, 0x2a,
	0x65, 0x0a, 0x0a, 0x53, 0x43, 0x4d, 0x57, 0x65, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x18, 0x53, 0x43, 0x4d, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x53,
	0x43, 0x4d, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x53, 0x43, 0x4d, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x01, 0x12, 0x17, 0x0a,
	0x13, 0x53, 0x43, 0x4d, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47, 0x49,
	0x54, 0x48, 0x55, 0x42, 0x10, 0x02, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_cfgmgmt_request_rollouts_proto_rawDescOnce sync.Once
	file_external_cfgmgmt_request_rollouts_proto_rawDescData = file_external_cfgmgmt_request_rollouts_proto_rawDesc
)

func file_external_cfgmgmt_request_rollouts_proto_rawDescGZIP() []byte {
	file_external_cfgmgmt_request_rollouts_proto_rawDescOnce.Do(func() {
		file_external_cfgmgmt_request_rollouts_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_cfgmgmt_request_rollouts_proto_rawDescData)
	})
	return file_external_cfgmgmt_request_rollouts_proto_rawDescData
}

var file_external_cfgmgmt_request_rollouts_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_external_cfgmgmt_request_rollouts_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_external_cfgmgmt_request_rollouts_proto_goTypes = []interface{}{
	(SCMType)(0),                                // 0: chef.automate.api.cfgmgmt.request.SCMType
	(SCMWebType)(0),                             // 1: chef.automate.api.cfgmgmt.request.SCMWebType
	(*CreateRollout)(nil),                       // 2: chef.automate.api.cfgmgmt.request.CreateRollout
	(*Rollouts)(nil),                            // 3: chef.automate.api.cfgmgmt.request.Rollouts
	(*RolloutById)(nil),                         // 4: chef.automate.api.cfgmgmt.request.RolloutById
	(*RolloutForChefRun)(nil),                   // 5: chef.automate.api.cfgmgmt.request.RolloutForChefRun
	(*ListNodeSegmentsWithRolloutProgress)(nil), // 6: chef.automate.api.cfgmgmt.request.ListNodeSegmentsWithRolloutProgress
	(*CreateRolloutTest)(nil),                   // 7: chef.automate.api.cfgmgmt.request.CreateRolloutTest
}
var file_external_cfgmgmt_request_rollouts_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.cfgmgmt.request.CreateRollout.scm_type:type_name -> chef.automate.api.cfgmgmt.request.SCMType
	1, // 1: chef.automate.api.cfgmgmt.request.CreateRollout.scm_web_type:type_name -> chef.automate.api.cfgmgmt.request.SCMWebType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_external_cfgmgmt_request_rollouts_proto_init() }
func file_external_cfgmgmt_request_rollouts_proto_init() {
	if File_external_cfgmgmt_request_rollouts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_cfgmgmt_request_rollouts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRollout); i {
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
		file_external_cfgmgmt_request_rollouts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rollouts); i {
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
		file_external_cfgmgmt_request_rollouts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolloutById); i {
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
		file_external_cfgmgmt_request_rollouts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RolloutForChefRun); i {
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
		file_external_cfgmgmt_request_rollouts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNodeSegmentsWithRolloutProgress); i {
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
		file_external_cfgmgmt_request_rollouts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRolloutTest); i {
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
			RawDescriptor: file_external_cfgmgmt_request_rollouts_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_cfgmgmt_request_rollouts_proto_goTypes,
		DependencyIndexes: file_external_cfgmgmt_request_rollouts_proto_depIdxs,
		EnumInfos:         file_external_cfgmgmt_request_rollouts_proto_enumTypes,
		MessageInfos:      file_external_cfgmgmt_request_rollouts_proto_msgTypes,
	}.Build()
	File_external_cfgmgmt_request_rollouts_proto = out.File
	file_external_cfgmgmt_request_rollouts_proto_rawDesc = nil
	file_external_cfgmgmt_request_rollouts_proto_goTypes = nil
	file_external_cfgmgmt_request_rollouts_proto_depIdxs = nil
}
