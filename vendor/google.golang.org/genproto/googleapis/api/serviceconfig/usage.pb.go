// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/usage.proto
// DO NOT EDIT!

package google_api // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Service access types.
//
// Access to restricted API features is always controlled by
// [visibility][google.api.Visibility], independent of the ServiceAccess type.
//
type Usage_ServiceAccess int32

const (
	// The service can only be seen/used by users identified in the service's
	// access control policy.
	//
	// If the service has not been whitelisted by your domain administrator
	// for out-of-org publishing, then this mode will be treated like
	// ORG_RESTRICTED.
	Usage_RESTRICTED Usage_ServiceAccess = 0
	// The service can be seen/used by anyone.
	//
	// If the service has not been whitelisted by your domain administrator
	// for out-of-org publishing, then this mode will be treated like
	// ORG_PUBLIC.
	//
	// The discovery document for the service will also be public and allow
	// unregistered access.
	Usage_PUBLIC Usage_ServiceAccess = 1
	// The service can be seen/used by users identified in the service's
	// access control policy and they are within the organization that owns the
	// service.
	//
	// Access is further constrained to the group
	// controlled by the administrator of the project/org that owns the
	// service.
	Usage_ORG_RESTRICTED Usage_ServiceAccess = 2
	// The service can be seen/used by the group of users controlled by the
	// administrator of the project/org that owns the service.
	Usage_ORG_PUBLIC Usage_ServiceAccess = 3
)

var Usage_ServiceAccess_name = map[int32]string{
	0: "RESTRICTED",
	1: "PUBLIC",
	2: "ORG_RESTRICTED",
	3: "ORG_PUBLIC",
}
var Usage_ServiceAccess_value = map[string]int32{
	"RESTRICTED":     0,
	"PUBLIC":         1,
	"ORG_RESTRICTED": 2,
	"ORG_PUBLIC":     3,
}

func (x Usage_ServiceAccess) String() string {
	return proto.EnumName(Usage_ServiceAccess_name, int32(x))
}
func (Usage_ServiceAccess) EnumDescriptor() ([]byte, []int) { return fileDescriptor14, []int{0, 0} }

// Configuration controlling usage of a service.
type Usage struct {
	// Controls which users can see or activate the service.
	ServiceAccess Usage_ServiceAccess `protobuf:"varint,4,opt,name=service_access,json=serviceAccess,enum=google.api.Usage_ServiceAccess" json:"service_access,omitempty"`
	// Requirements that must be satisfied before a consumer project can use the
	// service. Each requirement is of the form <service.name>/<requirement-id>;
	// for example 'serviceusage.googleapis.com/billing-enabled'.
	Requirements []string `protobuf:"bytes,1,rep,name=requirements" json:"requirements,omitempty"`
	// Services that must be activated in order for this service to be used.
	// The set of services activated as a result of these relations are all
	// activated in parallel with no guaranteed order of activation.
	// Each string is a service name, e.g. `calendar.googleapis.com`.
	DependsOnServices []string `protobuf:"bytes,2,rep,name=depends_on_services,json=dependsOnServices" json:"depends_on_services,omitempty"`
	// Services that must be contacted before a consumer can begin using the
	// service. Each service will be contacted in sequence, and, if any activation
	// call fails, the entire activation will fail. Each hook is of the form
	// <service.name>/<hook-id>, where <hook-id> is optional; for example:
	// 'robotservice.googleapis.com/default'.
	ActivationHooks []string `protobuf:"bytes,3,rep,name=activation_hooks,json=activationHooks" json:"activation_hooks,omitempty"`
	// Services that must be contacted before a consumer can deactivate a
	// service. Each service will be contacted in sequence, and, if any
	// deactivation call fails, the entire deactivation will fail. Each hook is
	// of the form <service.name>/<hook-id>, where <hook-id> is optional; for
	// example:
	// 'compute.googleapis.com/'.
	DeactivationHooks []string `protobuf:"bytes,5,rep,name=deactivation_hooks,json=deactivationHooks" json:"deactivation_hooks,omitempty"`
	// Individual rules for configuring usage on selected methods.
	Rules []*UsageRule `protobuf:"bytes,6,rep,name=rules" json:"rules,omitempty"`
}

func (m *Usage) Reset()                    { *m = Usage{} }
func (m *Usage) String() string            { return proto.CompactTextString(m) }
func (*Usage) ProtoMessage()               {}
func (*Usage) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *Usage) GetRules() []*UsageRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

// Usage configuration rules for the service.
//
// NOTE: Under development.
//
//
// Use this rule to configure unregistered calls for the service. Unregistered
// calls are calls that do not contain consumer project identity.
// (Example: calls that do not contain an API key).
// By default, API methods do not allow unregistered calls, and each method call
// must be identified by a consumer project identity. Use this rule to
// allow/disallow unregistered calls.
//
// Example of an API that wants to allow unregistered calls for entire service.
//
//     usage:
//       rules:
//       - selector: "*"
//         allow_unregistered_calls: true
//
// Example of a method that wants to allow unregistered calls.
//
//     usage:
//       rules:
//       - selector: "google.example.library.v1.LibraryService.CreateBook"
//         allow_unregistered_calls: true
type UsageRule struct {
	// Selects the methods to which this rule applies. Use '*' to indicate all
	// methods in all APIs.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax details.
	Selector string `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	// True, if the method allows unregistered calls; false otherwise.
	AllowUnregisteredCalls bool `protobuf:"varint,2,opt,name=allow_unregistered_calls,json=allowUnregisteredCalls" json:"allow_unregistered_calls,omitempty"`
}

func (m *UsageRule) Reset()                    { *m = UsageRule{} }
func (m *UsageRule) String() string            { return proto.CompactTextString(m) }
func (*UsageRule) ProtoMessage()               {}
func (*UsageRule) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func init() {
	proto.RegisterType((*Usage)(nil), "google.api.Usage")
	proto.RegisterType((*UsageRule)(nil), "google.api.UsageRule")
	proto.RegisterEnum("google.api.Usage_ServiceAccess", Usage_ServiceAccess_name, Usage_ServiceAccess_value)
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/usage.proto", fileDescriptor14)
}

var fileDescriptor14 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xc1, 0xcf, 0xd2, 0x30,
	0x18, 0xc6, 0x1d, 0x08, 0x81, 0x57, 0x99, 0x58, 0xa3, 0x59, 0x3c, 0x28, 0xd9, 0x09, 0x63, 0xdc,
	0x12, 0xbc, 0x78, 0x95, 0x89, 0x4a, 0x62, 0x02, 0x29, 0x70, 0x5e, 0x6a, 0xf7, 0x5a, 0x17, 0x47,
	0x8b, 0xed, 0x86, 0x7f, 0xbc, 0x97, 0xaf, 0x2b, 0xfb, 0x60, 0x7c, 0x5c, 0x96, 0xf4, 0xf9, 0xfd,
	0xda, 0x67, 0x6f, 0x53, 0x48, 0x84, 0x52, 0xa2, 0xc0, 0x48, 0xa8, 0x82, 0x49, 0x11, 0x29, 0x2d,
	0x62, 0x81, 0xf2, 0xa0, 0x55, 0xa9, 0xe2, 0x13, 0x62, 0x87, 0xdc, 0xc4, 0xf6, 0x13, 0x1b, 0xd4,
	0xc7, 0x9c, 0x23, 0x57, 0xf2, 0x57, 0x2e, 0xe2, 0xca, 0x30, 0x81, 0x91, 0x13, 0x09, 0x34, 0x87,
	0x58, 0x2b, 0xfc, 0xdf, 0x81, 0xde, 0xae, 0x66, 0xe4, 0x2b, 0xf8, 0xcd, 0x96, 0x94, 0x71, 0x8e,
	0xc6, 0x04, 0x8f, 0x27, 0xde, 0xd4, 0x9f, 0xbd, 0x8d, 0x2e, 0x7a, 0xe4, 0xd4, 0x68, 0x73, 0xf2,
	0x3e, 0x3b, 0x8d, 0x8e, 0x4c, 0x7b, 0x49, 0x42, 0x78, 0xaa, 0xf1, 0x6f, 0x95, 0x6b, 0xdc, 0xa3,
	0x2c, 0x4d, 0xe0, 0x4d, 0xba, 0xd3, 0x21, 0xbd, 0xca, 0x48, 0x04, 0x2f, 0x32, 0x3c, 0xa0, 0xcc,
	0x4c, 0xaa, 0x64, 0xda, 0xec, 0x37, 0x41, 0xc7, 0xa9, 0xcf, 0x1b, 0xb4, 0x92, 0x4d, 0x8f, 0x21,
	0xef, 0x60, 0xcc, 0x78, 0x99, 0x1f, 0x59, 0x99, 0x5b, 0xff, 0xb7, 0x52, 0x7f, 0x4c, 0xd0, 0x75,
	0xf2, 0xb3, 0x4b, 0xfe, 0xbd, 0x8e, 0xc9, 0x07, 0x20, 0x19, 0xde, 0xc8, 0xbd, 0xfb, 0x93, 0x1f,
	0xea, 0xef, 0xa1, 0xa7, 0xab, 0xc2, 0x76, 0xf7, 0xad, 0xf1, 0x64, 0xf6, 0xf2, 0x66, 0x58, 0x6a,
	0x29, 0x3d, 0x39, 0xe1, 0x0a, 0x46, 0x57, 0xa3, 0x13, 0x1f, 0x80, 0x2e, 0x36, 0x5b, 0xba, 0x4c,
	0xb6, 0x8b, 0x2f, 0xe3, 0x47, 0x04, 0xa0, 0xbf, 0xde, 0xcd, 0x7f, 0x2c, 0x93, 0xb1, 0x47, 0x08,
	0xf8, 0x2b, 0xfa, 0x2d, 0x6d, 0xf1, 0x4e, 0xed, 0xd7, 0x59, 0xe3, 0x74, 0x43, 0x06, 0xc3, 0x73,
	0x09, 0x79, 0x0d, 0x03, 0x83, 0x05, 0xf2, 0x52, 0x69, 0x7b, 0x69, 0x9e, 0xfd, 0xdf, 0xf3, 0x9a,
	0x7c, 0x82, 0x80, 0x15, 0x85, 0xfa, 0x97, 0x56, 0x52, 0xa3, 0xc8, 0x4d, 0x89, 0x1a, 0xb3, 0x94,
	0xdb, 0xac, 0xbe, 0x35, 0x6f, 0x3a, 0xa0, 0xaf, 0x1c, 0xdf, 0xb5, 0x70, 0x52, 0xd3, 0xf9, 0x1b,
	0xf0, 0xb9, 0xda, 0xb7, 0xc6, 0x9a, 0x83, 0xab, 0x5c, 0xd7, 0x4f, 0x61, 0xed, 0xfd, 0xec, 0xbb,
	0x37, 0xf1, 0xf1, 0x2e, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x19, 0x60, 0x91, 0x5a, 0x02, 0x00, 0x00,
}
