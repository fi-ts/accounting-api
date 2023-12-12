// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: metalstack/io/accounting/api/v1/network_traffic.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NetworkTraffic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	In     uint64 `protobuf:"varint,2,opt,name=in,proto3" json:"in,omitempty"`
	Out    uint64 `protobuf:"varint,3,opt,name=out,proto3" json:"out,omitempty"`
	// Deprecated: Marked as deprecated in metalstack/io/accounting/api/v1/network_traffic.proto.
	Total uint64 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *NetworkTraffic) Reset() {
	*x = NetworkTraffic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkTraffic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkTraffic) ProtoMessage() {}

func (x *NetworkTraffic) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkTraffic.ProtoReflect.Descriptor instead.
func (*NetworkTraffic) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkTraffic) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *NetworkTraffic) GetIn() uint64 {
	if x != nil {
		return x.In
	}
	return 0
}

func (x *NetworkTraffic) GetOut() uint64 {
	if x != nil {
		return x.Out
	}
	return 0
}

// Deprecated: Marked as deprecated in metalstack/io/accounting/api/v1/network_traffic.proto.
func (x *NetworkTraffic) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type NetworkTrafficReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report         *Report           `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
	TrafficReports []*NetworkTraffic `protobuf:"bytes,2,rep,name=traffic_reports,json=trafficReports,proto3" json:"traffic_reports,omitempty"`
}

func (x *NetworkTrafficReport) Reset() {
	*x = NetworkTrafficReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkTrafficReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkTrafficReport) ProtoMessage() {}

func (x *NetworkTrafficReport) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkTrafficReport.ProtoReflect.Descriptor instead.
func (*NetworkTrafficReport) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkTrafficReport) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

func (x *NetworkTrafficReport) GetTrafficReports() []*NetworkTraffic {
	if x != nil {
		return x.TrafficReports
	}
	return nil
}

type NetworkUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query  *UsageQuery `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Device *string     `protobuf:"bytes,2,opt,name=device,proto3,oneof" json:"device,omitempty"`
}

func (x *NetworkUsageRequest) Reset() {
	*x = NetworkUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkUsageRequest) ProtoMessage() {}

func (x *NetworkUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkUsageRequest.ProtoReflect.Descriptor instead.
func (*NetworkUsageRequest) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkUsageRequest) GetQuery() *UsageQuery {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *NetworkUsageRequest) GetDevice() string {
	if x != nil && x.Device != nil {
		return *x.Device
	}
	return ""
}

type NetworkUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From             *timestamppb.Timestamp    `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To               *timestamppb.Timestamp    `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Usage            []*NetworkUsage           `protobuf:"bytes,3,rep,name=usage,proto3" json:"usage,omitempty"`
	AccumulatedUsage *NetworkUsageAccumuluated `protobuf:"bytes,4,opt,name=accumulated_usage,json=accumulatedUsage,proto3" json:"accumulated_usage,omitempty"`
}

func (x *NetworkUsageResponse) Reset() {
	*x = NetworkUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkUsageResponse) ProtoMessage() {}

func (x *NetworkUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkUsageResponse.ProtoReflect.Descriptor instead.
func (*NetworkUsageResponse) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkUsageResponse) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *NetworkUsageResponse) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *NetworkUsageResponse) GetUsage() []*NetworkUsage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *NetworkUsageResponse) GetAccumulatedUsage() *NetworkUsageAccumuluated {
	if x != nil {
		return x.AccumulatedUsage
	}
	return nil
}

type NetworkUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partition      string               `protobuf:"bytes,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Tenant         string               `protobuf:"bytes,2,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Projectid      string               `protobuf:"bytes,3,opt,name=projectid,proto3" json:"projectid,omitempty"`
	Clusterid      string               `protobuf:"bytes,4,opt,name=clusterid,proto3" json:"clusterid,omitempty"`
	Clustername    string               `protobuf:"bytes,5,opt,name=clustername,proto3" json:"clustername,omitempty"`
	Device         string               `protobuf:"bytes,6,opt,name=device,proto3" json:"device,omitempty"`
	InBytes        string               `protobuf:"bytes,7,opt,name=in_bytes,json=inBytes,proto3" json:"in_bytes,omitempty"`
	OutBytes       string               `protobuf:"bytes,8,opt,name=out_bytes,json=outBytes,proto3" json:"out_bytes,omitempty"`
	Lifetime       *durationpb.Duration `protobuf:"bytes,11,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	Warnings       []string             `protobuf:"bytes,12,rep,name=warnings,proto3" json:"warnings,omitempty"`
	TotalBytes     string               `protobuf:"bytes,13,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	Projectname    string               `protobuf:"bytes,14,opt,name=projectname,proto3" json:"projectname,omitempty"`
	Annotations    []string             `protobuf:"bytes,15,rep,name=annotations,proto3" json:"annotations,omitempty"`
	Tenantname     string               `protobuf:"bytes,16,opt,name=tenantname,proto3" json:"tenantname,omitempty"`
	Contractnumber string               `protobuf:"bytes,17,opt,name=contractnumber,proto3" json:"contractnumber,omitempty"`
	Debtorid       string               `protobuf:"bytes,18,opt,name=debtorid,proto3" json:"debtorid,omitempty"`
}

func (x *NetworkUsage) Reset() {
	*x = NetworkUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkUsage) ProtoMessage() {}

func (x *NetworkUsage) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkUsage.ProtoReflect.Descriptor instead.
func (*NetworkUsage) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescGZIP(), []int{4}
}

func (x *NetworkUsage) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *NetworkUsage) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *NetworkUsage) GetProjectid() string {
	if x != nil {
		return x.Projectid
	}
	return ""
}

func (x *NetworkUsage) GetClusterid() string {
	if x != nil {
		return x.Clusterid
	}
	return ""
}

func (x *NetworkUsage) GetClustername() string {
	if x != nil {
		return x.Clustername
	}
	return ""
}

func (x *NetworkUsage) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *NetworkUsage) GetInBytes() string {
	if x != nil {
		return x.InBytes
	}
	return ""
}

func (x *NetworkUsage) GetOutBytes() string {
	if x != nil {
		return x.OutBytes
	}
	return ""
}

func (x *NetworkUsage) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

func (x *NetworkUsage) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *NetworkUsage) GetTotalBytes() string {
	if x != nil {
		return x.TotalBytes
	}
	return ""
}

func (x *NetworkUsage) GetProjectname() string {
	if x != nil {
		return x.Projectname
	}
	return ""
}

func (x *NetworkUsage) GetAnnotations() []string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *NetworkUsage) GetTenantname() string {
	if x != nil {
		return x.Tenantname
	}
	return ""
}

func (x *NetworkUsage) GetContractnumber() string {
	if x != nil {
		return x.Contractnumber
	}
	return ""
}

func (x *NetworkUsage) GetDebtorid() string {
	if x != nil {
		return x.Debtorid
	}
	return ""
}

type NetworkUsageAccumuluated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lifetime   *durationpb.Duration `protobuf:"bytes,1,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	InBytes    string               `protobuf:"bytes,2,opt,name=in_bytes,json=inBytes,proto3" json:"in_bytes,omitempty"`
	OutBytes   string               `protobuf:"bytes,3,opt,name=out_bytes,json=outBytes,proto3" json:"out_bytes,omitempty"`
	TotalBytes string               `protobuf:"bytes,4,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
}

func (x *NetworkUsageAccumuluated) Reset() {
	*x = NetworkUsageAccumuluated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkUsageAccumuluated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkUsageAccumuluated) ProtoMessage() {}

func (x *NetworkUsageAccumuluated) ProtoReflect() protoreflect.Message {
	mi := &file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkUsageAccumuluated.ProtoReflect.Descriptor instead.
func (*NetworkUsageAccumuluated) Descriptor() ([]byte, []int) {
	return file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescGZIP(), []int{5}
}

func (x *NetworkUsageAccumuluated) GetLifetime() *durationpb.Duration {
	if x != nil {
		return x.Lifetime
	}
	return nil
}

func (x *NetworkUsageAccumuluated) GetInBytes() string {
	if x != nil {
		return x.InBytes
	}
	return ""
}

func (x *NetworkUsageAccumuluated) GetOutBytes() string {
	if x != nil {
		return x.OutBytes
	}
	return ""
}

func (x *NetworkUsageAccumuluated) GetTotalBytes() string {
	if x != nil {
		return x.TotalBytes
	}
	return ""
}

var File_metalstack_io_accounting_api_v1_network_traffic_proto protoreflect.FileDescriptor

var file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDesc = []byte{
	0x0a, 0x35, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2f, 0x69, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x0e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x02, 0x18, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xb1, 0x01,
	0x0a, 0x14, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63,
	0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74,
	0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x58, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x66, 0x66,
	0x69, 0x63, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x22, 0x80, 0x01, 0x0a, 0x13, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x05, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x9f, 0x02, 0x0a, 0x14, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x43, 0x0a, 0x05, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x66,
	0x0a, 0x11, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x10, 0x61, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8e, 0x04, 0x0a, 0x0c, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6f, 0x75, 0x74, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x6c,
	0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x65, 0x62, 0x74, 0x6f, 0x72, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x65, 0x62, 0x74, 0x6f, 0x72, 0x69, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x18, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x6e, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69,
	0x6e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x75, 0x74, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x32, 0xf8, 0x01, 0x0a, 0x15, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x69,
	0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x35, 0x2e, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x1a, 0x26, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x74, 0x0a, 0x05, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x34, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e,
	0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x0d, 0x5a, 0x0b, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescOnce sync.Once
	file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescData = file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDesc
)

func file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescGZIP() []byte {
	file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescOnce.Do(func() {
		file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescData = protoimpl.X.CompressGZIP(file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescData)
	})
	return file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDescData
}

var file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_metalstack_io_accounting_api_v1_network_traffic_proto_goTypes = []interface{}{
	(*NetworkTraffic)(nil),           // 0: metalstack.io.accounting.api.v1.NetworkTraffic
	(*NetworkTrafficReport)(nil),     // 1: metalstack.io.accounting.api.v1.NetworkTrafficReport
	(*NetworkUsageRequest)(nil),      // 2: metalstack.io.accounting.api.v1.NetworkUsageRequest
	(*NetworkUsageResponse)(nil),     // 3: metalstack.io.accounting.api.v1.NetworkUsageResponse
	(*NetworkUsage)(nil),             // 4: metalstack.io.accounting.api.v1.NetworkUsage
	(*NetworkUsageAccumuluated)(nil), // 5: metalstack.io.accounting.api.v1.NetworkUsageAccumuluated
	(*Report)(nil),                   // 6: metalstack.io.accounting.api.v1.Report
	(*UsageQuery)(nil),               // 7: metalstack.io.accounting.api.v1.UsageQuery
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),      // 9: google.protobuf.Duration
	(*Empty)(nil),                    // 10: metalstack.io.accounting.api.v1.Empty
}
var file_metalstack_io_accounting_api_v1_network_traffic_proto_depIdxs = []int32{
	6,  // 0: metalstack.io.accounting.api.v1.NetworkTrafficReport.report:type_name -> metalstack.io.accounting.api.v1.Report
	0,  // 1: metalstack.io.accounting.api.v1.NetworkTrafficReport.traffic_reports:type_name -> metalstack.io.accounting.api.v1.NetworkTraffic
	7,  // 2: metalstack.io.accounting.api.v1.NetworkUsageRequest.query:type_name -> metalstack.io.accounting.api.v1.UsageQuery
	8,  // 3: metalstack.io.accounting.api.v1.NetworkUsageResponse.from:type_name -> google.protobuf.Timestamp
	8,  // 4: metalstack.io.accounting.api.v1.NetworkUsageResponse.to:type_name -> google.protobuf.Timestamp
	4,  // 5: metalstack.io.accounting.api.v1.NetworkUsageResponse.usage:type_name -> metalstack.io.accounting.api.v1.NetworkUsage
	5,  // 6: metalstack.io.accounting.api.v1.NetworkUsageResponse.accumulated_usage:type_name -> metalstack.io.accounting.api.v1.NetworkUsageAccumuluated
	9,  // 7: metalstack.io.accounting.api.v1.NetworkUsage.lifetime:type_name -> google.protobuf.Duration
	9,  // 8: metalstack.io.accounting.api.v1.NetworkUsageAccumuluated.lifetime:type_name -> google.protobuf.Duration
	1,  // 9: metalstack.io.accounting.api.v1.NetworkTrafficService.Modified:input_type -> metalstack.io.accounting.api.v1.NetworkTrafficReport
	2,  // 10: metalstack.io.accounting.api.v1.NetworkTrafficService.Usage:input_type -> metalstack.io.accounting.api.v1.NetworkUsageRequest
	10, // 11: metalstack.io.accounting.api.v1.NetworkTrafficService.Modified:output_type -> metalstack.io.accounting.api.v1.Empty
	3,  // 12: metalstack.io.accounting.api.v1.NetworkTrafficService.Usage:output_type -> metalstack.io.accounting.api.v1.NetworkUsageResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_metalstack_io_accounting_api_v1_network_traffic_proto_init() }
func file_metalstack_io_accounting_api_v1_network_traffic_proto_init() {
	if File_metalstack_io_accounting_api_v1_network_traffic_proto != nil {
		return
	}
	file_metalstack_io_accounting_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkTraffic); i {
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
		file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkTrafficReport); i {
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
		file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkUsageRequest); i {
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
		file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkUsageResponse); i {
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
		file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkUsage); i {
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
		file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkUsageAccumuluated); i {
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
	file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metalstack_io_accounting_api_v1_network_traffic_proto_goTypes,
		DependencyIndexes: file_metalstack_io_accounting_api_v1_network_traffic_proto_depIdxs,
		MessageInfos:      file_metalstack_io_accounting_api_v1_network_traffic_proto_msgTypes,
	}.Build()
	File_metalstack_io_accounting_api_v1_network_traffic_proto = out.File
	file_metalstack_io_accounting_api_v1_network_traffic_proto_rawDesc = nil
	file_metalstack_io_accounting_api_v1_network_traffic_proto_goTypes = nil
	file_metalstack_io_accounting_api_v1_network_traffic_proto_depIdxs = nil
}
