// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.21.12
// source: relay.proto

package relay

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{0}
}

type CapacityStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UtilizationRatio float32 `protobuf:"fixed32,1,opt,name=utilization_ratio,json=utilizationRatio,proto3" json:"utilization_ratio,omitempty"`
}

func (x *CapacityStatus) Reset() {
	*x = CapacityStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CapacityStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CapacityStatus) ProtoMessage() {}

func (x *CapacityStatus) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CapacityStatus.ProtoReflect.Descriptor instead.
func (*CapacityStatus) Descriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{1}
}

func (x *CapacityStatus) GetUtilizationRatio() float32 {
	if x != nil {
		return x.UtilizationRatio
	}
	return 0
}

type Usage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usage           uint64  `protobuf:"varint,1,opt,name=usage,proto3" json:"usage,omitempty"`
	UtilizationRate float32 `protobuf:"fixed32,2,opt,name=utilization_rate,json=utilizationRate,proto3" json:"utilization_rate,omitempty"` // how much of the alloted bandwidth has the bidder used
}

func (x *Usage) Reset() {
	*x = Usage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usage) ProtoMessage() {}

func (x *Usage) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usage.ProtoReflect.Descriptor instead.
func (*Usage) Descriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{2}
}

func (x *Usage) GetUsage() uint64 {
	if x != nil {
		return x.Usage
	}
	return 0
}

func (x *Usage) GetUtilizationRate() float32 {
	if x != nil {
		return x.UtilizationRate
	}
	return 0
}

type EndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Pubkey      []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Nonce       []byte `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Signature   []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *EndpointRequest) Reset() {
	*x = EndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointRequest) ProtoMessage() {}

func (x *EndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointRequest.ProtoReflect.Descriptor instead.
func (*EndpointRequest) Descriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{3}
}

func (x *EndpointRequest) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *EndpointRequest) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *EndpointRequest) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *EndpointRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type EndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url []string `protobuf:"bytes,1,rep,name=url,proto3" json:"url,omitempty"`
}

func (x *EndpointResponse) Reset() {
	*x = EndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointResponse) ProtoMessage() {}

func (x *EndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointResponse.ProtoReflect.Descriptor instead.
func (*EndpointResponse) Descriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{4}
}

func (x *EndpointResponse) GetUrl() []string {
	if x != nil {
		return x.Url
	}
	return nil
}

type ApiDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	FeatureFlag uint64 `protobuf:"varint,3,opt,name=feature_flag,json=featureFlag,proto3" json:"feature_flag,omitempty"`
}

func (x *ApiDescription) Reset() {
	*x = ApiDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiDescription) ProtoMessage() {}

func (x *ApiDescription) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiDescription.ProtoReflect.Descriptor instead.
func (*ApiDescription) Descriptor() ([]byte, []int) {
	return file_relay_proto_rawDescGZIP(), []int{5}
}

func (x *ApiDescription) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ApiDescription) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApiDescription) GetFeatureFlag() uint64 {
	if x != nil {
		return x.FeatureFlag
	}
	return 0
}

var File_relay_proto protoreflect.FileDescriptor

var file_relay_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x3d, 0x0a,
	0x0e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2b, 0x0a, 0x11, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x75, 0x74, 0x69, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x22, 0x48, 0x0a, 0x05,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x75,
	0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x75, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x22, 0x7f, 0x0a, 0x0f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x75, 0x62, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x75, 0x62,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x24, 0x0a, 0x10, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x6b, 0x0a,
	0x0e, 0x41, 0x70, 0x69, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x32, 0x3f, 0x0a, 0x08, 0x43, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x15, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x30, 0x01, 0x32, 0x5e, 0x0a, 0x05, 0x4d,
	0x65, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12,
	0x2a, 0x0a, 0x08, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x2e, 0x72, 0x65,
	0x6c, 0x61, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x32, 0x53, 0x0a, 0x08, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x47, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x65, 0x61, 0x72, 0x4e, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x32, 0x3b, 0x0a, 0x0b, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x41, 0x70, 0x69,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x33, 0x5a,
	0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x70, 0x61, 0x64, 0x2f, 0x73, 0x6f, 0x6c, 0x70, 0x69, 0x70, 0x65, 0x2d, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_proto_rawDescOnce sync.Once
	file_relay_proto_rawDescData = file_relay_proto_rawDesc
)

func file_relay_proto_rawDescGZIP() []byte {
	file_relay_proto_rawDescOnce.Do(func() {
		file_relay_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_proto_rawDescData)
	})
	return file_relay_proto_rawDescData
}

var file_relay_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_relay_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: relay.Empty
	(*CapacityStatus)(nil),   // 1: relay.CapacityStatus
	(*Usage)(nil),            // 2: relay.Usage
	(*EndpointRequest)(nil),  // 3: relay.EndpointRequest
	(*EndpointResponse)(nil), // 4: relay.EndpointResponse
	(*ApiDescription)(nil),   // 5: relay.ApiDescription
}
var file_relay_proto_depIdxs = []int32{
	0, // 0: relay.Capacity.OnStatus:input_type -> relay.Empty
	0, // 1: relay.Meter.GetStatus:input_type -> relay.Empty
	0, // 2: relay.Meter.OnStatus:input_type -> relay.Empty
	3, // 3: relay.Endpoint.GetClearNetAddress:input_type -> relay.EndpointRequest
	0, // 4: relay.Information.Get:input_type -> relay.Empty
	1, // 5: relay.Capacity.OnStatus:output_type -> relay.CapacityStatus
	2, // 6: relay.Meter.GetStatus:output_type -> relay.Usage
	2, // 7: relay.Meter.OnStatus:output_type -> relay.Usage
	4, // 8: relay.Endpoint.GetClearNetAddress:output_type -> relay.EndpointResponse
	5, // 9: relay.Information.Get:output_type -> relay.ApiDescription
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_relay_proto_init() }
func file_relay_proto_init() {
	if File_relay_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_relay_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CapacityStatus); i {
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
		file_relay_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usage); i {
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
		file_relay_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointRequest); i {
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
		file_relay_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointResponse); i {
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
		file_relay_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiDescription); i {
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
			RawDescriptor: file_relay_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   4,
		},
		GoTypes:           file_relay_proto_goTypes,
		DependencyIndexes: file_relay_proto_depIdxs,
		MessageInfos:      file_relay_proto_msgTypes,
	}.Build()
	File_relay_proto = out.File
	file_relay_proto_rawDesc = nil
	file_relay_proto_goTypes = nil
	file_relay_proto_depIdxs = nil
}