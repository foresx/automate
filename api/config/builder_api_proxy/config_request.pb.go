// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/builder_api_proxy/config_request.proto

package bldrapiproxy

import (
	fmt "fmt"
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConfigRequest struct {
	V1                   *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte            `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32             `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest) Reset()         { *m = ConfigRequest{} }
func (m *ConfigRequest) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest) ProtoMessage()    {}
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0}
}

func (m *ConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest.Unmarshal(m, b)
}
func (m *ConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest.Marshal(b, m, deterministic)
}
func (m *ConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest.Merge(m, src)
}
func (m *ConfigRequest) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest.Size(m)
}
func (m *ConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest proto.InternalMessageInfo

func (m *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if m != nil {
		return m.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	Sys                  *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc                  *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                    `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                     `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1) Reset()         { *m = ConfigRequest_V1{} }
func (m *ConfigRequest_V1) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1) ProtoMessage()    {}
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0, 0}
}

func (m *ConfigRequest_V1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1.Unmarshal(m, b)
}
func (m *ConfigRequest_V1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1.Merge(m, src)
}
func (m *ConfigRequest_V1) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1.Size(m)
}
func (m *ConfigRequest_V1) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1 proto.InternalMessageInfo

func (m *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if m != nil {
		return m.Sys
	}
	return nil
}

func (m *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if m != nil {
		return m.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	Mlsa                 *shared.Mlsa                     `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Tls                  *shared.TLSCredentials           `protobuf:"bytes,2,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Service              *ConfigRequest_V1_System_Service `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Log                  *ConfigRequest_V1_System_Logger  `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Nginx                *ConfigRequest_V1_System_Nginx   `protobuf:"bytes,5,opt,name=nginx,proto3" json:"nginx,omitempty" toml:"nginx,omitempty" mapstructure:"nginx,omitempty"`
	Http                 *ConfigRequest_V1_System_HTTP    `protobuf:"bytes,6,opt,name=http,proto3" json:"http,omitempty" toml:"http,omitempty" mapstructure:"http,omitempty"`
	Web                  *ConfigRequest_V1_System_Web     `protobuf:"bytes,7,opt,name=web,proto3" json:"web,omitempty" toml:"web,omitempty" mapstructure:"web,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System) Reset()         { *m = ConfigRequest_V1_System{} }
func (m *ConfigRequest_V1_System) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System) ProtoMessage()    {}
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0, 0, 0}
}

func (m *ConfigRequest_V1_System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System.Merge(m, src)
}
func (m *ConfigRequest_V1_System) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System.Size(m)
}
func (m *ConfigRequest_V1_System) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System proto.InternalMessageInfo

func (m *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if m != nil {
		return m.Mlsa
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if m != nil {
		return m.Tls
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetLog() *ConfigRequest_V1_System_Logger {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetNginx() *ConfigRequest_V1_System_Nginx {
	if m != nil {
		return m.Nginx
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetHttp() *ConfigRequest_V1_System_HTTP {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *ConfigRequest_V1_System) GetWeb() *ConfigRequest_V1_System_Web {
	if m != nil {
		return m.Web
	}
	return nil
}

type ConfigRequest_V1_System_Service struct {
	Host                 *wrappers.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrappers.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Service) Reset()         { *m = ConfigRequest_V1_System_Service{} }
func (m *ConfigRequest_V1_System_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0, 0, 0, 0}
}

func (m *ConfigRequest_V1_System_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Service.Size(m)
}
func (m *ConfigRequest_V1_System_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Service proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Service) GetHost() *wrappers.StringValue {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfigRequest_V1_System_Service) GetPort() *wrappers.Int32Value {
	if m != nil {
		return m.Port
	}
	return nil
}

type ConfigRequest_V1_System_Logger struct {
	Level                *wrappers.StringValue `protobuf:"bytes,1,opt,name=level,proto3" json:"level,omitempty" toml:"level,omitempty" mapstructure:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Logger) Reset()         { *m = ConfigRequest_V1_System_Logger{} }
func (m *ConfigRequest_V1_System_Logger) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Logger) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Logger) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0, 0, 0, 1}
}

func (m *ConfigRequest_V1_System_Logger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Logger.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Logger) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Logger.Size(m)
}
func (m *ConfigRequest_V1_System_Logger) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Logger.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Logger proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Logger) GetLevel() *wrappers.StringValue {
	if m != nil {
		return m.Level
	}
	return nil
}

type ConfigRequest_V1_System_Nginx struct {
	WorkerConnections    *wrappers.Int32Value `protobuf:"bytes,1,opt,name=worker_connections,json=workerConnections,proto3" json:"worker_connections,omitempty" toml:"worker_connections,omitempty" mapstructure:"worker_connections,omitempty"`
	WorkerProcesses      *wrappers.Int32Value `protobuf:"bytes,2,opt,name=worker_processes,json=workerProcesses,proto3" json:"worker_processes,omitempty" toml:"worker_processes,omitempty" mapstructure:"worker_processes,omitempty"`
	WorkerRlimitNofile   *wrappers.Int32Value `protobuf:"bytes,3,opt,name=worker_rlimit_nofile,json=workerRlimitNofile,proto3" json:"worker_rlimit_nofile,omitempty" toml:"worker_rlimit_nofile,omitempty" mapstructure:"worker_rlimit_nofile,omitempty"`
	MaxBodySize          *wrappers.Int32Value `protobuf:"bytes,4,opt,name=max_body_size,json=maxBodySize,proto3" json:"max_body_size,omitempty" toml:"max_body_size,omitempty" mapstructure:"max_body_size,omitempty"`
	ProxySendTimeout     *wrappers.Int32Value `protobuf:"bytes,5,opt,name=proxy_send_timeout,json=proxySendTimeout,proto3" json:"proxy_send_timeout,omitempty" toml:"proxy_send_timeout,omitempty" mapstructure:"proxy_send_timeout,omitempty"`
	ProxyReadTimeout     *wrappers.Int32Value `protobuf:"bytes,6,opt,name=proxy_read_timeout,json=proxyReadTimeout,proto3" json:"proxy_read_timeout,omitempty" toml:"proxy_read_timeout,omitempty" mapstructure:"proxy_read_timeout,omitempty"`
	ProxyConnectTimeout  *wrappers.Int32Value `protobuf:"bytes,7,opt,name=proxy_connect_timeout,json=proxyConnectTimeout,proto3" json:"proxy_connect_timeout,omitempty" toml:"proxy_connect_timeout,omitempty" mapstructure:"proxy_connect_timeout,omitempty"`
	EnableCaching        *wrappers.BoolValue  `protobuf:"bytes,8,opt,name=enable_caching,json=enableCaching,proto3" json:"enable_caching,omitempty" toml:"enable_caching,omitempty" mapstructure:"enable_caching,omitempty"`
	EnableGzip           *wrappers.BoolValue  `protobuf:"bytes,9,opt,name=enable_gzip,json=enableGzip,proto3" json:"enable_gzip,omitempty" toml:"enable_gzip,omitempty" mapstructure:"enable_gzip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Nginx) Reset()         { *m = ConfigRequest_V1_System_Nginx{} }
func (m *ConfigRequest_V1_System_Nginx) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Nginx) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Nginx) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0, 0, 0, 2}
}

func (m *ConfigRequest_V1_System_Nginx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Nginx.Size(m)
}
func (m *ConfigRequest_V1_System_Nginx) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Nginx.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Nginx proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Nginx) GetWorkerConnections() *wrappers.Int32Value {
	if m != nil {
		return m.WorkerConnections
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetWorkerProcesses() *wrappers.Int32Value {
	if m != nil {
		return m.WorkerProcesses
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetWorkerRlimitNofile() *wrappers.Int32Value {
	if m != nil {
		return m.WorkerRlimitNofile
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetMaxBodySize() *wrappers.Int32Value {
	if m != nil {
		return m.MaxBodySize
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetProxySendTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.ProxySendTimeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetProxyReadTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.ProxyReadTimeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetProxyConnectTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.ProxyConnectTimeout
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetEnableCaching() *wrappers.BoolValue {
	if m != nil {
		return m.EnableCaching
	}
	return nil
}

func (m *ConfigRequest_V1_System_Nginx) GetEnableGzip() *wrappers.BoolValue {
	if m != nil {
		return m.EnableGzip
	}
	return nil
}

type ConfigRequest_V1_System_HTTP struct {
	KeepaliveConnections *wrappers.Int32Value `protobuf:"bytes,1,opt,name=keepalive_connections,json=keepaliveConnections,proto3" json:"keepalive_connections,omitempty" toml:"keepalive_connections,omitempty" mapstructure:"keepalive_connections,omitempty"`
	KeepaliveRequests    *wrappers.Int32Value `protobuf:"bytes,2,opt,name=keepalive_requests,json=keepaliveRequests,proto3" json:"keepalive_requests,omitempty" toml:"keepalive_requests,omitempty" mapstructure:"keepalive_requests,omitempty"`
	KeepaliveTimeout     *wrappers.Int32Value `protobuf:"bytes,3,opt,name=keepalive_timeout,json=keepaliveTimeout,proto3" json:"keepalive_timeout,omitempty" toml:"keepalive_timeout,omitempty" mapstructure:"keepalive_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte               `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_HTTP) Reset()         { *m = ConfigRequest_V1_System_HTTP{} }
func (m *ConfigRequest_V1_System_HTTP) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_HTTP) ProtoMessage()    {}
func (*ConfigRequest_V1_System_HTTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0, 0, 0, 3}
}

func (m *ConfigRequest_V1_System_HTTP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_HTTP.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_HTTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_HTTP.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_HTTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_HTTP.Merge(m, src)
}
func (m *ConfigRequest_V1_System_HTTP) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_HTTP.Size(m)
}
func (m *ConfigRequest_V1_System_HTTP) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_HTTP.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_HTTP proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_HTTP) GetKeepaliveConnections() *wrappers.Int32Value {
	if m != nil {
		return m.KeepaliveConnections
	}
	return nil
}

func (m *ConfigRequest_V1_System_HTTP) GetKeepaliveRequests() *wrappers.Int32Value {
	if m != nil {
		return m.KeepaliveRequests
	}
	return nil
}

func (m *ConfigRequest_V1_System_HTTP) GetKeepaliveTimeout() *wrappers.Int32Value {
	if m != nil {
		return m.KeepaliveTimeout
	}
	return nil
}

type ConfigRequest_V1_System_Web struct {
	CookieDomain         *wrappers.StringValue `protobuf:"bytes,1,opt,name=cookie_domain,json=cookieDomain,proto3" json:"cookie_domain,omitempty" toml:"cookie_domain,omitempty" mapstructure:"cookie_domain,omitempty"`
	Environment          *wrappers.StringValue `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty" toml:"environment,omitempty" mapstructure:"environment,omitempty"`
	UseGravatar          *wrappers.BoolValue   `protobuf:"bytes,3,opt,name=use_gravatar,json=useGravatar,proto3" json:"use_gravatar,omitempty" toml:"use_gravatar,omitempty" mapstructure:"use_gravatar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                 `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_System_Web) Reset()         { *m = ConfigRequest_V1_System_Web{} }
func (m *ConfigRequest_V1_System_Web) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_System_Web) ProtoMessage()    {}
func (*ConfigRequest_V1_System_Web) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0, 0, 0, 4}
}

func (m *ConfigRequest_V1_System_Web) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_System_Web.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_System_Web) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_System_Web.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_System_Web) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_System_Web.Merge(m, src)
}
func (m *ConfigRequest_V1_System_Web) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_System_Web.Size(m)
}
func (m *ConfigRequest_V1_System_Web) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_System_Web.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_System_Web proto.InternalMessageInfo

func (m *ConfigRequest_V1_System_Web) GetCookieDomain() *wrappers.StringValue {
	if m != nil {
		return m.CookieDomain
	}
	return nil
}

func (m *ConfigRequest_V1_System_Web) GetEnvironment() *wrappers.StringValue {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *ConfigRequest_V1_System_Web) GetUseGravatar() *wrappers.BoolValue {
	if m != nil {
		return m.UseGravatar
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *ConfigRequest_V1_Service) Reset()         { *m = ConfigRequest_V1_Service{} }
func (m *ConfigRequest_V1_Service) String() string { return proto.CompactTextString(m) }
func (*ConfigRequest_V1_Service) ProtoMessage()    {}
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_051f76e364601840, []int{0, 0, 1}
}

func (m *ConfigRequest_V1_Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigRequest_V1_Service.Unmarshal(m, b)
}
func (m *ConfigRequest_V1_Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigRequest_V1_Service.Marshal(b, m, deterministic)
}
func (m *ConfigRequest_V1_Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigRequest_V1_Service.Merge(m, src)
}
func (m *ConfigRequest_V1_Service) XXX_Size() int {
	return xxx_messageInfo_ConfigRequest_V1_Service.Size(m)
}
func (m *ConfigRequest_V1_Service) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigRequest_V1_Service.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigRequest_V1_Service proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConfigRequest)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest")
	proto.RegisterType((*ConfigRequest_V1)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest.V1")
	proto.RegisterType((*ConfigRequest_V1_System)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest.V1.System")
	proto.RegisterType((*ConfigRequest_V1_System_Service)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest.V1.System.Service")
	proto.RegisterType((*ConfigRequest_V1_System_Logger)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest.V1.System.Logger")
	proto.RegisterType((*ConfigRequest_V1_System_Nginx)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest.V1.System.Nginx")
	proto.RegisterType((*ConfigRequest_V1_System_HTTP)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest.V1.System.HTTP")
	proto.RegisterType((*ConfigRequest_V1_System_Web)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest.V1.System.Web")
	proto.RegisterType((*ConfigRequest_V1_Service)(nil), "chef.automate.domain.builder_api_proxy.ConfigRequest.V1.Service")
}

func init() {
	proto.RegisterFile("api/config/builder_api_proxy/config_request.proto", fileDescriptor_051f76e364601840)
}

var fileDescriptor_051f76e364601840 = []byte{
	// 912 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x96, 0xdf, 0x6e, 0xdb, 0x36,
	0x17, 0xc0, 0xe1, 0xf8, 0x4f, 0xbe, 0x32, 0x4d, 0xeb, 0xb0, 0x2d, 0x20, 0xa8, 0x45, 0x11, 0x7c,
	0x17, 0xc3, 0x30, 0xc0, 0xf2, 0xec, 0xee, 0x62, 0x58, 0xbb, 0x75, 0xb5, 0xb3, 0x26, 0x0d, 0xda,
	0x2c, 0x93, 0xb3, 0xb4, 0xd8, 0x2e, 0x04, 0x4a, 0x3e, 0x96, 0x89, 0x50, 0xa4, 0x46, 0xd2, 0x4e,
	0x9c, 0xbb, 0xbd, 0xc1, 0xd0, 0x47, 0xea, 0xcd, 0x9e, 0x61, 0x2f, 0xb1, 0x9b, 0xde, 0xec, 0x72,
	0x10, 0x49, 0xcb, 0xd9, 0x02, 0xb8, 0x5e, 0x7d, 0x17, 0x85, 0xfc, 0xfd, 0xc4, 0x73, 0x78, 0xce,
	0xb1, 0x50, 0x87, 0xe4, 0xb4, 0x9d, 0x08, 0x3e, 0xa2, 0x69, 0x3b, 0x9e, 0x50, 0x36, 0x04, 0x19,
	0x91, 0x9c, 0x46, 0xb9, 0x14, 0x17, 0x33, 0xb7, 0x10, 0x49, 0xf8, 0x65, 0x02, 0x4a, 0x07, 0xb9,
	0x14, 0x5a, 0xe0, 0x4f, 0x92, 0x31, 0x8c, 0x02, 0x32, 0xd1, 0x22, 0x23, 0x1a, 0x82, 0xa1, 0xc8,
	0x08, 0xe5, 0xc1, 0x35, 0xd8, 0x7f, 0x78, 0x45, 0xad, 0xc6, 0x44, 0xc2, 0xb0, 0x9d, 0x32, 0x11,
	0x13, 0x66, 0x3d, 0xfe, 0xfd, 0xeb, 0xeb, 0x9a, 0x29, 0xb7, 0x78, 0x98, 0x88, 0x2c, 0x17, 0x1c,
	0xb8, 0x56, 0xed, 0xf9, 0xab, 0x5a, 0xa9, 0xcc, 0x93, 0xb6, 0x59, 0x4f, 0x5a, 0x29, 0xf0, 0x16,
	0xe9, 0xb6, 0x1c, 0x5f, 0xa8, 0x48, 0xb7, 0x78, 0x68, 0x13, 0xce, 0x85, 0x26, 0x9a, 0x0a, 0x3e,
	0x77, 0x3d, 0x4c, 0x85, 0x48, 0x19, 0x58, 0x32, 0x9e, 0x8c, 0xda, 0xe7, 0x92, 0xe4, 0x39, 0x48,
	0xb7, 0xfe, 0xff, 0xdf, 0x76, 0xd0, 0x76, 0xdf, 0x78, 0x42, 0x1b, 0x28, 0x3e, 0x40, 0x1b, 0xd3,
	0x8e, 0x57, 0xd9, 0xad, 0x7c, 0xba, 0xd5, 0xfd, 0x32, 0x58, 0x2d, 0xde, 0xe0, 0x1f, 0x8a, 0xe0,
	0xb4, 0x13, 0x6e, 0x4c, 0x3b, 0xfe, 0xdb, 0x26, 0xda, 0x38, 0xed, 0xe0, 0x1f, 0x50, 0x55, 0xcd,
	0x94, 0x33, 0x3e, 0xfd, 0x58, 0x63, 0x30, 0x98, 0x29, 0x0d, 0x59, 0x58, 0xb8, 0x70, 0x88, 0xaa,
	0x6a, 0x9a, 0x78, 0x1b, 0x46, 0xf9, 0xed, 0xc7, 0x2b, 0x41, 0x4e, 0x69, 0x02, 0x61, 0x21, 0xf3,
	0xff, 0xb8, 0x85, 0x1a, 0xf6, 0x1d, 0xf8, 0x0b, 0x54, 0xcb, 0x98, 0x22, 0xee, 0xc8, 0xbb, 0xff,
	0xf2, 0x53, 0x3e, 0x92, 0x24, 0xb0, 0xc9, 0x0f, 0x5e, 0x31, 0x45, 0x42, 0xb3, 0x1b, 0x3f, 0x41,
	0x55, 0xcd, 0x94, 0x3b, 0xd4, 0x67, 0xcb, 0xa0, 0x93, 0x97, 0x83, 0xbe, 0x84, 0x21, 0x70, 0x4d,
	0x09, 0x53, 0x61, 0x81, 0x61, 0x82, 0x36, 0x95, 0x3d, 0x8e, 0x57, 0x35, 0x86, 0xfd, 0x35, 0x33,
	0x55, 0x46, 0x37, 0xf7, 0xe2, 0x37, 0xa8, 0xca, 0x44, 0xea, 0xd5, 0x8c, 0xfe, 0xf9, 0xba, 0xfa,
	0x97, 0x22, 0x4d, 0x41, 0x86, 0x85, 0x12, 0xff, 0x8c, 0xea, 0x3c, 0xa5, 0xfc, 0xc2, 0xab, 0x1b,
	0xf7, 0x77, 0xeb, 0xba, 0x8f, 0x0a, 0x59, 0x68, 0x9d, 0xf8, 0x0d, 0xaa, 0x8d, 0xb5, 0xce, 0xbd,
	0x86, 0x71, 0xef, 0xad, 0xeb, 0x3e, 0x38, 0x39, 0x39, 0x0e, 0x8d, 0x11, 0xff, 0x88, 0xaa, 0xe7,
	0x10, 0x7b, 0x9b, 0x46, 0xdc, 0x5f, 0x57, 0xfc, 0x1a, 0xe2, 0xb0, 0xf0, 0xf9, 0xbf, 0x56, 0xd0,
	0xa6, 0x4b, 0x3e, 0xfe, 0x1c, 0xd5, 0xc6, 0x42, 0x69, 0x57, 0x4a, 0x0f, 0x02, 0xdb, 0x8e, 0xc1,
	0xbc, 0x1d, 0x83, 0x81, 0x96, 0x94, 0xa7, 0xa7, 0x84, 0x4d, 0x20, 0x34, 0x3b, 0xf1, 0x1e, 0xaa,
	0xe5, 0x42, 0x6a, 0x57, 0x47, 0xf7, 0xaf, 0x11, 0x2f, 0xb8, 0x7e, 0xd4, 0x35, 0x40, 0xef, 0xce,
	0xbb, 0xf7, 0xde, 0x6d, 0x9b, 0x9d, 0xe6, 0x5f, 0x47, 0xbe, 0xf9, 0x23, 0x34, 0xb4, 0xff, 0x04,
	0x35, 0xec, 0x05, 0xe1, 0x2e, 0xaa, 0x33, 0x98, 0x02, 0x5b, 0xe9, 0x08, 0x76, 0xab, 0xff, 0xb6,
	0x8e, 0xea, 0xe6, 0x0e, 0xf0, 0x21, 0xc2, 0xe7, 0x42, 0x9e, 0x81, 0x8c, 0x12, 0xc1, 0x39, 0x24,
	0x66, 0xb6, 0x38, 0xd5, 0xb2, 0xb3, 0x85, 0x3b, 0x16, 0xeb, 0x2f, 0x28, 0xfc, 0x1c, 0x35, 0x9d,
	0x2b, 0x97, 0x22, 0x01, 0xa5, 0x40, 0xad, 0x10, 0x65, 0x78, 0xdb, 0x42, 0xc7, 0x73, 0x06, 0xbf,
	0x42, 0x77, 0x9d, 0x47, 0x32, 0x9a, 0x51, 0x1d, 0x71, 0x31, 0xa2, 0x6c, 0xde, 0x37, 0x4b, 0x5d,
	0x2e, 0x98, 0xd0, 0x70, 0x47, 0x06, 0xc3, 0x4f, 0xd1, 0x76, 0x46, 0x2e, 0xa2, 0x58, 0x0c, 0x67,
	0x91, 0xa2, 0x97, 0xe0, 0x1a, 0x64, 0xa9, 0x67, 0x2b, 0x23, 0x17, 0x3d, 0x31, 0x9c, 0x0d, 0xe8,
	0x25, 0xe0, 0x17, 0x08, 0x9b, 0xca, 0x88, 0x14, 0xf0, 0x61, 0xa4, 0x69, 0x06, 0x62, 0xa2, 0x5d,
	0x2b, 0x2c, 0xb5, 0x34, 0x0d, 0x36, 0x00, 0x3e, 0x3c, 0xb1, 0xd0, 0x42, 0x25, 0x81, 0x2c, 0x54,
	0x8d, 0x55, 0x55, 0x21, 0x90, 0x52, 0xf5, 0x3d, 0xba, 0x67, 0x55, 0xee, 0xe2, 0x4a, 0xdb, 0xe6,
	0x87, 0x6d, 0x77, 0x0c, 0xe9, 0xee, 0x6e, 0x2e, 0x7c, 0x86, 0x6e, 0x01, 0x27, 0x31, 0x83, 0x28,
	0x21, 0xc9, 0x98, 0xf2, 0xd4, 0xfb, 0x9f, 0x31, 0xf9, 0xd7, 0x4c, 0x3d, 0x21, 0x98, 0x15, 0x6d,
	0x5b, 0xa2, 0x6f, 0x01, 0xfc, 0x18, 0x6d, 0x39, 0x45, 0x7a, 0x49, 0x73, 0xef, 0xc6, 0x07, 0x79,
	0x64, 0xb7, 0xef, 0x5f, 0xd2, 0xdc, 0xff, 0xb3, 0x82, 0x6a, 0x45, 0xf3, 0xe2, 0x63, 0x74, 0xef,
	0x0c, 0x20, 0x27, 0x8c, 0x4e, 0xe1, 0xbf, 0x96, 0xe5, 0xdd, 0x92, 0xbc, 0x5a, 0x99, 0x87, 0x08,
	0x2f, 0x8c, 0xee, 0x17, 0x7f, 0xa5, 0xda, 0xdc, 0x29, 0x31, 0x37, 0x0f, 0x14, 0x3e, 0x40, 0x8b,
	0x7f, 0x96, 0x39, 0x5f, 0xa1, 0x34, 0x9b, 0x25, 0xe5, 0x12, 0xee, 0xff, 0x5e, 0x41, 0xd5, 0xd7,
	0x10, 0xe3, 0x67, 0x68, 0x3b, 0x11, 0xe2, 0x8c, 0x42, 0x64, 0x67, 0xd2, 0x4a, 0x9d, 0x7c, 0xd3,
	0x22, 0x7b, 0x86, 0xc0, 0xdf, 0x14, 0x89, 0x9f, 0x52, 0x29, 0x78, 0x06, 0x7c, 0x3e, 0x5b, 0x96,
	0x0b, 0xae, 0x02, 0xf8, 0x6b, 0x74, 0x73, 0xa2, 0x20, 0x4a, 0x25, 0x99, 0x12, 0x4d, 0xa4, 0x8b,
	0x67, 0xd9, 0xcd, 0x6d, 0x4d, 0x14, 0xec, 0xbb, 0xed, 0xfe, 0x8d, 0x72, 0x20, 0x7e, 0xb5, 0xfb,
	0xee, 0xbd, 0xf7, 0x00, 0xf9, 0xe5, 0x77, 0x8d, 0x9b, 0xaf, 0x2d, 0x92, 0xd3, 0x96, 0xa9, 0xba,
	0xde, 0xde, 0x4f, 0xbd, 0x94, 0xea, 0xf1, 0x24, 0x0e, 0x12, 0x91, 0xb5, 0x8b, 0xa1, 0x5c, 0x7e,
	0x05, 0xb5, 0x97, 0x7d, 0xb1, 0x3d, 0x8e, 0xd9, 0x50, 0x92, 0x9c, 0x9a, 0x87, 0xb8, 0x61, 0xce,
	0xf4, 0xe8, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x65, 0x0c, 0x9d, 0xe5, 0x09, 0x00, 0x00,
}
