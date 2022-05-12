// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: config/report_manager/config_request.proto

package report_manager

import (
	shared "github.com/chef/automate/api/config/shared"
	_ "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V1 *ConfigRequest_V1 `protobuf:"bytes,1,opt,name=v1,proto3" json:"v1,omitempty" toml:"v1,omitempty" mapstructure:"v1,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_config_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_config_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_config_report_manager_config_request_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigRequest) GetV1() *ConfigRequest_V1 {
	if x != nil {
		return x.V1
	}
	return nil
}

type ConfigRequest_V1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys *ConfigRequest_V1_System  `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty" toml:"sys,omitempty" mapstructure:"sys,omitempty"`
	Svc *ConfigRequest_V1_Service `protobuf:"bytes,2,opt,name=svc,proto3" json:"svc,omitempty" toml:"svc,omitempty" mapstructure:"svc,omitempty"`
}

func (x *ConfigRequest_V1) Reset() {
	*x = ConfigRequest_V1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_config_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1) ProtoMessage() {}

func (x *ConfigRequest_V1) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_config_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1) Descriptor() ([]byte, []int) {
	return file_config_report_manager_config_request_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfigRequest_V1) GetSys() *ConfigRequest_V1_System {
	if x != nil {
		return x.Sys
	}
	return nil
}

func (x *ConfigRequest_V1) GetSvc() *ConfigRequest_V1_Service {
	if x != nil {
		return x.Svc
	}
	return nil
}

type ConfigRequest_V1_System struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mlsa     *shared.Mlsa                      `protobuf:"bytes,1,opt,name=mlsa,proto3" json:"mlsa,omitempty" toml:"mlsa,omitempty" mapstructure:"mlsa,omitempty"`
	Service  *ConfigRequest_V1_System_Service  `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty" toml:"service,omitempty" mapstructure:"service,omitempty"`
	Tls      *shared.TLSCredentials            `protobuf:"bytes,3,opt,name=tls,proto3" json:"tls,omitempty" toml:"tls,omitempty" mapstructure:"tls,omitempty"`
	Log      *shared.Log                       `protobuf:"bytes,4,opt,name=log,proto3" json:"log,omitempty" toml:"log,omitempty" mapstructure:"log,omitempty"`
	Objstore *ConfigRequest_V1_System_Objstore `protobuf:"bytes,5,opt,name=objstore,proto3" json:"objstore,omitempty" toml:"objstore,omitempty" mapstructure:"objstore,omitempty"`
	Minio    *ConfigRequest_V1_System_Minio    `protobuf:"bytes,6,opt,name=minio,proto3" json:"minio,omitempty" toml:"minio,omitempty" mapstructure:"minio,omitempty"`
}

func (x *ConfigRequest_V1_System) Reset() {
	*x = ConfigRequest_V1_System{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_config_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System) ProtoMessage() {}

func (x *ConfigRequest_V1_System) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_config_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System) Descriptor() ([]byte, []int) {
	return file_config_report_manager_config_request_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ConfigRequest_V1_System) GetMlsa() *shared.Mlsa {
	if x != nil {
		return x.Mlsa
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetService() *ConfigRequest_V1_System_Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetTls() *shared.TLSCredentials {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetLog() *shared.Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetObjstore() *ConfigRequest_V1_System_Objstore {
	if x != nil {
		return x.Objstore
	}
	return nil
}

func (x *ConfigRequest_V1_System) GetMinio() *ConfigRequest_V1_System_Minio {
	if x != nil {
		return x.Minio
	}
	return nil
}

type ConfigRequest_V1_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConfigRequest_V1_Service) Reset() {
	*x = ConfigRequest_V1_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_config_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_Service) ProtoMessage() {}

func (x *ConfigRequest_V1_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_config_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_Service) Descriptor() ([]byte, []int) {
	return file_config_report_manager_config_request_proto_rawDescGZIP(), []int{0, 0, 1}
}

type ConfigRequest_V1_System_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated: Do not use.
	Host                 *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty" toml:"host,omitempty" mapstructure:"host,omitempty"`
	Port                 *wrapperspb.Int32Value  `protobuf:"bytes,2,opt,name=port,proto3" json:"port,omitempty" toml:"port,omitempty" mapstructure:"port,omitempty"`
	EnableLargeReporting *wrapperspb.BoolValue   `protobuf:"bytes,3,opt,name=enable_large_reporting,json=enableLargeReporting,proto3" json:"enable_large_reporting,omitempty" toml:"enable_large_reporting,omitempty" mapstructure:"enable_large_reporting,omitempty"`
}

func (x *ConfigRequest_V1_System_Service) Reset() {
	*x = ConfigRequest_V1_System_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_config_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Service) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Service) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_config_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Service.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Service) Descriptor() ([]byte, []int) {
	return file_config_report_manager_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 0}
}

// Deprecated: Do not use.
func (x *ConfigRequest_V1_System_Service) GetHost() *wrapperspb.StringValue {
	if x != nil {
		return x.Host
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetPort() *wrapperspb.Int32Value {
	if x != nil {
		return x.Port
	}
	return nil
}

func (x *ConfigRequest_V1_System_Service) GetEnableLargeReporting() *wrapperspb.BoolValue {
	if x != nil {
		return x.EnableLargeReporting
	}
	return nil
}

type ConfigRequest_V1_System_Objstore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bucket *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty" toml:"bucket,omitempty" mapstructure:"bucket,omitempty"`
}

func (x *ConfigRequest_V1_System_Objstore) Reset() {
	*x = ConfigRequest_V1_System_Objstore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_config_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Objstore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Objstore) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Objstore) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_config_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Objstore.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Objstore) Descriptor() ([]byte, []int) {
	return file_config_report_manager_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 1}
}

func (x *ConfigRequest_V1_System_Objstore) GetBucket() *wrapperspb.StringValue {
	if x != nil {
		return x.Bucket
	}
	return nil
}

type ConfigRequest_V1_System_Minio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint     *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty" toml:"endpoint,omitempty" mapstructure:"endpoint,omitempty"`
	RootUser     *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=root_user,json=rootUser,proto3" json:"root_user,omitempty" toml:"root_user,omitempty" mapstructure:"root_user,omitempty"`
	RootPassword *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=root_password,json=rootPassword,proto3" json:"root_password,omitempty" toml:"root_password,omitempty" mapstructure:"root_password,omitempty"`
}

func (x *ConfigRequest_V1_System_Minio) Reset() {
	*x = ConfigRequest_V1_System_Minio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_report_manager_config_request_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest_V1_System_Minio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest_V1_System_Minio) ProtoMessage() {}

func (x *ConfigRequest_V1_System_Minio) ProtoReflect() protoreflect.Message {
	mi := &file_config_report_manager_config_request_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest_V1_System_Minio.ProtoReflect.Descriptor instead.
func (*ConfigRequest_V1_System_Minio) Descriptor() ([]byte, []int) {
	return file_config_report_manager_config_request_proto_rawDescGZIP(), []int{0, 0, 0, 2}
}

func (x *ConfigRequest_V1_System_Minio) GetEndpoint() *wrapperspb.StringValue {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *ConfigRequest_V1_System_Minio) GetRootUser() *wrapperspb.StringValue {
	if x != nil {
		return x.RootUser
	}
	return nil
}

func (x *ConfigRequest_V1_System_Minio) GetRootPassword() *wrapperspb.StringValue {
	if x != nil {
		return x.RootPassword
	}
	return nil
}

var File_config_report_manager_config_request_proto protoreflect.FileDescriptor

var file_config_report_manager_config_request_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x1a, 0x1a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x61, 0x32, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x32, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x09, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x02, 0x76, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x52, 0x02, 0x76, 0x31,
	0x1a, 0xe0, 0x08, 0x0a, 0x02, 0x56, 0x31, 0x12, 0x4e, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x52, 0x03, 0x73, 0x79, 0x73, 0x12, 0x4f, 0x0a, 0x03, 0x73, 0x76, 0x63, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x03, 0x73, 0x76, 0x63, 0x1a, 0xad, 0x07, 0x0a, 0x06, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x34, 0x0a, 0x04, 0x6d, 0x6c, 0x73, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4d,
	0x6c, 0x73, 0x61, 0x52, 0x04, 0x6d, 0x6c, 0x73, 0x61, 0x12, 0x5e, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56,
	0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x74, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12, 0x61, 0x0a, 0x08, 0x6f, 0x62,
	0x6a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x62, 0x6a, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x58, 0x0a,
	0x05, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x56, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6f,
	0x52, 0x05, 0x6d, 0x69, 0x6e, 0x69, 0x6f, 0x1a, 0xda, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x16, 0xc2, 0xf3, 0x18, 0x12, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x10, 0xa8, 0x4f, 0x1a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x50, 0x0a, 0x16, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x61, 0x72,
	0x67, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x1a, 0x40, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x34, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x1a, 0xbf, 0x01, 0x0a, 0x05, 0x4d, 0x69, 0x6e, 0x69, 0x6f,
	0x12, 0x38, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x09, 0x72, 0x6f,
	0x6f, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x72, 0x6f, 0x6f,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0d, 0x72, 0x6f, 0x6f, 0x74, 0x5f, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x72, 0x6f, 0x6f, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x1a, 0x09, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3a, 0x1c, 0xc2, 0xf3, 0x18, 0x18, 0x0a, 0x16, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_report_manager_config_request_proto_rawDescOnce sync.Once
	file_config_report_manager_config_request_proto_rawDescData = file_config_report_manager_config_request_proto_rawDesc
)

func file_config_report_manager_config_request_proto_rawDescGZIP() []byte {
	file_config_report_manager_config_request_proto_rawDescOnce.Do(func() {
		file_config_report_manager_config_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_report_manager_config_request_proto_rawDescData)
	})
	return file_config_report_manager_config_request_proto_rawDescData
}

var file_config_report_manager_config_request_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_config_report_manager_config_request_proto_goTypes = []interface{}{
	(*ConfigRequest)(nil),                    // 0: chef.automate.domain.report_manager.ConfigRequest
	(*ConfigRequest_V1)(nil),                 // 1: chef.automate.domain.report_manager.ConfigRequest.V1
	(*ConfigRequest_V1_System)(nil),          // 2: chef.automate.domain.report_manager.ConfigRequest.V1.System
	(*ConfigRequest_V1_Service)(nil),         // 3: chef.automate.domain.report_manager.ConfigRequest.V1.Service
	(*ConfigRequest_V1_System_Service)(nil),  // 4: chef.automate.domain.report_manager.ConfigRequest.V1.System.Service
	(*ConfigRequest_V1_System_Objstore)(nil), // 5: chef.automate.domain.report_manager.ConfigRequest.V1.System.Objstore
	(*ConfigRequest_V1_System_Minio)(nil),    // 6: chef.automate.domain.report_manager.ConfigRequest.V1.System.Minio
	(*shared.Mlsa)(nil),                      // 7: chef.automate.infra.config.Mlsa
	(*shared.TLSCredentials)(nil),            // 8: chef.automate.infra.config.TLSCredentials
	(*shared.Log)(nil),                       // 9: chef.automate.infra.config.Log
	(*wrapperspb.StringValue)(nil),           // 10: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),            // 11: google.protobuf.Int32Value
	(*wrapperspb.BoolValue)(nil),             // 12: google.protobuf.BoolValue
}
var file_config_report_manager_config_request_proto_depIdxs = []int32{
	1,  // 0: chef.automate.domain.report_manager.ConfigRequest.v1:type_name -> chef.automate.domain.report_manager.ConfigRequest.V1
	2,  // 1: chef.automate.domain.report_manager.ConfigRequest.V1.sys:type_name -> chef.automate.domain.report_manager.ConfigRequest.V1.System
	3,  // 2: chef.automate.domain.report_manager.ConfigRequest.V1.svc:type_name -> chef.automate.domain.report_manager.ConfigRequest.V1.Service
	7,  // 3: chef.automate.domain.report_manager.ConfigRequest.V1.System.mlsa:type_name -> chef.automate.infra.config.Mlsa
	4,  // 4: chef.automate.domain.report_manager.ConfigRequest.V1.System.service:type_name -> chef.automate.domain.report_manager.ConfigRequest.V1.System.Service
	8,  // 5: chef.automate.domain.report_manager.ConfigRequest.V1.System.tls:type_name -> chef.automate.infra.config.TLSCredentials
	9,  // 6: chef.automate.domain.report_manager.ConfigRequest.V1.System.log:type_name -> chef.automate.infra.config.Log
	5,  // 7: chef.automate.domain.report_manager.ConfigRequest.V1.System.objstore:type_name -> chef.automate.domain.report_manager.ConfigRequest.V1.System.Objstore
	6,  // 8: chef.automate.domain.report_manager.ConfigRequest.V1.System.minio:type_name -> chef.automate.domain.report_manager.ConfigRequest.V1.System.Minio
	10, // 9: chef.automate.domain.report_manager.ConfigRequest.V1.System.Service.host:type_name -> google.protobuf.StringValue
	11, // 10: chef.automate.domain.report_manager.ConfigRequest.V1.System.Service.port:type_name -> google.protobuf.Int32Value
	12, // 11: chef.automate.domain.report_manager.ConfigRequest.V1.System.Service.enable_large_reporting:type_name -> google.protobuf.BoolValue
	10, // 12: chef.automate.domain.report_manager.ConfigRequest.V1.System.Objstore.bucket:type_name -> google.protobuf.StringValue
	10, // 13: chef.automate.domain.report_manager.ConfigRequest.V1.System.Minio.endpoint:type_name -> google.protobuf.StringValue
	10, // 14: chef.automate.domain.report_manager.ConfigRequest.V1.System.Minio.root_user:type_name -> google.protobuf.StringValue
	10, // 15: chef.automate.domain.report_manager.ConfigRequest.V1.System.Minio.root_password:type_name -> google.protobuf.StringValue
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_config_report_manager_config_request_proto_init() }
func file_config_report_manager_config_request_proto_init() {
	if File_config_report_manager_config_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_report_manager_config_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_config_report_manager_config_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1); i {
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
		file_config_report_manager_config_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System); i {
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
		file_config_report_manager_config_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_Service); i {
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
		file_config_report_manager_config_request_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Service); i {
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
		file_config_report_manager_config_request_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Objstore); i {
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
		file_config_report_manager_config_request_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest_V1_System_Minio); i {
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
			RawDescriptor: file_config_report_manager_config_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_report_manager_config_request_proto_goTypes,
		DependencyIndexes: file_config_report_manager_config_request_proto_depIdxs,
		MessageInfos:      file_config_report_manager_config_request_proto_msgTypes,
	}.Build()
	File_config_report_manager_config_request_proto = out.File
	file_config_report_manager_config_request_proto_rawDesc = nil
	file_config_report_manager_config_request_proto_goTypes = nil
	file_config_report_manager_config_request_proto_depIdxs = nil
}
