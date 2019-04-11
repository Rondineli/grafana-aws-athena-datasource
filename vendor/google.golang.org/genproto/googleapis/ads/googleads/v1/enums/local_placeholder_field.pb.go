// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/local_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for Local placeholder fields.
type LocalPlaceholderFieldEnum_LocalPlaceholderField int32

const (
	// Not specified.
	LocalPlaceholderFieldEnum_UNSPECIFIED LocalPlaceholderFieldEnum_LocalPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	LocalPlaceholderFieldEnum_UNKNOWN LocalPlaceholderFieldEnum_LocalPlaceholderField = 1
	// Data Type: STRING. Required. Unique ID.
	LocalPlaceholderFieldEnum_DEAL_ID LocalPlaceholderFieldEnum_LocalPlaceholderField = 2
	// Data Type: STRING. Required. Main headline with local deal title to be
	// shown in dynamic ad.
	LocalPlaceholderFieldEnum_DEAL_NAME LocalPlaceholderFieldEnum_LocalPlaceholderField = 3
	// Data Type: STRING. Local deal subtitle to be shown in dynamic ad.
	LocalPlaceholderFieldEnum_SUBTITLE LocalPlaceholderFieldEnum_LocalPlaceholderField = 4
	// Data Type: STRING. Description of local deal to be shown in dynamic ad.
	LocalPlaceholderFieldEnum_DESCRIPTION LocalPlaceholderFieldEnum_LocalPlaceholderField = 5
	// Data Type: STRING. Price to be shown in the ad. Highly recommended for
	// dynamic ads. Example: "100.00 USD"
	LocalPlaceholderFieldEnum_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 6
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	LocalPlaceholderFieldEnum_FORMATTED_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 7
	// Data Type: STRING. Sale price to be shown in the ad.
	// Example: "80.00 USD"
	LocalPlaceholderFieldEnum_SALE_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 8
	// Data Type: STRING. Formatted sale price to be shown in the ad.
	// Example: "On sale for $80.00", "$60 - $80"
	LocalPlaceholderFieldEnum_FORMATTED_SALE_PRICE LocalPlaceholderFieldEnum_LocalPlaceholderField = 9
	// Data Type: URL. Image to be displayed in the ad.
	LocalPlaceholderFieldEnum_IMAGE_URL LocalPlaceholderFieldEnum_LocalPlaceholderField = 10
	// Data Type: STRING. Complete property address, including postal code.
	LocalPlaceholderFieldEnum_ADDRESS LocalPlaceholderFieldEnum_LocalPlaceholderField = 11
	// Data Type: STRING. Category of local deal used to group like items
	// together for recommendation engine.
	LocalPlaceholderFieldEnum_CATEGORY LocalPlaceholderFieldEnum_LocalPlaceholderField = 12
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	LocalPlaceholderFieldEnum_CONTEXTUAL_KEYWORDS LocalPlaceholderFieldEnum_LocalPlaceholderField = 13
	// Data Type: URL_LIST. Required. Final URLs to be used in ad when using
	// Upgraded URLs; the more specific the better (e.g. the individual URL of a
	// specific local deal and its location).
	LocalPlaceholderFieldEnum_FINAL_URLS LocalPlaceholderFieldEnum_LocalPlaceholderField = 14
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	LocalPlaceholderFieldEnum_FINAL_MOBILE_URLS LocalPlaceholderFieldEnum_LocalPlaceholderField = 15
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	LocalPlaceholderFieldEnum_TRACKING_URL LocalPlaceholderFieldEnum_LocalPlaceholderField = 16
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	LocalPlaceholderFieldEnum_ANDROID_APP_LINK LocalPlaceholderFieldEnum_LocalPlaceholderField = 17
	// Data Type: STRING_LIST. List of recommended local deal IDs to show
	// together with this item.
	LocalPlaceholderFieldEnum_SIMILAR_DEAL_IDS LocalPlaceholderFieldEnum_LocalPlaceholderField = 18
	// Data Type: STRING. iOS app link.
	LocalPlaceholderFieldEnum_IOS_APP_LINK LocalPlaceholderFieldEnum_LocalPlaceholderField = 19
	// Data Type: INT64. iOS app store ID.
	LocalPlaceholderFieldEnum_IOS_APP_STORE_ID LocalPlaceholderFieldEnum_LocalPlaceholderField = 20
)

var LocalPlaceholderFieldEnum_LocalPlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "DEAL_ID",
	3:  "DEAL_NAME",
	4:  "SUBTITLE",
	5:  "DESCRIPTION",
	6:  "PRICE",
	7:  "FORMATTED_PRICE",
	8:  "SALE_PRICE",
	9:  "FORMATTED_SALE_PRICE",
	10: "IMAGE_URL",
	11: "ADDRESS",
	12: "CATEGORY",
	13: "CONTEXTUAL_KEYWORDS",
	14: "FINAL_URLS",
	15: "FINAL_MOBILE_URLS",
	16: "TRACKING_URL",
	17: "ANDROID_APP_LINK",
	18: "SIMILAR_DEAL_IDS",
	19: "IOS_APP_LINK",
	20: "IOS_APP_STORE_ID",
}
var LocalPlaceholderFieldEnum_LocalPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":          0,
	"UNKNOWN":              1,
	"DEAL_ID":              2,
	"DEAL_NAME":            3,
	"SUBTITLE":             4,
	"DESCRIPTION":          5,
	"PRICE":                6,
	"FORMATTED_PRICE":      7,
	"SALE_PRICE":           8,
	"FORMATTED_SALE_PRICE": 9,
	"IMAGE_URL":            10,
	"ADDRESS":              11,
	"CATEGORY":             12,
	"CONTEXTUAL_KEYWORDS":  13,
	"FINAL_URLS":           14,
	"FINAL_MOBILE_URLS":    15,
	"TRACKING_URL":         16,
	"ANDROID_APP_LINK":     17,
	"SIMILAR_DEAL_IDS":     18,
	"IOS_APP_LINK":         19,
	"IOS_APP_STORE_ID":     20,
}

func (x LocalPlaceholderFieldEnum_LocalPlaceholderField) String() string {
	return proto.EnumName(LocalPlaceholderFieldEnum_LocalPlaceholderField_name, int32(x))
}
func (LocalPlaceholderFieldEnum_LocalPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_local_placeholder_field_acfbfe8fb3d3dc88, []int{0, 0}
}

// Values for Local placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type LocalPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalPlaceholderFieldEnum) Reset()         { *m = LocalPlaceholderFieldEnum{} }
func (m *LocalPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*LocalPlaceholderFieldEnum) ProtoMessage()    {}
func (*LocalPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_local_placeholder_field_acfbfe8fb3d3dc88, []int{0}
}
func (m *LocalPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *LocalPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *LocalPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalPlaceholderFieldEnum.Merge(dst, src)
}
func (m *LocalPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_LocalPlaceholderFieldEnum.Size(m)
}
func (m *LocalPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LocalPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LocalPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.LocalPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.LocalPlaceholderFieldEnum_LocalPlaceholderField", LocalPlaceholderFieldEnum_LocalPlaceholderField_name, LocalPlaceholderFieldEnum_LocalPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/local_placeholder_field.proto", fileDescriptor_local_placeholder_field_acfbfe8fb3d3dc88)
}

var fileDescriptor_local_placeholder_field_acfbfe8fb3d3dc88 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xdd, 0x6e, 0xda, 0x30,
	0x14, 0x1e, 0xb0, 0xfe, 0x60, 0x68, 0x71, 0x0d, 0xd5, 0xb6, 0x6a, 0xbd, 0x68, 0x1f, 0x20, 0x11,
	0xda, 0x5d, 0x7a, 0x65, 0x12, 0x83, 0x2c, 0x82, 0x13, 0xd9, 0x0e, 0x5d, 0x27, 0xa4, 0x28, 0x23,
	0x59, 0x86, 0x14, 0x62, 0x44, 0x68, 0x1f, 0x62, 0x8f, 0xb1, 0xab, 0x69, 0x8f, 0xb2, 0x47, 0xd9,
	0xcd, 0x5e, 0x61, 0x72, 0xc2, 0xcf, 0x2e, 0xba, 0xde, 0x58, 0xe7, 0x9c, 0xef, 0xe7, 0x1c, 0xf9,
	0x1c, 0x70, 0x97, 0x2a, 0x95, 0x66, 0x89, 0x19, 0xc5, 0x85, 0x59, 0x85, 0x3a, 0x7a, 0xea, 0x9b,
	0x49, 0xfe, 0xb8, 0x2c, 0xcc, 0x4c, 0xcd, 0xa3, 0x2c, 0x5c, 0x65, 0xd1, 0x3c, 0xf9, 0xaa, 0xb2,
	0x38, 0x59, 0x87, 0x5f, 0x16, 0x49, 0x16, 0x1b, 0xab, 0xb5, 0xda, 0x28, 0x74, 0x5d, 0x29, 0x8c,
	0x28, 0x2e, 0x8c, 0xbd, 0xd8, 0x78, 0xea, 0x1b, 0xa5, 0xf8, 0xea, 0xfd, 0xce, 0x7b, 0xb5, 0x30,
	0xa3, 0x3c, 0x57, 0x9b, 0x68, 0xb3, 0x50, 0x79, 0x51, 0x89, 0x6f, 0x7f, 0x34, 0xc0, 0x3b, 0x57,
	0xdb, 0xfb, 0x07, 0xf7, 0xa1, 0x36, 0x27, 0xf9, 0xe3, 0xf2, 0xf6, 0x5b, 0x03, 0x5c, 0x3e, 0x8b,
	0xa2, 0x0e, 0x68, 0x05, 0x4c, 0xf8, 0xc4, 0xa6, 0x43, 0x4a, 0x1c, 0xf8, 0x0a, 0xb5, 0xc0, 0x49,
	0xc0, 0xc6, 0xcc, 0xbb, 0x67, 0xb0, 0xa6, 0x13, 0x87, 0x60, 0x37, 0xa4, 0x0e, 0xac, 0xa3, 0x33,
	0xd0, 0x2c, 0x13, 0x86, 0x27, 0x04, 0x36, 0x50, 0x1b, 0x9c, 0x8a, 0x60, 0x20, 0xa9, 0x74, 0x09,
	0x7c, 0xad, 0x7d, 0x1c, 0x22, 0x6c, 0x4e, 0x7d, 0x49, 0x3d, 0x06, 0x8f, 0x50, 0x13, 0x1c, 0xf9,
	0x9c, 0xda, 0x04, 0x1e, 0xa3, 0x2e, 0xe8, 0x0c, 0x3d, 0x3e, 0xc1, 0x52, 0x12, 0x27, 0xac, 0x8a,
	0x27, 0xe8, 0x1c, 0x00, 0x81, 0x5d, 0xb2, 0xcd, 0x4f, 0xd1, 0x5b, 0xd0, 0x3b, 0x90, 0xfe, 0x41,
	0x9a, 0xba, 0x2f, 0x9d, 0xe0, 0x11, 0x09, 0x03, 0xee, 0x42, 0xa0, 0x67, 0xc2, 0x8e, 0xc3, 0x89,
	0x10, 0xb0, 0xa5, 0x87, 0xb0, 0xb1, 0x24, 0x23, 0x8f, 0x3f, 0xc0, 0x36, 0x7a, 0x03, 0xba, 0xb6,
	0xc7, 0x24, 0xf9, 0x28, 0x03, 0xec, 0x86, 0x63, 0xf2, 0x70, 0xef, 0x71, 0x47, 0xc0, 0x33, 0xdd,
	0x6c, 0x48, 0x19, 0x76, 0xb5, 0x85, 0x80, 0xe7, 0xe8, 0x12, 0x5c, 0x54, 0xf9, 0xc4, 0x1b, 0x50,
	0x97, 0x54, 0xe5, 0x0e, 0x82, 0xa0, 0x2d, 0x39, 0xb6, 0xc7, 0x94, 0x8d, 0xca, 0x66, 0x10, 0xf5,
	0x00, 0xc4, 0xcc, 0xe1, 0x1e, 0x75, 0x42, 0xec, 0xfb, 0xa1, 0x4b, 0xd9, 0x18, 0x5e, 0xe8, 0xaa,
	0xa0, 0x13, 0xea, 0x62, 0x1e, 0x6e, 0xbf, 0x47, 0x40, 0xa4, 0xd5, 0xd4, 0x13, 0x07, 0x5e, 0x57,
	0xf3, 0x76, 0x15, 0x21, 0x3d, 0x4e, 0xf4, 0x3f, 0xf6, 0x06, 0x7f, 0x6a, 0xe0, 0x66, 0xae, 0x96,
	0xc6, 0x8b, 0xeb, 0x1e, 0x5c, 0x3d, 0xbb, 0x2f, 0x5f, 0x2f, 0xdb, 0xaf, 0x7d, 0x1a, 0x6c, 0xc5,
	0xa9, 0xca, 0xa2, 0x3c, 0x35, 0xd4, 0x3a, 0x35, 0xd3, 0x24, 0x2f, 0x4f, 0x61, 0x77, 0x78, 0xab,
	0x45, 0xf1, 0x9f, 0x3b, 0xbc, 0x2b, 0xdf, 0xef, 0xf5, 0xc6, 0x08, 0xe3, 0x9f, 0xf5, 0xeb, 0x51,
	0x65, 0x85, 0xe3, 0xc2, 0xa8, 0x42, 0x1d, 0x4d, 0xfb, 0x86, 0xbe, 0x9c, 0xe2, 0xd7, 0x0e, 0x9f,
	0xe1, 0xb8, 0x98, 0xed, 0xf1, 0xd9, 0xb4, 0x3f, 0x2b, 0xf1, 0xdf, 0xf5, 0x9b, 0xaa, 0x68, 0x59,
	0x38, 0x2e, 0x2c, 0x6b, 0xcf, 0xb0, 0xac, 0x69, 0xdf, 0xb2, 0x4a, 0xce, 0xe7, 0xe3, 0x72, 0xb0,
	0x0f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x7d, 0xf3, 0x33, 0x1f, 0x03, 0x00, 0x00,
}