// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: block_certificate.proto

package pb

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

// A certificate proofing the validity of block meta information.
type BlockCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The aggregated signature of the certification committee active during the
	// block's creation, certifying the block's meta information.
	Signature *AggregatedSignature `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The chain ID of the chain the block is part of.
	ChainId uint64 `protobuf:"varint,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	// The block number.
	Number uint64 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	// The 32-byte hash of the block.
	Hash []byte `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	// The 32-byte hash of the state root of the block.
	StateRoot []byte `protobuf:"bytes,5,opt,name=stateRoot,proto3" json:"stateRoot,omitempty"`
}

func (x *BlockCertificate) Reset() {
	*x = BlockCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockCertificate) ProtoMessage() {}

func (x *BlockCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_block_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockCertificate.ProtoReflect.Descriptor instead.
func (*BlockCertificate) Descriptor() ([]byte, []int) {
	return file_block_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *BlockCertificate) GetSignature() *AggregatedSignature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *BlockCertificate) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *BlockCertificate) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BlockCertificate) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BlockCertificate) GetStateRoot() []byte {
	if x != nil {
		return x.StateRoot
	}
	return nil
}

var File_block_certificate_proto protoreflect.FileDescriptor

var file_block_certificate_proto_rawDesc = []byte{
	0x0a, 0x17, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x2e, 0x73, 0x63, 0x63, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbf, 0x01, 0x0a, 0x10, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63,
	0x2e, 0x73, 0x63, 0x63, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x52, 0x6f,
	0x6f, 0x74, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_block_certificate_proto_rawDescOnce sync.Once
	file_block_certificate_proto_rawDescData = file_block_certificate_proto_rawDesc
)

func file_block_certificate_proto_rawDescGZIP() []byte {
	file_block_certificate_proto_rawDescOnce.Do(func() {
		file_block_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_certificate_proto_rawDescData)
	})
	return file_block_certificate_proto_rawDescData
}

var file_block_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_block_certificate_proto_goTypes = []interface{}{
	(*BlockCertificate)(nil),    // 0: sonic.scc.cert.proto.BlockCertificate
	(*AggregatedSignature)(nil), // 1: sonic.scc.cert.proto.AggregatedSignature
}
var file_block_certificate_proto_depIdxs = []int32{
	1, // 0: sonic.scc.cert.proto.BlockCertificate.signature:type_name -> sonic.scc.cert.proto.AggregatedSignature
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_block_certificate_proto_init() }
func file_block_certificate_proto_init() {
	if File_block_certificate_proto != nil {
		return
	}
	file_signature_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_block_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockCertificate); i {
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
			RawDescriptor: file_block_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_block_certificate_proto_goTypes,
		DependencyIndexes: file_block_certificate_proto_depIdxs,
		MessageInfos:      file_block_certificate_proto_msgTypes,
	}.Build()
	File_block_certificate_proto = out.File
	file_block_certificate_proto_rawDesc = nil
	file_block_certificate_proto_goTypes = nil
	file_block_certificate_proto_depIdxs = nil
}
