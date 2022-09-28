// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: external/sso/sso.proto

package sso

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetSsoConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaContents         string   `protobuf:"bytes,1,opt,name=ca_contents,json=caContents,proto3" json:"ca_contents,omitempty"`
	SsoUrl             string   `protobuf:"bytes,2,opt,name=sso_url,json=ssoUrl,proto3" json:"sso_url,omitempty"`
	EmailAttr          string   `protobuf:"bytes,3,opt,name=email_attr,json=emailAttr,proto3" json:"email_attr,omitempty"`
	UsernameAttr       string   `protobuf:"bytes,4,opt,name=username_attr,json=usernameAttr,proto3" json:"username_attr,omitempty"`
	GroupsAttr         string   `protobuf:"bytes,5,opt,name=groups_attr,json=groupsAttr,proto3" json:"groups_attr,omitempty"`
	AllowedGroups      []string `protobuf:"bytes,6,rep,name=allowed_groups,json=allowedGroups,proto3" json:"allowed_groups,omitempty"`
	EntityIssuer       string   `protobuf:"bytes,7,opt,name=entity_issuer,json=entityIssuer,proto3" json:"entity_issuer,omitempty"`
	NameIdPolicyFormat string   `protobuf:"bytes,8,opt,name=name_id_policy_format,json=nameIdPolicyFormat,proto3" json:"name_id_policy_format,omitempty"`
}

func (x *GetSsoConfigResponse) Reset() {
	*x = GetSsoConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_sso_sso_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSsoConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSsoConfigResponse) ProtoMessage() {}

func (x *GetSsoConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_sso_sso_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSsoConfigResponse.ProtoReflect.Descriptor instead.
func (*GetSsoConfigResponse) Descriptor() ([]byte, []int) {
	return file_external_sso_sso_proto_rawDescGZIP(), []int{0}
}

func (x *GetSsoConfigResponse) GetCaContents() string {
	if x != nil {
		return x.CaContents
	}
	return ""
}

func (x *GetSsoConfigResponse) GetSsoUrl() string {
	if x != nil {
		return x.SsoUrl
	}
	return ""
}

func (x *GetSsoConfigResponse) GetEmailAttr() string {
	if x != nil {
		return x.EmailAttr
	}
	return ""
}

func (x *GetSsoConfigResponse) GetUsernameAttr() string {
	if x != nil {
		return x.UsernameAttr
	}
	return ""
}

func (x *GetSsoConfigResponse) GetGroupsAttr() string {
	if x != nil {
		return x.GroupsAttr
	}
	return ""
}

func (x *GetSsoConfigResponse) GetAllowedGroups() []string {
	if x != nil {
		return x.AllowedGroups
	}
	return nil
}

func (x *GetSsoConfigResponse) GetEntityIssuer() string {
	if x != nil {
		return x.EntityIssuer
	}
	return ""
}

func (x *GetSsoConfigResponse) GetNameIdPolicyFormat() string {
	if x != nil {
		return x.NameIdPolicyFormat
	}
	return ""
}

type AuthenticateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AuthenticateTokenResponse) Reset() {
	*x = AuthenticateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_sso_sso_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateTokenResponse) ProtoMessage() {}

func (x *AuthenticateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_external_sso_sso_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateTokenResponse.ProtoReflect.Descriptor instead.
func (*AuthenticateTokenResponse) Descriptor() ([]byte, []int) {
	return file_external_sso_sso_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticateTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_external_sso_sso_proto protoreflect.FileDescriptor

var file_external_sso_sso_proto_rawDesc = []byte{
	0x0a, 0x16, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x73,
	0x73, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x73, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x02, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x73, 0x6f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x73, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x73, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x41, 0x74, 0x74, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x41, 0x74, 0x74, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x41, 0x74, 0x74, 0x72, 0x12, 0x25, 0x0a, 0x0e,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x15, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0x35, 0x0a, 0x19, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xc5, 0x02, 0x0a, 0x10, 0x53, 0x73, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53,
	0x73, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x73, 0x6f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73,
	0x73, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x8a, 0xb5, 0x18, 0x1c, 0x0a, 0x0a, 0x73,
	0x73, 0x6f, 0x3a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x73, 0x73, 0x6f, 0x3a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x67, 0x65, 0x74, 0x12, 0x9e, 0x01, 0x0a, 0x0c, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x30, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x1d, 0x0a, 0x0a, 0x73,
	0x73, 0x6f, 0x3a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0f, 0x73, 0x73, 0x6f, 0x3a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x70, 0x6f, 0x73, 0x74, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x73, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_sso_sso_proto_rawDescOnce sync.Once
	file_external_sso_sso_proto_rawDescData = file_external_sso_sso_proto_rawDesc
)

func file_external_sso_sso_proto_rawDescGZIP() []byte {
	file_external_sso_sso_proto_rawDescOnce.Do(func() {
		file_external_sso_sso_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_sso_sso_proto_rawDescData)
	})
	return file_external_sso_sso_proto_rawDescData
}

var file_external_sso_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_external_sso_sso_proto_goTypes = []interface{}{
	(*GetSsoConfigResponse)(nil),      // 0: chef.automate.api.sso.GetSsoConfigResponse
	(*AuthenticateTokenResponse)(nil), // 1: chef.automate.api.sso.AuthenticateTokenResponse
	(*emptypb.Empty)(nil),             // 2: google.protobuf.Empty
}
var file_external_sso_sso_proto_depIdxs = []int32{
	2, // 0: chef.automate.api.sso.SsoConfigService.GetSsoConfig:input_type -> google.protobuf.Empty
	2, // 1: chef.automate.api.sso.SsoConfigService.Authenticate:input_type -> google.protobuf.Empty
	0, // 2: chef.automate.api.sso.SsoConfigService.GetSsoConfig:output_type -> chef.automate.api.sso.GetSsoConfigResponse
	1, // 3: chef.automate.api.sso.SsoConfigService.Authenticate:output_type -> chef.automate.api.sso.AuthenticateTokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_sso_sso_proto_init() }
func file_external_sso_sso_proto_init() {
	if File_external_sso_sso_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_sso_sso_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSsoConfigResponse); i {
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
		file_external_sso_sso_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateTokenResponse); i {
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
			RawDescriptor: file_external_sso_sso_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_external_sso_sso_proto_goTypes,
		DependencyIndexes: file_external_sso_sso_proto_depIdxs,
		MessageInfos:      file_external_sso_sso_proto_msgTypes,
	}.Build()
	File_external_sso_sso_proto = out.File
	file_external_sso_sso_proto_rawDesc = nil
	file_external_sso_sso_proto_goTypes = nil
	file_external_sso_sso_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SsoConfigServiceClient is the client API for SsoConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SsoConfigServiceClient interface {
	GetSsoConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSsoConfigResponse, error)
	Authenticate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AuthenticateTokenResponse, error)
}

type ssoConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSsoConfigServiceClient(cc grpc.ClientConnInterface) SsoConfigServiceClient {
	return &ssoConfigServiceClient{cc}
}

func (c *ssoConfigServiceClient) GetSsoConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSsoConfigResponse, error) {
	out := new(GetSsoConfigResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.sso.SsoConfigService/GetSsoConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ssoConfigServiceClient) Authenticate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AuthenticateTokenResponse, error) {
	out := new(AuthenticateTokenResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.sso.SsoConfigService/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SsoConfigServiceServer is the server API for SsoConfigService service.
type SsoConfigServiceServer interface {
	GetSsoConfig(context.Context, *emptypb.Empty) (*GetSsoConfigResponse, error)
	Authenticate(context.Context, *emptypb.Empty) (*AuthenticateTokenResponse, error)
}

// UnimplementedSsoConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSsoConfigServiceServer struct {
}

func (*UnimplementedSsoConfigServiceServer) GetSsoConfig(context.Context, *emptypb.Empty) (*GetSsoConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSsoConfig not implemented")
}
func (*UnimplementedSsoConfigServiceServer) Authenticate(context.Context, *emptypb.Empty) (*AuthenticateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}

func RegisterSsoConfigServiceServer(s *grpc.Server, srv SsoConfigServiceServer) {
	s.RegisterService(&_SsoConfigService_serviceDesc, srv)
}

func _SsoConfigService_GetSsoConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoConfigServiceServer).GetSsoConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.sso.SsoConfigService/GetSsoConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoConfigServiceServer).GetSsoConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SsoConfigService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SsoConfigServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.sso.SsoConfigService/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SsoConfigServiceServer).Authenticate(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _SsoConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.sso.SsoConfigService",
	HandlerType: (*SsoConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSsoConfig",
			Handler:    _SsoConfigService_GetSsoConfig_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _SsoConfigService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "external/sso/sso.proto",
}
