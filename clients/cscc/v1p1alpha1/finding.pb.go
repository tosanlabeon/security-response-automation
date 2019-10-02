// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1p1alpha1/finding.proto

package securitycenter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The state of the finding.
type Finding_State int32

const (
	// Unspecified state.
	Finding_STATE_UNSPECIFIED Finding_State = 0
	// The finding requires attention and has not been addressed yet.
	Finding_ACTIVE Finding_State = 1
	// The finding has been fixed, triaged as a non-issue or otherwise addressed
	// and is no longer active.
	Finding_INACTIVE Finding_State = 2
)

var Finding_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	2: "INACTIVE",
}

var Finding_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"INACTIVE":          2,
}

func (x Finding_State) String() string {
	return proto.EnumName(Finding_State_name, int32(x))
}

func (Finding_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_986b20153e409057, []int{0, 0}
}

// Cloud Security Command Center (Cloud SCC) finding.
//
// A finding is a record of assessment data (security, risk, health or privacy)
// ingested into Cloud SCC for presentation, notification, analysis,
// policy testing, and enforcement. For example, an XSS vulnerability in an
// App Engine application is a finding.
type Finding struct {
	// The relative resource name of this finding. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/sources/456/findings/789"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The relative resource name of the source the finding belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// This field is immutable after creation time.
	// For example:
	// "organizations/123/sources/456"
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// The full resource name of the Google Cloud Platform (GCP) resource this
	// finding is for. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	// This field is immutable after creation time.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The state of the finding.
	State Finding_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.securitycenter.v1p1alpha1.Finding_State" json:"state,omitempty"`
	// The additional taxonomy group within findings from a given source.
	// This field is immutable after creation time.
	// Example: "XSS_FLASH_INJECTION"
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	// The URI that, if available, points to a web page outside of Cloud SCC
	// where additional information about the finding can be found. This field is
	// guaranteed to be either empty or a well formed URL.
	ExternalUri string `protobuf:"bytes,6,opt,name=external_uri,json=externalUri,proto3" json:"external_uri,omitempty"`
	// Source specific properties. These properties are managed by the source
	// that writes the finding. The key names in the source_properties map must be
	// between 1 and 255 characters, and must start with a letter and contain
	// alphanumeric characters or underscores only.
	SourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=source_properties,json=sourceProperties,proto3" json:"source_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. User specified security marks. These marks are entirely
	// managed by the user and come from the SecurityMarks resource that belongs
	// to the finding.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the event took place. For example, if the finding
	// represents an open firewall it would capture the time the detector believes
	// the firewall became open. The accuracy is determined by the detector.
	EventTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// The time at which the finding was created in Cloud SCC.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Finding) Reset()         { *m = Finding{} }
func (m *Finding) String() string { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()    {}
func (*Finding) Descriptor() ([]byte, []int) {
	return fileDescriptor_986b20153e409057, []int{0}
}

func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (m *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(m, src)
}
func (m *Finding) XXX_Size() int {
	return xxx_messageInfo_Finding.Size(m)
}
func (m *Finding) XXX_DiscardUnknown() {
	xxx_messageInfo_Finding.DiscardUnknown(m)
}

var xxx_messageInfo_Finding proto.InternalMessageInfo

func (m *Finding) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Finding) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Finding) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Finding) GetState() Finding_State {
	if m != nil {
		return m.State
	}
	return Finding_STATE_UNSPECIFIED
}

func (m *Finding) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Finding) GetExternalUri() string {
	if m != nil {
		return m.ExternalUri
	}
	return ""
}

func (m *Finding) GetSourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.SourceProperties
	}
	return nil
}

func (m *Finding) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Finding) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *Finding) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.v1p1alpha1.Finding_State", Finding_State_name, Finding_State_value)
	proto.RegisterType((*Finding)(nil), "google.cloud.securitycenter.v1p1alpha1.Finding")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1p1alpha1.Finding.SourcePropertiesEntry")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1p1alpha1/finding.proto", fileDescriptor_986b20153e409057)
}

var fileDescriptor_986b20153e409057 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xef, 0x6a, 0xd3, 0x50,
	0x14, 0x37, 0xd9, 0xda, 0xb5, 0xa7, 0xdd, 0xe8, 0x2e, 0x6c, 0x84, 0x22, 0x58, 0x27, 0x8c, 0x2a,
	0x92, 0xd0, 0xaa, 0x30, 0xd3, 0x4f, 0xb5, 0x66, 0x52, 0xc4, 0x52, 0xd2, 0xae, 0x1f, 0xb4, 0x50,
	0xee, 0xb2, 0xbb, 0x18, 0x96, 0xde, 0x1b, 0x6e, 0x6e, 0x8a, 0xfd, 0xea, 0xe3, 0xf8, 0x22, 0x82,
	0x8f, 0xe2, 0x53, 0x48, 0xee, 0x4d, 0xaa, 0xa9, 0xc2, 0xea, 0xb7, 0x7b, 0xce, 0xef, 0xcf, 0x39,
	0x39, 0xe7, 0x04, 0x5e, 0xfa, 0x8c, 0xf9, 0x21, 0xb1, 0xbc, 0x90, 0x25, 0x37, 0x56, 0x4c, 0xbc,
	0x84, 0x07, 0x62, 0xed, 0x11, 0x2a, 0x08, 0xb7, 0x56, 0x9d, 0xa8, 0x83, 0xc3, 0xe8, 0x33, 0xee,
	0x58, 0xb7, 0x01, 0xbd, 0x09, 0xa8, 0x6f, 0x46, 0x9c, 0x09, 0x86, 0xce, 0x95, 0xca, 0x94, 0x2a,
	0xb3, 0xa8, 0x32, 0x7f, 0xab, 0x9a, 0xbd, 0x1d, 0xdd, 0x73, 0x64, 0xb1, 0xc4, 0xfc, 0x2e, 0x56,
	0x45, 0x9a, 0x0f, 0x33, 0xb1, 0x8c, 0xae, 0x93, 0x5b, 0x2b, 0x16, 0x3c, 0xf1, 0x44, 0x86, 0x3e,
	0xda, 0x46, 0x45, 0xb0, 0x24, 0xb1, 0xc0, 0xcb, 0x68, 0x4b, 0x8e, 0xa3, 0xc0, 0xc2, 0x94, 0x32,
	0x81, 0x45, 0xc0, 0x68, 0x66, 0x7e, 0xf6, 0xbd, 0x04, 0x07, 0x97, 0xea, 0x9b, 0x10, 0x82, 0x7d,
	0x8a, 0x97, 0xc4, 0xd0, 0x5a, 0x5a, 0xbb, 0xea, 0xca, 0x37, 0x3a, 0x85, 0x72, 0x84, 0x39, 0xa1,
	0xc2, 0xd0, 0x65, 0x36, 0x8b, 0xd0, 0x13, 0x38, 0xe4, 0x24, 0x66, 0x09, 0xf7, 0xc8, 0x42, 0x8a,
	0xf6, 0x24, 0x5c, 0xcf, 0x93, 0xa3, 0x54, 0xfc, 0x1e, 0x4a, 0xb1, 0xc0, 0x82, 0x18, 0xfb, 0x2d,
	0xad, 0x7d, 0xd4, 0x7d, 0x65, 0xee, 0x36, 0x2e, 0x33, 0x6b, 0xc8, 0x9c, 0xa4, 0x62, 0x57, 0x79,
	0xa0, 0x26, 0x54, 0x3c, 0x2c, 0x88, 0xcf, 0xf8, 0xda, 0x28, 0xc9, 0x62, 0x9b, 0x18, 0x3d, 0x86,
	0x3a, 0xf9, 0x22, 0x08, 0xa7, 0x38, 0x5c, 0x24, 0x3c, 0x30, 0xca, 0x12, 0xaf, 0xe5, 0xb9, 0x2b,
	0x1e, 0x20, 0x0e, 0xc7, 0x59, 0xbb, 0x11, 0x67, 0x11, 0xe1, 0x22, 0x20, 0xb1, 0x71, 0xd0, 0xda,
	0x6b, 0xd7, 0xba, 0xce, 0x7f, 0xf7, 0x25, 0x8d, 0xc6, 0x1b, 0x1f, 0x87, 0x0a, 0xbe, 0x76, 0x1b,
	0xf1, 0x56, 0x1a, 0xcd, 0xe1, 0xa8, 0xb8, 0x51, 0xa3, 0xd2, 0xd2, 0xda, 0xb5, 0xdd, 0x07, 0x31,
	0xc9, 0x90, 0x0f, 0xa9, 0xd8, 0x3d, 0x8c, 0xff, 0x0c, 0xd1, 0x6b, 0x00, 0xb2, 0x22, 0x54, 0x2c,
	0xd2, 0x8d, 0x1b, 0x55, 0xe9, 0xdc, 0xcc, 0x9d, 0xf3, 0x73, 0x30, 0xa7, 0xf9, 0x39, 0xb8, 0x55,
	0xc9, 0x4e, 0x63, 0xd4, 0x83, 0x9a, 0xc7, 0x09, 0x16, 0x44, 0x69, 0xe1, 0x5e, 0x2d, 0x28, 0x7a,
	0x9a, 0x68, 0x7e, 0x82, 0x93, 0x7f, 0x0e, 0x00, 0x35, 0x60, 0xef, 0x8e, 0xac, 0xb3, 0xf3, 0x49,
	0x9f, 0xe8, 0x39, 0x94, 0x56, 0x38, 0x4c, 0x88, 0x3c, 0x9e, 0x5a, 0xf7, 0xf4, 0xaf, 0x0a, 0xb3,
	0x14, 0x75, 0x15, 0xc9, 0xd6, 0x2f, 0xb4, 0xb3, 0x0b, 0x28, 0xc9, 0xad, 0xa3, 0x13, 0x38, 0x9e,
	0x4c, 0xfb, 0x53, 0x67, 0x71, 0x35, 0x9a, 0x8c, 0x9d, 0xc1, 0xf0, 0x72, 0xe8, 0xbc, 0x6d, 0x3c,
	0x40, 0x00, 0xe5, 0xfe, 0x60, 0x3a, 0x9c, 0x39, 0x0d, 0x0d, 0xd5, 0xa1, 0x32, 0x1c, 0x65, 0x91,
	0xfe, 0xe6, 0xab, 0x0e, 0xcf, 0x3c, 0xb6, 0xdc, 0x71, 0xb4, 0x63, 0xed, 0xe3, 0x34, 0x63, 0xfa,
	0x2c, 0xc4, 0xd4, 0x37, 0x19, 0xf7, 0x2d, 0x9f, 0x50, 0xd9, 0x9a, 0xa5, 0x20, 0x1c, 0x05, 0xf1,
	0x7d, 0xff, 0x6c, 0xaf, 0x88, 0x7c, 0xd3, 0xcf, 0xdf, 0x29, 0xdb, 0x81, 0x6c, 0x20, 0xdf, 0xe0,
	0x40, 0x35, 0x30, 0xeb, 0x8c, 0x3b, 0x7d, 0xa9, 0xfb, 0x91, 0x13, 0xe7, 0x92, 0x38, 0x2f, 0x12,
	0xe7, 0xb3, 0x4d, 0x81, 0x9f, 0xfa, 0x53, 0x45, 0xb4, 0x6d, 0xc9, 0xb4, 0xed, 0x22, 0xd5, 0xb6,
	0x53, 0xae, 0x32, 0xbd, 0x2e, 0xcb, 0xf6, 0x5f, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x96,
	0x86, 0x62, 0xcf, 0x04, 0x00, 0x00,
}