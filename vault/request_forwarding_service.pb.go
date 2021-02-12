// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: vault/request_forwarding_service.proto

package vault

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	forwarding "github.com/hashicorp/vault/helper/forwarding"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// ClusterAddr is used to send up a standby node's address to the active
	// node upon heartbeat
	ClusterAddr string `protobuf:"bytes,2,opt,name=cluster_addr,json=clusterAddr,proto3" json:"cluster_addr,omitempty"`
	// ClusterAddrs is used to send up a list of cluster addresses to a dr
	// primary from a dr secondary
	ClusterAddrs     []string         `protobuf:"bytes,3,rep,name=cluster_addrs,json=clusterAddrs,proto3" json:"cluster_addrs,omitempty"`
	RaftAppliedIndex uint64           `protobuf:"varint,4,opt,name=raft_applied_index,json=raftAppliedIndex,proto3" json:"raft_applied_index,omitempty"`
	RaftNodeID       string           `protobuf:"bytes,5,opt,name=raft_node_id,json=raftNodeId,proto3" json:"raft_node_id,omitempty"`
	NodeInfo         *NodeInformation `protobuf:"bytes,6,opt,name=node_info,json=nodeInfo,proto3" json:"node_info,omitempty"`
	RaftTerm         uint64           `protobuf:"varint,7,opt,name=raft_term,json=raftTerm,proto3" json:"raft_term,omitempty"`
	RaftNonVoter     bool             `protobuf:"varint,8,opt,name=raft_non_voter,json=raftNonVoter,proto3" json:"raft_non_voter,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_request_forwarding_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vault_request_forwarding_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_vault_request_forwarding_service_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EchoRequest) GetClusterAddr() string {
	if x != nil {
		return x.ClusterAddr
	}
	return ""
}

func (x *EchoRequest) GetClusterAddrs() []string {
	if x != nil {
		return x.ClusterAddrs
	}
	return nil
}

func (x *EchoRequest) GetRaftAppliedIndex() uint64 {
	if x != nil {
		return x.RaftAppliedIndex
	}
	return 0
}

func (x *EchoRequest) GetRaftNodeID() string {
	if x != nil {
		return x.RaftNodeID
	}
	return ""
}

func (x *EchoRequest) GetNodeInfo() *NodeInformation {
	if x != nil {
		return x.NodeInfo
	}
	return nil
}

func (x *EchoRequest) GetRaftTerm() uint64 {
	if x != nil {
		return x.RaftTerm
	}
	return 0
}

func (x *EchoRequest) GetRaftNonVoter() bool {
	if x != nil {
		return x.RaftNonVoter
	}
	return false
}

type EchoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message          string           `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ClusterAddrs     []string         `protobuf:"bytes,2,rep,name=cluster_addrs,json=clusterAddrs,proto3" json:"cluster_addrs,omitempty"`
	ReplicationState uint32           `protobuf:"varint,3,opt,name=replication_state,json=replicationState,proto3" json:"replication_state,omitempty"`
	RaftAppliedIndex uint64           `protobuf:"varint,4,opt,name=raft_applied_index,json=raftAppliedIndex,proto3" json:"raft_applied_index,omitempty"`
	RaftNodeID       string           `protobuf:"bytes,5,opt,name=raft_node_id,json=raftNodeId,proto3" json:"raft_node_id,omitempty"`
	NodeInfo         *NodeInformation `protobuf:"bytes,6,opt,name=node_info,json=nodeInfo,proto3" json:"node_info,omitempty"`
}

func (x *EchoReply) Reset() {
	*x = EchoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_request_forwarding_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoReply) ProtoMessage() {}

func (x *EchoReply) ProtoReflect() protoreflect.Message {
	mi := &file_vault_request_forwarding_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoReply.ProtoReflect.Descriptor instead.
func (*EchoReply) Descriptor() ([]byte, []int) {
	return file_vault_request_forwarding_service_proto_rawDescGZIP(), []int{1}
}

func (x *EchoReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EchoReply) GetClusterAddrs() []string {
	if x != nil {
		return x.ClusterAddrs
	}
	return nil
}

func (x *EchoReply) GetReplicationState() uint32 {
	if x != nil {
		return x.ReplicationState
	}
	return 0
}

func (x *EchoReply) GetRaftAppliedIndex() uint64 {
	if x != nil {
		return x.RaftAppliedIndex
	}
	return 0
}

func (x *EchoReply) GetRaftNodeID() string {
	if x != nil {
		return x.RaftNodeID
	}
	return ""
}

func (x *EchoReply) GetNodeInfo() *NodeInformation {
	if x != nil {
		return x.NodeInfo
	}
	return nil
}

type NodeInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterAddr      string `protobuf:"bytes,1,opt,name=cluster_addr,json=clusterAddr,proto3" json:"cluster_addr,omitempty"`
	ApiAddr          string `protobuf:"bytes,2,opt,name=api_addr,json=apiAddr,proto3" json:"api_addr,omitempty"`
	Mode             string `protobuf:"bytes,3,opt,name=mode,proto3" json:"mode,omitempty"`
	NodeID           string `protobuf:"bytes,4,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	ReplicationState uint32 `protobuf:"varint,5,opt,name=replication_state,json=replicationState,proto3" json:"replication_state,omitempty"`
}

func (x *NodeInformation) Reset() {
	*x = NodeInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_request_forwarding_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInformation) ProtoMessage() {}

func (x *NodeInformation) ProtoReflect() protoreflect.Message {
	mi := &file_vault_request_forwarding_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInformation.ProtoReflect.Descriptor instead.
func (*NodeInformation) Descriptor() ([]byte, []int) {
	return file_vault_request_forwarding_service_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInformation) GetClusterAddr() string {
	if x != nil {
		return x.ClusterAddr
	}
	return ""
}

func (x *NodeInformation) GetApiAddr() string {
	if x != nil {
		return x.ApiAddr
	}
	return ""
}

func (x *NodeInformation) GetMode() string {
	if x != nil {
		return x.Mode
	}
	return ""
}

func (x *NodeInformation) GetNodeID() string {
	if x != nil {
		return x.NodeID
	}
	return ""
}

func (x *NodeInformation) GetReplicationState() uint32 {
	if x != nil {
		return x.ReplicationState
	}
	return 0
}

type ClientKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	X    []byte `protobuf:"bytes,2,opt,name=x,proto3" json:"x,omitempty"`
	Y    []byte `protobuf:"bytes,3,opt,name=y,proto3" json:"y,omitempty"`
	D    []byte `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
}

func (x *ClientKey) Reset() {
	*x = ClientKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_request_forwarding_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientKey) ProtoMessage() {}

func (x *ClientKey) ProtoReflect() protoreflect.Message {
	mi := &file_vault_request_forwarding_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientKey.ProtoReflect.Descriptor instead.
func (*ClientKey) Descriptor() ([]byte, []int) {
	return file_vault_request_forwarding_service_proto_rawDescGZIP(), []int{3}
}

func (x *ClientKey) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ClientKey) GetX() []byte {
	if x != nil {
		return x.X
	}
	return nil
}

func (x *ClientKey) GetY() []byte {
	if x != nil {
		return x.Y
	}
	return nil
}

func (x *ClientKey) GetD() []byte {
	if x != nil {
		return x.D
	}
	return nil
}

type PerfStandbyElectionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PerfStandbyElectionInput) Reset() {
	*x = PerfStandbyElectionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_request_forwarding_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerfStandbyElectionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerfStandbyElectionInput) ProtoMessage() {}

func (x *PerfStandbyElectionInput) ProtoReflect() protoreflect.Message {
	mi := &file_vault_request_forwarding_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerfStandbyElectionInput.ProtoReflect.Descriptor instead.
func (*PerfStandbyElectionInput) Descriptor() ([]byte, []int) {
	return file_vault_request_forwarding_service_proto_rawDescGZIP(), []int{4}
}

type PerfStandbyElectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                 string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClusterID          string     `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	PrimaryClusterAddr string     `protobuf:"bytes,3,opt,name=primary_cluster_addr,json=primaryClusterAddr,proto3" json:"primary_cluster_addr,omitempty"`
	CaCert             []byte     `protobuf:"bytes,4,opt,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	ClientCert         []byte     `protobuf:"bytes,5,opt,name=client_cert,json=clientCert,proto3" json:"client_cert,omitempty"`
	ClientKey          *ClientKey `protobuf:"bytes,6,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
}

func (x *PerfStandbyElectionResponse) Reset() {
	*x = PerfStandbyElectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vault_request_forwarding_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PerfStandbyElectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PerfStandbyElectionResponse) ProtoMessage() {}

func (x *PerfStandbyElectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vault_request_forwarding_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PerfStandbyElectionResponse.ProtoReflect.Descriptor instead.
func (*PerfStandbyElectionResponse) Descriptor() ([]byte, []int) {
	return file_vault_request_forwarding_service_proto_rawDescGZIP(), []int{5}
}

func (x *PerfStandbyElectionResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *PerfStandbyElectionResponse) GetClusterID() string {
	if x != nil {
		return x.ClusterID
	}
	return ""
}

func (x *PerfStandbyElectionResponse) GetPrimaryClusterAddr() string {
	if x != nil {
		return x.PrimaryClusterAddr
	}
	return ""
}

func (x *PerfStandbyElectionResponse) GetCaCert() []byte {
	if x != nil {
		return x.CaCert
	}
	return nil
}

func (x *PerfStandbyElectionResponse) GetClientCert() []byte {
	if x != nil {
		return x.ClientCert
	}
	return nil
}

func (x *PerfStandbyElectionResponse) GetClientKey() *ClientKey {
	if x != nil {
		return x.ClientKey
	}
	return nil
}

var File_vault_request_forwarding_service_proto protoreflect.FileDescriptor

var file_vault_request_forwarding_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x1a,
	0x1d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69,
	0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7,
	0x02, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x73,
	0x12, 0x2c, 0x0a, 0x12, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x72, 0x61,
	0x66, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20,
	0x0a, 0x0c, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x61, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x33, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x74, 0x65,
	0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x61, 0x66, 0x74, 0x54, 0x65,
	0x72, 0x6d, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x5f, 0x76,
	0x6f, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x61, 0x66, 0x74,
	0x4e, 0x6f, 0x6e, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x22, 0xfc, 0x01, 0x0a, 0x09, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10,
	0x72, 0x61, 0x66, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x20, 0x0a, 0x0c, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x61, 0x66, 0x74, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x0f, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x69, 0x41, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x10, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x49, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x79,
	0x12, 0x0c, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x64, 0x22, 0x1a,
	0x0a, 0x18, 0x50, 0x65, 0x72, 0x66, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79, 0x45, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xe9, 0x01, 0x0a, 0x1b, 0x50,
	0x65, 0x72, 0x66, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x62, 0x79, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x63,
	0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x61,
	0x43, 0x65, 0x72, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x43, 0x65, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x61, 0x75, 0x6c,
	0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x32, 0xf0, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x0e,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13,
	0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x04, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x12, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x21, 0x50,
	0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x62,
	0x79, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1f, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x62, 0x79, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x1a, 0x22, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x66, 0x53, 0x74,
	0x61, 0x6e, 0x64, 0x62, 0x79, 0x45, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vault_request_forwarding_service_proto_rawDescOnce sync.Once
	file_vault_request_forwarding_service_proto_rawDescData = file_vault_request_forwarding_service_proto_rawDesc
)

func file_vault_request_forwarding_service_proto_rawDescGZIP() []byte {
	file_vault_request_forwarding_service_proto_rawDescOnce.Do(func() {
		file_vault_request_forwarding_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_vault_request_forwarding_service_proto_rawDescData)
	})
	return file_vault_request_forwarding_service_proto_rawDescData
}

var file_vault_request_forwarding_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vault_request_forwarding_service_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),                 // 0: vault.EchoRequest
	(*EchoReply)(nil),                   // 1: vault.EchoReply
	(*NodeInformation)(nil),             // 2: vault.NodeInformation
	(*ClientKey)(nil),                   // 3: vault.ClientKey
	(*PerfStandbyElectionInput)(nil),    // 4: vault.PerfStandbyElectionInput
	(*PerfStandbyElectionResponse)(nil), // 5: vault.PerfStandbyElectionResponse
	(*forwarding.Request)(nil),          // 6: forwarding.Request
	(*forwarding.Response)(nil),         // 7: forwarding.Response
}
var file_vault_request_forwarding_service_proto_depIDxs = []int32{
	2, // 0: vault.EchoRequest.node_info:type_name -> vault.NodeInformation
	2, // 1: vault.EchoReply.node_info:type_name -> vault.NodeInformation
	3, // 2: vault.PerfStandbyElectionResponse.client_key:type_name -> vault.ClientKey
	6, // 3: vault.RequestForwarding.ForwardRequest:input_type -> forwarding.Request
	0, // 4: vault.RequestForwarding.Echo:input_type -> vault.EchoRequest
	4, // 5: vault.RequestForwarding.PerformanceStandbyElectionRequest:input_type -> vault.PerfStandbyElectionInput
	7, // 6: vault.RequestForwarding.ForwardRequest:output_type -> forwarding.Response
	1, // 7: vault.RequestForwarding.Echo:output_type -> vault.EchoReply
	5, // 8: vault.RequestForwarding.PerformanceStandbyElectionRequest:output_type -> vault.PerfStandbyElectionResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_vault_request_forwarding_service_proto_init() }
func file_vault_request_forwarding_service_proto_init() {
	if File_vault_request_forwarding_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vault_request_forwarding_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_vault_request_forwarding_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoReply); i {
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
		file_vault_request_forwarding_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInformation); i {
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
		file_vault_request_forwarding_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientKey); i {
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
		file_vault_request_forwarding_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerfStandbyElectionInput); i {
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
		file_vault_request_forwarding_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PerfStandbyElectionResponse); i {
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
			RawDescriptor: file_vault_request_forwarding_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vault_request_forwarding_service_proto_goTypes,
		DependencyIndexes: file_vault_request_forwarding_service_proto_depIDxs,
		MessageInfos:      file_vault_request_forwarding_service_proto_msgTypes,
	}.Build()
	File_vault_request_forwarding_service_proto = out.File
	file_vault_request_forwarding_service_proto_rawDesc = nil
	file_vault_request_forwarding_service_proto_goTypes = nil
	file_vault_request_forwarding_service_proto_depIDxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RequestForwardingClient is the client API for RequestForwarding service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RequestForwardingClient interface {
	ForwardRequest(ctx context.Context, in *forwarding.Request, opts ...grpc.CallOption) (*forwarding.Response, error)
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
	PerformanceStandbyElectionRequest(ctx context.Context, in *PerfStandbyElectionInput, opts ...grpc.CallOption) (RequestForwarding_PerformanceStandbyElectionRequestClient, error)
}

type requestForwardingClient struct {
	cc grpc.ClientConnInterface
}

func NewRequestForwardingClient(cc grpc.ClientConnInterface) RequestForwardingClient {
	return &requestForwardingClient{cc}
}

func (c *requestForwardingClient) ForwardRequest(ctx context.Context, in *forwarding.Request, opts ...grpc.CallOption) (*forwarding.Response, error) {
	out := new(forwarding.Response)
	err := c.cc.Invoke(ctx, "/vault.RequestForwarding/ForwardRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestForwardingClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/vault.RequestForwarding/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *requestForwardingClient) PerformanceStandbyElectionRequest(ctx context.Context, in *PerfStandbyElectionInput, opts ...grpc.CallOption) (RequestForwarding_PerformanceStandbyElectionRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RequestForwarding_serviceDesc.Streams[0], "/vault.RequestForwarding/PerformanceStandbyElectionRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &requestForwardingPerformanceStandbyElectionRequestClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RequestForwarding_PerformanceStandbyElectionRequestClient interface {
	Recv() (*PerfStandbyElectionResponse, error)
	grpc.ClientStream
}

type requestForwardingPerformanceStandbyElectionRequestClient struct {
	grpc.ClientStream
}

func (x *requestForwardingPerformanceStandbyElectionRequestClient) Recv() (*PerfStandbyElectionResponse, error) {
	m := new(PerfStandbyElectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RequestForwardingServer is the server API for RequestForwarding service.
type RequestForwardingServer interface {
	ForwardRequest(context.Context, *forwarding.Request) (*forwarding.Response, error)
	Echo(context.Context, *EchoRequest) (*EchoReply, error)
	PerformanceStandbyElectionRequest(*PerfStandbyElectionInput, RequestForwarding_PerformanceStandbyElectionRequestServer) error
}

// UnimplementedRequestForwardingServer can be embedded to have forward compatible implementations.
type UnimplementedRequestForwardingServer struct {
}

func (*UnimplementedRequestForwardingServer) ForwardRequest(context.Context, *forwarding.Request) (*forwarding.Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ForwardRequest not implemented")
}
func (*UnimplementedRequestForwardingServer) Echo(context.Context, *EchoRequest) (*EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (*UnimplementedRequestForwardingServer) PerformanceStandbyElectionRequest(*PerfStandbyElectionInput, RequestForwarding_PerformanceStandbyElectionRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method PerformanceStandbyElectionRequest not implemented")
}

func RegisterRequestForwardingServer(s *grpc.Server, srv RequestForwardingServer) {
	s.RegisterService(&_RequestForwarding_serviceDesc, srv)
}

func _RequestForwarding_ForwardRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(forwarding.Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestForwardingServer).ForwardRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vault.RequestForwarding/ForwardRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestForwardingServer).ForwardRequest(ctx, req.(*forwarding.Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestForwarding_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RequestForwardingServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vault.RequestForwarding/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RequestForwardingServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RequestForwarding_PerformanceStandbyElectionRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PerfStandbyElectionInput)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RequestForwardingServer).PerformanceStandbyElectionRequest(m, &requestForwardingPerformanceStandbyElectionRequestServer{stream})
}

type RequestForwarding_PerformanceStandbyElectionRequestServer interface {
	Send(*PerfStandbyElectionResponse) error
	grpc.ServerStream
}

type requestForwardingPerformanceStandbyElectionRequestServer struct {
	grpc.ServerStream
}

func (x *requestForwardingPerformanceStandbyElectionRequestServer) Send(m *PerfStandbyElectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RequestForwarding_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vault.RequestForwarding",
	HandlerType: (*RequestForwardingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ForwardRequest",
			Handler:    _RequestForwarding_ForwardRequest_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _RequestForwarding_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PerformanceStandbyElectionRequest",
			Handler:       _RequestForwarding_PerformanceStandbyElectionRequest_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vault/request_forwarding_service.proto",
}
