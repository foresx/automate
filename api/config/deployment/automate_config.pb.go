// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/config/deployment/automate_config.proto

package deployment

import (
	fmt "fmt"
	applications "github.com/chef/automate/api/config/applications"
	authn "github.com/chef/automate/api/config/authn"
	authz "github.com/chef/automate/api/config/authz"
	backup_gateway "github.com/chef/automate/api/config/backup_gateway"
	bifrost "github.com/chef/automate/api/config/bifrost"
	bookshelf "github.com/chef/automate/api/config/bookshelf"
	builder_api "github.com/chef/automate/api/config/builder_api"
	builder_api_proxy "github.com/chef/automate/api/config/builder_api_proxy"
	builder_memcached "github.com/chef/automate/api/config/builder_memcached"
	cereal "github.com/chef/automate/api/config/cereal"
	cfgmgmt "github.com/chef/automate/api/config/cfgmgmt"
	compliance "github.com/chef/automate/api/config/compliance"
	cs_nginx "github.com/chef/automate/api/config/cs_nginx"
	data_feed "github.com/chef/automate/api/config/data_feed"
	dex "github.com/chef/automate/api/config/dex"
	elasticsearch "github.com/chef/automate/api/config/elasticsearch"
	erchef "github.com/chef/automate/api/config/erchef"
	es_sidecar "github.com/chef/automate/api/config/es_sidecar"
	esgateway "github.com/chef/automate/api/config/esgateway"
	event "github.com/chef/automate/api/config/event"
	event_feed "github.com/chef/automate/api/config/event_feed"
	event_gateway "github.com/chef/automate/api/config/event_gateway"
	gateway "github.com/chef/automate/api/config/gateway"
	ingest "github.com/chef/automate/api/config/ingest"
	license_control "github.com/chef/automate/api/config/license_control"
	load_balancer "github.com/chef/automate/api/config/load_balancer"
	local_user "github.com/chef/automate/api/config/local_user"
	minio "github.com/chef/automate/api/config/minio"
	nodemanager "github.com/chef/automate/api/config/nodemanager"
	notifications "github.com/chef/automate/api/config/notifications"
	pg_gateway "github.com/chef/automate/api/config/pg_gateway"
	pg_sidecar "github.com/chef/automate/api/config/pg_sidecar"
	postgresql "github.com/chef/automate/api/config/postgresql"
	prometheus "github.com/chef/automate/api/config/prometheus"
	secrets "github.com/chef/automate/api/config/secrets"
	session "github.com/chef/automate/api/config/session"
	shared "github.com/chef/automate/api/config/shared"
	teams "github.com/chef/automate/api/config/teams"
	ui "github.com/chef/automate/api/config/ui"
	workflow_nginx "github.com/chef/automate/api/config/workflow_nginx"
	workflow_server "github.com/chef/automate/api/config/workflow_server"
	proto "github.com/golang/protobuf/proto"
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

type AutomateConfig struct {
	Global               *shared.GlobalConfig             `protobuf:"bytes,19,opt,name=global,proto3" json:"global,omitempty" toml:"global,omitempty" mapstructure:"global,omitempty"`
	AuthN                *authn.ConfigRequest             `protobuf:"bytes,1,opt,name=auth_n,json=authN,proto3" json:"auth_n,omitempty" toml:"auth_n,omitempty" mapstructure:"auth_n,omitempty"`
	AuthZ                *authz.ConfigRequest             `protobuf:"bytes,2,opt,name=auth_z,json=authZ,proto3" json:"auth_z,omitempty" toml:"auth_z,omitempty" mapstructure:"auth_z,omitempty"`
	Compliance           *compliance.ConfigRequest        `protobuf:"bytes,10,opt,name=compliance,proto3" json:"compliance,omitempty" toml:"compliance,omitempty" mapstructure:"compliance,omitempty"`
	ConfigMgmt           *cfgmgmt.ConfigRequest           `protobuf:"bytes,6,opt,name=config_mgmt,json=configMgmt,proto3" json:"config_mgmt,omitempty" toml:"config_mgmt,omitempty" mapstructure:"config_mgmt,omitempty"`
	Deployment           *ConfigRequest                   `protobuf:"bytes,3,opt,name=deployment,proto3" json:"deployment,omitempty" toml:"deployment,omitempty" mapstructure:"deployment,omitempty"`
	Dex                  *dex.ConfigRequest               `protobuf:"bytes,4,opt,name=dex,proto3" json:"dex,omitempty" toml:"dex,omitempty" mapstructure:"dex,omitempty"`
	Elasticsearch        *elasticsearch.ConfigRequest     `protobuf:"bytes,7,opt,name=elasticsearch,proto3" json:"elasticsearch,omitempty" toml:"elasticsearch,omitempty" mapstructure:"elasticsearch,omitempty"`
	Esgateway            *esgateway.ConfigRequest         `protobuf:"bytes,31,opt,name=esgateway,proto3" json:"esgateway,omitempty" toml:"esgateway,omitempty" mapstructure:"esgateway,omitempty"`
	EsSidecar            *es_sidecar.ConfigRequest        `protobuf:"bytes,11,opt,name=es_sidecar,json=esSidecar,proto3" json:"es_sidecar,omitempty" toml:"es_sidecar,omitempty" mapstructure:"es_sidecar,omitempty"`
	Gateway              *gateway.ConfigRequest           `protobuf:"bytes,5,opt,name=gateway,proto3" json:"gateway,omitempty" toml:"gateway,omitempty" mapstructure:"gateway,omitempty"`
	Ingest               *ingest.ConfigRequest            `protobuf:"bytes,13,opt,name=ingest,proto3" json:"ingest,omitempty" toml:"ingest,omitempty" mapstructure:"ingest,omitempty"`
	LoadBalancer         *load_balancer.ConfigRequest     `protobuf:"bytes,8,opt,name=load_balancer,json=loadBalancer,proto3" json:"load_balancer,omitempty" toml:"load_balancer,omitempty" mapstructure:"load_balancer,omitempty"`
	LocalUser            *local_user.ConfigRequest        `protobuf:"bytes,12,opt,name=local_user,json=localUser,proto3" json:"local_user,omitempty" toml:"local_user,omitempty" mapstructure:"local_user,omitempty"`
	LicenseControl       *license_control.ConfigRequest   `protobuf:"bytes,16,opt,name=license_control,json=licenseControl,proto3" json:"license_control,omitempty" toml:"license_control,omitempty" mapstructure:"license_control,omitempty"`
	Notifications        *notifications.ConfigRequest     `protobuf:"bytes,14,opt,name=notifications,proto3" json:"notifications,omitempty" toml:"notifications,omitempty" mapstructure:"notifications,omitempty"`
	Postgresql           *postgresql.ConfigRequest        `protobuf:"bytes,15,opt,name=postgresql,proto3" json:"postgresql,omitempty" toml:"postgresql,omitempty" mapstructure:"postgresql,omitempty"`
	Session              *session.ConfigRequest           `protobuf:"bytes,17,opt,name=session,proto3" json:"session,omitempty" toml:"session,omitempty" mapstructure:"session,omitempty"`
	Teams                *teams.ConfigRequest             `protobuf:"bytes,18,opt,name=teams,proto3" json:"teams,omitempty" toml:"teams,omitempty" mapstructure:"teams,omitempty"`
	UI                   *ui.ConfigRequest                `protobuf:"bytes,9,opt,name=u_i,json=uI,proto3" json:"u_i,omitempty" toml:"u_i,omitempty" mapstructure:"u_i,omitempty"`
	Secrets              *secrets.ConfigRequest           `protobuf:"bytes,21,opt,name=secrets,proto3" json:"secrets,omitempty" toml:"secrets,omitempty" mapstructure:"secrets,omitempty"`
	BackupGateway        *backup_gateway.ConfigRequest    `protobuf:"bytes,29,opt,name=backup_gateway,json=backupGateway,proto3" json:"backup_gateway,omitempty" toml:"backup_gateway,omitempty" mapstructure:"backup_gateway,omitempty"`
	PgSidecar            *pg_sidecar.ConfigRequest        `protobuf:"bytes,35,opt,name=pg_sidecar,json=pgSidecar,proto3" json:"pg_sidecar,omitempty" toml:"pg_sidecar,omitempty" mapstructure:"pg_sidecar,omitempty"`
	PgGateway            *pg_gateway.ConfigRequest        `protobuf:"bytes,34,opt,name=pg_gateway,json=pgGateway,proto3" json:"pg_gateway,omitempty" toml:"pg_gateway,omitempty" mapstructure:"pg_gateway,omitempty"`
	Applications         *applications.ConfigRequest      `protobuf:"bytes,36,opt,name=applications,proto3" json:"applications,omitempty" toml:"applications,omitempty" mapstructure:"applications,omitempty"`
	Bookshelf            *bookshelf.ConfigRequest         `protobuf:"bytes,22,opt,name=bookshelf,proto3" json:"bookshelf,omitempty" toml:"bookshelf,omitempty" mapstructure:"bookshelf,omitempty"`
	Bifrost              *bifrost.ConfigRequest           `protobuf:"bytes,23,opt,name=bifrost,proto3" json:"bifrost,omitempty" toml:"bifrost,omitempty" mapstructure:"bifrost,omitempty"`
	Erchef               *erchef.ConfigRequest            `protobuf:"bytes,24,opt,name=erchef,proto3" json:"erchef,omitempty" toml:"erchef,omitempty" mapstructure:"erchef,omitempty"`
	CsNginx              *cs_nginx.ConfigRequest          `protobuf:"bytes,25,opt,name=cs_nginx,json=csNginx,proto3" json:"cs_nginx,omitempty" toml:"cs_nginx,omitempty" mapstructure:"cs_nginx,omitempty"`
	Workflow             *workflow_server.ConfigRequest   `protobuf:"bytes,27,opt,name=workflow,proto3" json:"workflow,omitempty" toml:"workflow,omitempty" mapstructure:"workflow,omitempty"`
	WorkflowNginx        *workflow_nginx.ConfigRequest    `protobuf:"bytes,28,opt,name=workflow_nginx,json=workflowNginx,proto3" json:"workflow_nginx,omitempty" toml:"workflow_nginx,omitempty" mapstructure:"workflow_nginx,omitempty"`
	EventService         *event.ConfigRequest             `protobuf:"bytes,30,opt,name=event_service,json=eventService,proto3" json:"event_service,omitempty" toml:"event_service,omitempty" mapstructure:"event_service,omitempty"`
	Nodemanager          *nodemanager.ConfigRequest       `protobuf:"bytes,33,opt,name=nodemanager,proto3" json:"nodemanager,omitempty" toml:"nodemanager,omitempty" mapstructure:"nodemanager,omitempty"`
	EventGateway         *event_gateway.ConfigRequest     `protobuf:"bytes,37,opt,name=event_gateway,json=eventGateway,proto3" json:"event_gateway,omitempty" toml:"event_gateway,omitempty" mapstructure:"event_gateway,omitempty"`
	Prometheus           *prometheus.ConfigRequest        `protobuf:"bytes,32,opt,name=prometheus,proto3" json:"prometheus,omitempty" toml:"prometheus,omitempty" mapstructure:"prometheus,omitempty"`
	DataFeedService      *data_feed.ConfigRequest         `protobuf:"bytes,38,opt,name=data_feed_service,json=dataFeedService,proto3" json:"data_feed_service,omitempty" toml:"data_feed_service,omitempty" mapstructure:"data_feed_service,omitempty"`
	EventFeedService     *event_feed.ConfigRequest        `protobuf:"bytes,39,opt,name=event_feed_service,json=eventFeedService,proto3" json:"event_feed_service,omitempty" toml:"event_feed_service,omitempty" mapstructure:"event_feed_service,omitempty"`
	Cereal               *cereal.ConfigRequest            `protobuf:"bytes,40,opt,name=cereal,proto3" json:"cereal,omitempty" toml:"cereal,omitempty" mapstructure:"cereal,omitempty"`
	BuilderApi           *builder_api.ConfigRequest       `protobuf:"bytes,41,opt,name=builder_api,json=builderApi,proto3" json:"builder_api,omitempty" toml:"builder_api,omitempty" mapstructure:"builder_api,omitempty"`
	BuilderApiProxy      *builder_api_proxy.ConfigRequest `protobuf:"bytes,42,opt,name=builder_api_proxy,json=builderApiProxy,proto3" json:"builder_api_proxy,omitempty" toml:"builder_api_proxy,omitempty" mapstructure:"builder_api_proxy,omitempty"`
	Minio                *minio.ConfigRequest             `protobuf:"bytes,43,opt,name=minio,proto3" json:"minio,omitempty" toml:"minio,omitempty" mapstructure:"minio,omitempty"`
	BuilderMemcached     *builder_memcached.ConfigRequest `protobuf:"bytes,44,opt,name=builder_memcached,json=builderMemcached,proto3" json:"builder_memcached,omitempty" toml:"builder_memcached,omitempty" mapstructure:"builder_memcached,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte                           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32                            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *AutomateConfig) Reset()         { *m = AutomateConfig{} }
func (m *AutomateConfig) String() string { return proto.CompactTextString(m) }
func (*AutomateConfig) ProtoMessage()    {}
func (*AutomateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff33eb1f2117d94c, []int{0}
}

func (m *AutomateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutomateConfig.Unmarshal(m, b)
}
func (m *AutomateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutomateConfig.Marshal(b, m, deterministic)
}
func (m *AutomateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutomateConfig.Merge(m, src)
}
func (m *AutomateConfig) XXX_Size() int {
	return xxx_messageInfo_AutomateConfig.Size(m)
}
func (m *AutomateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AutomateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AutomateConfig proto.InternalMessageInfo

func (m *AutomateConfig) GetGlobal() *shared.GlobalConfig {
	if m != nil {
		return m.Global
	}
	return nil
}

func (m *AutomateConfig) GetAuthN() *authn.ConfigRequest {
	if m != nil {
		return m.AuthN
	}
	return nil
}

func (m *AutomateConfig) GetAuthZ() *authz.ConfigRequest {
	if m != nil {
		return m.AuthZ
	}
	return nil
}

func (m *AutomateConfig) GetCompliance() *compliance.ConfigRequest {
	if m != nil {
		return m.Compliance
	}
	return nil
}

func (m *AutomateConfig) GetConfigMgmt() *cfgmgmt.ConfigRequest {
	if m != nil {
		return m.ConfigMgmt
	}
	return nil
}

func (m *AutomateConfig) GetDeployment() *ConfigRequest {
	if m != nil {
		return m.Deployment
	}
	return nil
}

func (m *AutomateConfig) GetDex() *dex.ConfigRequest {
	if m != nil {
		return m.Dex
	}
	return nil
}

func (m *AutomateConfig) GetElasticsearch() *elasticsearch.ConfigRequest {
	if m != nil {
		return m.Elasticsearch
	}
	return nil
}

func (m *AutomateConfig) GetEsgateway() *esgateway.ConfigRequest {
	if m != nil {
		return m.Esgateway
	}
	return nil
}

func (m *AutomateConfig) GetEsSidecar() *es_sidecar.ConfigRequest {
	if m != nil {
		return m.EsSidecar
	}
	return nil
}

func (m *AutomateConfig) GetGateway() *gateway.ConfigRequest {
	if m != nil {
		return m.Gateway
	}
	return nil
}

func (m *AutomateConfig) GetIngest() *ingest.ConfigRequest {
	if m != nil {
		return m.Ingest
	}
	return nil
}

func (m *AutomateConfig) GetLoadBalancer() *load_balancer.ConfigRequest {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

func (m *AutomateConfig) GetLocalUser() *local_user.ConfigRequest {
	if m != nil {
		return m.LocalUser
	}
	return nil
}

func (m *AutomateConfig) GetLicenseControl() *license_control.ConfigRequest {
	if m != nil {
		return m.LicenseControl
	}
	return nil
}

func (m *AutomateConfig) GetNotifications() *notifications.ConfigRequest {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func (m *AutomateConfig) GetPostgresql() *postgresql.ConfigRequest {
	if m != nil {
		return m.Postgresql
	}
	return nil
}

func (m *AutomateConfig) GetSession() *session.ConfigRequest {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *AutomateConfig) GetTeams() *teams.ConfigRequest {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *AutomateConfig) GetUI() *ui.ConfigRequest {
	if m != nil {
		return m.UI
	}
	return nil
}

func (m *AutomateConfig) GetSecrets() *secrets.ConfigRequest {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *AutomateConfig) GetBackupGateway() *backup_gateway.ConfigRequest {
	if m != nil {
		return m.BackupGateway
	}
	return nil
}

func (m *AutomateConfig) GetPgSidecar() *pg_sidecar.ConfigRequest {
	if m != nil {
		return m.PgSidecar
	}
	return nil
}

func (m *AutomateConfig) GetPgGateway() *pg_gateway.ConfigRequest {
	if m != nil {
		return m.PgGateway
	}
	return nil
}

func (m *AutomateConfig) GetApplications() *applications.ConfigRequest {
	if m != nil {
		return m.Applications
	}
	return nil
}

func (m *AutomateConfig) GetBookshelf() *bookshelf.ConfigRequest {
	if m != nil {
		return m.Bookshelf
	}
	return nil
}

func (m *AutomateConfig) GetBifrost() *bifrost.ConfigRequest {
	if m != nil {
		return m.Bifrost
	}
	return nil
}

func (m *AutomateConfig) GetErchef() *erchef.ConfigRequest {
	if m != nil {
		return m.Erchef
	}
	return nil
}

func (m *AutomateConfig) GetCsNginx() *cs_nginx.ConfigRequest {
	if m != nil {
		return m.CsNginx
	}
	return nil
}

func (m *AutomateConfig) GetWorkflow() *workflow_server.ConfigRequest {
	if m != nil {
		return m.Workflow
	}
	return nil
}

func (m *AutomateConfig) GetWorkflowNginx() *workflow_nginx.ConfigRequest {
	if m != nil {
		return m.WorkflowNginx
	}
	return nil
}

func (m *AutomateConfig) GetEventService() *event.ConfigRequest {
	if m != nil {
		return m.EventService
	}
	return nil
}

func (m *AutomateConfig) GetNodemanager() *nodemanager.ConfigRequest {
	if m != nil {
		return m.Nodemanager
	}
	return nil
}

func (m *AutomateConfig) GetEventGateway() *event_gateway.ConfigRequest {
	if m != nil {
		return m.EventGateway
	}
	return nil
}

func (m *AutomateConfig) GetPrometheus() *prometheus.ConfigRequest {
	if m != nil {
		return m.Prometheus
	}
	return nil
}

func (m *AutomateConfig) GetDataFeedService() *data_feed.ConfigRequest {
	if m != nil {
		return m.DataFeedService
	}
	return nil
}

func (m *AutomateConfig) GetEventFeedService() *event_feed.ConfigRequest {
	if m != nil {
		return m.EventFeedService
	}
	return nil
}

func (m *AutomateConfig) GetCereal() *cereal.ConfigRequest {
	if m != nil {
		return m.Cereal
	}
	return nil
}

func (m *AutomateConfig) GetBuilderApi() *builder_api.ConfigRequest {
	if m != nil {
		return m.BuilderApi
	}
	return nil
}

func (m *AutomateConfig) GetBuilderApiProxy() *builder_api_proxy.ConfigRequest {
	if m != nil {
		return m.BuilderApiProxy
	}
	return nil
}

func (m *AutomateConfig) GetMinio() *minio.ConfigRequest {
	if m != nil {
		return m.Minio
	}
	return nil
}

func (m *AutomateConfig) GetBuilderMemcached() *builder_memcached.ConfigRequest {
	if m != nil {
		return m.BuilderMemcached
	}
	return nil
}

func init() {
	proto.RegisterType((*AutomateConfig)(nil), "chef.automate.domain.deployment.AutomateConfig")
}

func init() {
	proto.RegisterFile("api/config/deployment/automate_config.proto", fileDescriptor_ff33eb1f2117d94c)
}

var fileDescriptor_ff33eb1f2117d94c = []byte{
	// 1278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x98, 0x5b, 0x73, 0xdc, 0x34,
	0x14, 0xc7, 0xa7, 0xcd, 0x5d, 0xc9, 0xe6, 0x22, 0x28, 0xa8, 0x29, 0x34, 0x25, 0xa5, 0xcd, 0xdd,
	0x9b, 0xb6, 0xf0, 0xc0, 0x0b, 0x90, 0x74, 0x20, 0x43, 0x68, 0x42, 0x69, 0x09, 0xd0, 0x4e, 0x67,
	0x3c, 0x5a, 0x5b, 0xeb, 0xd5, 0xc4, 0xb6, 0x5c, 0x5f, 0x9a, 0xcb, 0x97, 0xe6, 0x2b, 0x30, 0x96,
	0x7c, 0x91, 0xb5, 0xb2, 0xbc, 0x8f, 0xd1, 0xf9, 0xe9, 0x9f, 0xe3, 0x73, 0x74, 0xce, 0x91, 0x16,
	0xec, 0xe1, 0x88, 0xf6, 0x1d, 0x16, 0x0e, 0xa9, 0xd7, 0x77, 0x49, 0xe4, 0xb3, 0x9b, 0x80, 0x84,
	0x69, 0x1f, 0x67, 0x29, 0x0b, 0x70, 0x4a, 0x6c, 0x61, 0xb2, 0xa2, 0x98, 0xa5, 0x0c, 0x6e, 0x38,
	0x23, 0x32, 0xb4, 0x4a, 0x9b, 0xe5, 0xb2, 0x00, 0xd3, 0xd0, 0xaa, 0xb7, 0xad, 0xef, 0x4b, 0x6a,
	0x38, 0x8a, 0x7c, 0xea, 0xe0, 0x94, 0xb2, 0x30, 0x29, 0xd6, 0xec, 0x98, 0x7c, 0xcc, 0x48, 0x92,
	0x0a, 0xb9, 0xf5, 0x27, 0x32, 0x9d, 0xa5, 0xa3, 0x70, 0x32, 0xec, 0x56, 0x8f, 0x59, 0x12, 0x36,
	0xc0, 0xce, 0x65, 0x16, 0xd9, 0x1e, 0x4e, 0xc9, 0x15, 0xbe, 0xd1, 0xf3, 0x5b, 0x32, 0x4f, 0x87,
	0x31, 0x4b, 0x52, 0x3d, 0xb8, 0x23, 0x83, 0x8c, 0x5d, 0x26, 0x23, 0xe2, 0x0f, 0xf5, 0xe8, 0x53,
	0x09, 0x25, 0x71, 0x1e, 0xad, 0xee, 0x4f, 0x22, 0x9f, 0xf2, 0x80, 0x77, 0xba, 0xe8, 0x0c, 0xbd,
	0xc0, 0x0b, 0x5a, 0xc0, 0x5d, 0x19, 0x64, 0x41, 0xe4, 0x53, 0x1c, 0x3a, 0x44, 0xcf, 0x3e, 0x94,
	0xd8, 0x64, 0x84, 0x63, 0xe2, 0xf6, 0x3d, 0x9f, 0x0d, 0xb0, 0x5f, 0xd8, 0xb7, 0x65, 0xad, 0xc4,
	0x0e, 0x3d, 0x1a, 0x5e, 0x77, 0xff, 0x57, 0xe9, 0xec, 0x68, 0xd9, 0xc7, 0x0d, 0xb6, 0x45, 0xf0,
	0x40, 0x0e, 0x8b, 0x8f, 0x93, 0x94, 0x3a, 0x09, 0xc1, 0xb1, 0x33, 0xea, 0x4e, 0x0c, 0x49, 0x8c,
	0xc9, 0xde, 0x6d, 0xa0, 0x76, 0x42, 0x5d, 0xe2, 0xe0, 0xb8, 0x3b, 0xea, 0x46, 0x51, 0x39, 0xdb,
	0x34, 0xf4, 0x48, 0xdb, 0x01, 0xea, 0x4b, 0x9c, 0x4f, 0x1d, 0x12, 0x26, 0xbc, 0xae, 0xd2, 0x98,
	0xf9, 0xdd, 0x71, 0xf0, 0x19, 0x76, 0xed, 0x01, 0xf6, 0xf3, 0x84, 0xc6, 0xdd, 0x1f, 0xe7, 0x33,
	0x07, 0xfb, 0x76, 0x96, 0xb4, 0xb1, 0x72, 0xbd, 0x87, 0xcc, 0x25, 0x01, 0x0e, 0xb1, 0xd7, 0x06,
	0x1f, 0x34, 0xe0, 0x94, 0x0e, 0xcd, 0xf5, 0x2c, 0xfb, 0x11, 0xb1, 0x24, 0xf5, 0x62, 0x92, 0x7c,
	0xf4, 0x27, 0x60, 0x3d, 0x7b, 0xe2, 0xe4, 0x45, 0x9e, 0x39, 0x79, 0x0d, 0x36, 0x66, 0x01, 0x49,
	0x47, 0x24, 0x4b, 0xba, 0x13, 0x9d, 0x10, 0x27, 0x26, 0xe9, 0x44, 0x60, 0x92, 0x50, 0x36, 0x41,
	0xab, 0x4a, 0x09, 0x0e, 0x5a, 0xf4, 0x36, 0x25, 0x2c, 0xa3, 0xdd, 0xed, 0xec, 0x8a, 0xc5, 0x97,
	0x43, 0x9f, 0x5d, 0x99, 0x8a, 0xb1, 0xaf, 0xe3, 0x13, 0x12, 0x7f, 0x9a, 0x24, 0xb9, 0xbc, 0x07,
	0x99, 0x93, 0x20, 0x17, 0x9b, 0x8b, 0x53, 0x6c, 0x0f, 0x09, 0x71, 0x27, 0x28, 0x36, 0xae, 0xdc,
	0xce, 0xca, 0x35, 0xe4, 0x90, 0x98, 0x60, 0xbf, 0xfb, 0xdc, 0x0e, 0x32, 0xea, 0xbb, 0x24, 0xb6,
	0xeb, 0x25, 0x05, 0x7e, 0xa6, 0x87, 0xed, 0x28, 0x66, 0xd7, 0x37, 0xdd, 0x99, 0x0b, 0x68, 0x48,
	0xd9, 0xe4, 0xca, 0x01, 0x09, 0x1c, 0xec, 0x8c, 0x5a, 0xbe, 0x70, 0xf3, 0xbf, 0xfb, 0x60, 0xf9,
	0xa8, 0x18, 0x99, 0x2f, 0x39, 0x00, 0x7f, 0x06, 0xb3, 0xa2, 0xe5, 0xa2, 0xcf, 0x1e, 0xdd, 0xd9,
	0x5e, 0x7c, 0xbe, 0x6d, 0x35, 0x07, 0x2b, 0x0d, 0x87, 0x31, 0xb6, 0x8a, 0xd1, 0x7b, 0xc2, 0x49,
	0xb1, 0xf3, 0x4d, 0xb1, 0x2f, 0x57, 0xc8, 0x47, 0xa1, 0x1d, 0xa2, 0x3b, 0x5c, 0x61, 0xc7, 0xd2,
	0x8e, 0x66, 0x3e, 0x55, 0xad, 0x62, 0xaf, 0xf0, 0xea, 0xcd, 0x4c, 0xbe, 0x78, 0x5e, 0x29, 0xdc,
	0xa2, 0xbb, 0x5d, 0x0a, 0xb7, 0x3a, 0x85, 0xf7, 0xf0, 0x02, 0x80, 0x7a, 0xd6, 0x20, 0xc0, 0x55,
	0xbe, 0xd7, 0xab, 0xd4, 0x5c, 0xf9, 0x55, 0x4d, 0x45, 0x49, 0x08, 0xbe, 0x02, 0x8b, 0x45, 0x1c,
	0xf3, 0x79, 0x87, 0x66, 0xb9, 0xee, 0x5e, 0x8b, 0xae, 0x18, 0x8a, 0xe3, 0x6a, 0xf9, 0x9f, 0x67,
	0x5e, 0x90, 0xc2, 0x73, 0x00, 0xea, 0xd1, 0x84, 0xa6, 0xb8, 0x98, 0x65, 0x75, 0xdc, 0x63, 0x54,
	0xbd, 0xda, 0x02, 0x7f, 0x00, 0x53, 0x2e, 0xb9, 0x46, 0xd3, 0x5c, 0x68, 0xab, 0x4d, 0xe8, 0x5a,
	0x51, 0xc8, 0xf7, 0xc0, 0xbf, 0x41, 0xaf, 0x31, 0xd4, 0xd0, 0x1c, 0x17, 0x39, 0xd4, 0x26, 0xbf,
	0x41, 0x2a, 0x6a, 0x4d, 0x19, 0x78, 0x0a, 0x16, 0xaa, 0xe9, 0x87, 0x36, 0xb8, 0xe6, 0xbe, 0x5e,
	0xb3, 0xa4, 0x14, 0xbd, 0x7a, 0x3b, 0x7c, 0x05, 0x40, 0x3d, 0x1e, 0xd1, 0x22, 0x17, 0x3b, 0x68,
	0x11, 0x2b, 0xb1, 0x71, 0xb5, 0xb7, 0xc2, 0x00, 0x8f, 0xc0, 0x5c, 0xe9, 0xd7, 0x8c, 0x36, 0x60,
	0x38, 0xa2, 0xfa, 0x03, 0x51, 0xee, 0x83, 0xc7, 0x60, 0x56, 0x8c, 0x56, 0xd4, 0xe3, 0x0a, 0xbb,
	0xfa, 0x90, 0x0b, 0x46, 0x11, 0x29, 0x76, 0xc2, 0x0b, 0xd0, 0x6b, 0x4c, 0x51, 0x34, 0x6f, 0x08,
	0x7c, 0x83, 0x54, 0x04, 0x97, 0x72, 0xe3, 0x71, 0x61, 0x83, 0x67, 0x00, 0xd4, 0xd3, 0x16, 0x2d,
	0x99, 0x8e, 0x56, 0xcd, 0xa9, 0xc1, 0xe2, 0x96, 0x8b, 0x84, 0xc4, 0xf0, 0x03, 0x58, 0x51, 0x2e,
	0x07, 0x68, 0x95, 0x6b, 0xbe, 0x68, 0xd1, 0x6c, 0xc2, 0x8a, 0xf0, 0x72, 0x61, 0x7e, 0x29, 0xac,
	0xf0, 0x1f, 0xd0, 0x6b, 0x4c, 0x70, 0xb4, 0xcc, 0xb5, 0x9f, 0xe9, 0xb5, 0x1b, 0xa8, 0x7a, 0xfa,
	0x1a, 0xc6, 0x3c, 0x0a, 0xf5, 0xac, 0x47, 0x2b, 0x86, 0x13, 0x53, 0x63, 0x6a, 0x7d, 0xd5, 0x16,
	0xf8, 0x0b, 0x98, 0x2b, 0x26, 0x2c, 0x5a, 0x33, 0x55, 0x7e, 0x01, 0xa9, 0xc7, 0xa6, 0x58, 0x86,
	0x3f, 0x81, 0x19, 0x3e, 0x7f, 0x11, 0x34, 0x35, 0x37, 0x8e, 0xa8, 0xcd, 0x8d, 0x2f, 0xc2, 0x43,
	0x30, 0x95, 0xd9, 0x14, 0x2d, 0xf0, 0xed, 0x1b, 0xca, 0xf6, 0x8c, 0x2a, 0x9b, 0xee, 0x66, 0xbf,
	0x09, 0xcf, 0xf9, 0x25, 0x02, 0xdd, 0x33, 0x7b, 0xce, 0xa1, 0x71, 0xcf, 0xf9, 0x32, 0xfc, 0x17,
	0x2c, 0x37, 0x5f, 0x2f, 0xe8, 0x6b, 0x6d, 0xa6, 0x44, 0x4c, 0x9b, 0xa8, 0x9a, 0x29, 0x61, 0x3d,
	0xa9, 0x6b, 0xbb, 0xbe, 0x3d, 0xa1, 0xc7, 0xa6, 0x4c, 0x79, 0x6d, 0xb5, 0x1d, 0x79, 0x65, 0x6d,
	0x0b, 0xb5, 0xd2, 0xc7, 0x4d, 0xb3, 0x5a, 0x4b, 0xdf, 0x89, 0xbc, 0xd2, 0xb7, 0xbf, 0xc0, 0x92,
	0xfc, 0x5e, 0x44, 0xdf, 0x6a, 0x2b, 0xb4, 0x9c, 0x49, 0x12, 0xa9, 0x56, 0xa8, 0x6c, 0x83, 0xbf,
	0x83, 0x85, 0xea, 0xc1, 0x86, 0xbe, 0xd0, 0xba, 0x58, 0x48, 0x56, 0x98, 0xea, 0x62, 0x65, 0xc8,
	0xf3, 0x5b, 0x3c, 0x13, 0xd1, 0x97, 0xa6, 0xfc, 0x16, 0x90, 0x9a, 0xdf, 0x62, 0x39, 0x6f, 0x68,
	0xe2, 0x65, 0x88, 0x90, 0xa9, 0xa1, 0x09, 0x46, 0x6d, 0x68, 0x62, 0x15, 0x9e, 0x80, 0xf9, 0xf2,
	0x65, 0x86, 0xee, 0x6b, 0x1b, 0x7e, 0x39, 0x1f, 0x0b, 0x4a, 0x75, 0xc6, 0x49, 0xce, 0xf3, 0x65,
	0xf8, 0x07, 0x98, 0x2f, 0xef, 0x8a, 0xe8, 0x81, 0xa9, 0xd9, 0x28, 0x37, 0x4a, 0x45, 0xaf, 0x12,
	0x81, 0xef, 0xc0, 0x72, 0xf3, 0xb2, 0x8a, 0xbe, 0xe2, 0xb2, 0xcf, 0x3b, 0x64, 0x75, 0x5e, 0xf6,
	0x4a, 0xab, 0xf0, 0xf5, 0x1c, 0xf4, 0xc4, 0x65, 0x32, 0x77, 0x81, 0x3a, 0x04, 0x3d, 0x34, 0x95,
	0x36, 0x47, 0xd5, 0xc3, 0xc1, 0x17, 0xdf, 0x8a, 0xed, 0xf0, 0x4f, 0xb0, 0x28, 0x3d, 0x80, 0xd0,
	0x37, 0x5c, 0xad, 0xdf, 0xd6, 0x0f, 0x2b, 0x50, 0xd1, 0x94, 0x35, 0xf8, 0x84, 0x97, 0x6f, 0xd2,
	0xe8, 0x89, 0xa9, 0xc9, 0x36, 0x50, 0xad, 0xab, 0x65, 0x75, 0xe4, 0x3d, 0xb6, 0x7a, 0xcb, 0xa0,
	0x47, 0xa6, 0x5a, 0xab, 0xb0, 0xb1, 0x1e, 0x5b, 0x59, 0xe0, 0x3b, 0xb0, 0x56, 0xdd, 0xe0, 0xab,
	0x68, 0x3e, 0x35, 0x95, 0x47, 0x85, 0x2b, 0xaa, 0x2b, 0xb9, 0xe1, 0x57, 0x42, 0xdc, 0x32, 0xa8,
	0x1f, 0x00, 0xac, 0x6f, 0xfc, 0x95, 0xf6, 0x96, 0x69, 0x36, 0xd6, 0xbc, 0x22, 0xbe, 0xca, 0x2d,
	0xb2, 0xfa, 0x11, 0x98, 0x15, 0x6f, 0x04, 0xb4, 0xad, 0xcd, 0x7d, 0x71, 0x6f, 0xe6, 0x88, 0x5a,
	0x3a, 0x62, 0x15, 0xbe, 0x06, 0x8b, 0xd2, 0x8b, 0x00, 0xed, 0x98, 0xb2, 0x2e, 0x81, 0x6a, 0x34,
	0x0b, 0xd3, 0x51, 0x44, 0x21, 0x06, 0x6b, 0x63, 0x6f, 0x0c, 0xb4, 0x6b, 0xba, 0x0d, 0x8f, 0xe1,
	0x6a, 0x54, 0x6b, 0xf5, 0xd7, 0xb9, 0x19, 0xfe, 0x08, 0x66, 0xf8, 0x9b, 0x04, 0xed, 0x19, 0x9e,
	0x0b, 0x9c, 0x50, 0x87, 0x19, 0x5f, 0x94, 0x5d, 0xac, 0x1e, 0x2b, 0x68, 0x9f, 0x6b, 0x7d, 0xa7,
	0x1f, 0x2b, 0x2a, 0xad, 0xa6, 0xa6, 0x00, 0xce, 0x4a, 0xfb, 0xe9, 0xf4, 0xfc, 0xfa, 0xea, 0x83,
	0xd3, 0xe9, 0xf9, 0xcf, 0x57, 0xef, 0x1d, 0x1f, 0xbe, 0xb7, 0x3c, 0x9a, 0x8e, 0xb2, 0x41, 0x7e,
	0xef, 0xef, 0xf3, 0x5f, 0xc1, 0x4a, 0xfd, 0xbe, 0xf6, 0x27, 0xa3, 0xc1, 0x2c, 0x7f, 0x2a, 0xbd,
	0xf8, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x03, 0x1a, 0x6d, 0x8e, 0x14, 0x00, 0x00,
}
