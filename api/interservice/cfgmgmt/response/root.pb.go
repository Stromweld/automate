// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: interservice/cfgmgmt/response/root.proto

package response

import (
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

type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty" toml:"version,omitempty" mapstructure:"version,omitempty"`
	Sha     string `protobuf:"bytes,3,opt,name=sha,proto3" json:"sha,omitempty" toml:"sha,omitempty" mapstructure:"sha,omitempty"`
	Built   string `protobuf:"bytes,4,opt,name=built,proto3" json:"built,omitempty" toml:"built,omitempty" mapstructure:"built,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_response_root_proto_rawDescGZIP(), []int{0}
}

func (x *VersionInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *VersionInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionInfo) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *VersionInfo) GetBuilt() string {
	if x != nil {
		return x.Built
	}
	return ""
}

// Health message
//
// The config-mgmt-service health is constructed with:
//   - Status:
//     => ok:             Everything is alright
//     => initialization: The service is in its initialization process
//     => warning:        Something might be wrong?
//     => critical:       Something is wrong!
//
// @afiune: Here we can add more health information to the response
type Health struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty" toml:"status,omitempty" mapstructure:"status,omitempty"`
}

func (x *Health) Reset() {
	*x = Health{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Health) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Health) ProtoMessage() {}

func (x *Health) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Health.ProtoReflect.Descriptor instead.
func (*Health) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_response_root_proto_rawDescGZIP(), []int{1}
}

func (x *Health) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type Organizations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organizations []string `protobuf:"bytes,1,rep,name=organizations,proto3" json:"organizations,omitempty" toml:"organizations,omitempty" mapstructure:"organizations,omitempty"`
}

func (x *Organizations) Reset() {
	*x = Organizations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organizations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organizations) ProtoMessage() {}

func (x *Organizations) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organizations.ProtoReflect.Descriptor instead.
func (*Organizations) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_response_root_proto_rawDescGZIP(), []int{2}
}

func (x *Organizations) GetOrganizations() []string {
	if x != nil {
		return x.Organizations
	}
	return nil
}

type SourceFQDNS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceFqdns []string `protobuf:"bytes,1,rep,name=source_fqdns,json=sourceFqdns,proto3" json:"source_fqdns,omitempty" toml:"source_fqdns,omitempty" mapstructure:"source_fqdns,omitempty"`
}

func (x *SourceFQDNS) Reset() {
	*x = SourceFQDNS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceFQDNS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceFQDNS) ProtoMessage() {}

func (x *SourceFQDNS) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceFQDNS.ProtoReflect.Descriptor instead.
func (*SourceFQDNS) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_response_root_proto_rawDescGZIP(), []int{3}
}

func (x *SourceFQDNS) GetSourceFqdns() []string {
	if x != nil {
		return x.SourceFqdns
	}
	return nil
}

type ExportData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty" toml:"content,omitempty" mapstructure:"content,omitempty"`
}

func (x *ExportData) Reset() {
	*x = ExportData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExportData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportData) ProtoMessage() {}

func (x *ExportData) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportData.ProtoReflect.Descriptor instead.
func (*ExportData) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_response_root_proto_rawDescGZIP(), []int{4}
}

func (x *ExportData) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type ReportExportData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty" toml:"content,omitempty" mapstructure:"content,omitempty"`
}

func (x *ReportExportData) Reset() {
	*x = ReportExportData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportExportData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportExportData) ProtoMessage() {}

func (x *ReportExportData) ProtoReflect() protoreflect.Message {
	mi := &file_interservice_cfgmgmt_response_root_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportExportData.ProtoReflect.Descriptor instead.
func (*ReportExportData) Descriptor() ([]byte, []int) {
	return file_interservice_cfgmgmt_response_root_proto_rawDescGZIP(), []int{5}
}

func (x *ReportExportData) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_interservice_cfgmgmt_response_root_proto protoreflect.FileDescriptor

var file_interservice_cfgmgmt_response_root_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f,
	0x72, 0x6f, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x63, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x68, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x22, 0x20, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x0d, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x30, 0x0a, 0x0b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x51, 0x44, 0x4e, 0x53, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x66, 0x71, 0x64, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x71, 0x64, 0x6e,
	0x73, 0x22, 0x26, 0x0a, 0x0a, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x66, 0x67, 0x6d, 0x67, 0x6d, 0x74, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interservice_cfgmgmt_response_root_proto_rawDescOnce sync.Once
	file_interservice_cfgmgmt_response_root_proto_rawDescData = file_interservice_cfgmgmt_response_root_proto_rawDesc
)

func file_interservice_cfgmgmt_response_root_proto_rawDescGZIP() []byte {
	file_interservice_cfgmgmt_response_root_proto_rawDescOnce.Do(func() {
		file_interservice_cfgmgmt_response_root_proto_rawDescData = protoimpl.X.CompressGZIP(file_interservice_cfgmgmt_response_root_proto_rawDescData)
	})
	return file_interservice_cfgmgmt_response_root_proto_rawDescData
}

var file_interservice_cfgmgmt_response_root_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_interservice_cfgmgmt_response_root_proto_goTypes = []interface{}{
	(*VersionInfo)(nil),      // 0: chef.automate.domain.cfgmgmt.response.VersionInfo
	(*Health)(nil),           // 1: chef.automate.domain.cfgmgmt.response.Health
	(*Organizations)(nil),    // 2: chef.automate.domain.cfgmgmt.response.Organizations
	(*SourceFQDNS)(nil),      // 3: chef.automate.domain.cfgmgmt.response.SourceFQDNS
	(*ExportData)(nil),       // 4: chef.automate.domain.cfgmgmt.response.ExportData
	(*ReportExportData)(nil), // 5: chef.automate.domain.cfgmgmt.response.ReportExportData
}
var file_interservice_cfgmgmt_response_root_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_interservice_cfgmgmt_response_root_proto_init() }
func file_interservice_cfgmgmt_response_root_proto_init() {
	if File_interservice_cfgmgmt_response_root_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interservice_cfgmgmt_response_root_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionInfo); i {
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
		file_interservice_cfgmgmt_response_root_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Health); i {
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
		file_interservice_cfgmgmt_response_root_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organizations); i {
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
		file_interservice_cfgmgmt_response_root_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceFQDNS); i {
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
		file_interservice_cfgmgmt_response_root_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExportData); i {
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
		file_interservice_cfgmgmt_response_root_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportExportData); i {
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
			RawDescriptor: file_interservice_cfgmgmt_response_root_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interservice_cfgmgmt_response_root_proto_goTypes,
		DependencyIndexes: file_interservice_cfgmgmt_response_root_proto_depIdxs,
		MessageInfos:      file_interservice_cfgmgmt_response_root_proto_msgTypes,
	}.Build()
	File_interservice_cfgmgmt_response_root_proto = out.File
	file_interservice_cfgmgmt_response_root_proto_rawDesc = nil
	file_interservice_cfgmgmt_response_root_proto_goTypes = nil
	file_interservice_cfgmgmt_response_root_proto_depIdxs = nil
}
