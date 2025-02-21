// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: committee_certificate.proto

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

// A certificate proofing the validity of a committee for a specific period.
type CommitteeCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The aggregated signature of the preceding period's committee, certifying
	// the current committee.
	Signature *AggregatedSignature `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// The chain ID of the chain the committee is active on.
	ChainId uint64 `protobuf:"varint,2,opt,name=chainId,proto3" json:"chainId,omitempty"`
	// The period in which the committee is active.
	Period uint64 `protobuf:"varint,3,opt,name=period,proto3" json:"period,omitempty"`
	// The members of the committee.
	Members []*Member `protobuf:"bytes,4,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *CommitteeCertificate) Reset() {
	*x = CommitteeCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_committee_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitteeCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitteeCertificate) ProtoMessage() {}

func (x *CommitteeCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_committee_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitteeCertificate.ProtoReflect.Descriptor instead.
func (*CommitteeCertificate) Descriptor() ([]byte, []int) {
	return file_committee_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *CommitteeCertificate) GetSignature() *AggregatedSignature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *CommitteeCertificate) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *CommitteeCertificate) GetPeriod() uint64 {
	if x != nil {
		return x.Period
	}
	return 0
}

func (x *CommitteeCertificate) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

// A member of the committee.
type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The 48-byte BLS public key of the member.
	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	// The 96-byte proof of possession of the private key, proofing that at some
	// point in time the member had access to the private key corresponding to
	// the public key.
	ProofOfPossession []byte `protobuf:"bytes,2,opt,name=proofOfPossession,proto3" json:"proofOfPossession,omitempty"`
	// The voting power of the member. This value defines the weight of the
	// member's vote in the committee.
	VotingPower uint64 `protobuf:"varint,3,opt,name=votingPower,proto3" json:"votingPower,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_committee_certificate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_committee_certificate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_committee_certificate_proto_rawDescGZIP(), []int{1}
}

func (x *Member) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Member) GetProofOfPossession() []byte {
	if x != nil {
		return x.ProofOfPossession
	}
	return nil
}

func (x *Member) GetVotingPower() uint64 {
	if x != nil {
		return x.VotingPower
	}
	return 0
}

var File_committee_certificate_proto protoreflect.FileDescriptor

var file_committee_certificate_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73,
	0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x73, 0x63, 0x63, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74,
	0x65, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x47, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x73, 0x6f, 0x6e, 0x69, 0x63, 0x2e, 0x73, 0x63, 0x63, 0x2e, 0x63, 0x65, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x64, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x6e, 0x69,
	0x63, 0x2e, 0x73, 0x63, 0x63, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x22, 0x76, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x6f,
	0x66, 0x4f, 0x66, 0x50, 0x6f, 0x73, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x4f, 0x66, 0x50, 0x6f, 0x73, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x76, 0x6f, 0x74,
	0x69, 0x6e, 0x67, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_committee_certificate_proto_rawDescOnce sync.Once
	file_committee_certificate_proto_rawDescData = file_committee_certificate_proto_rawDesc
)

func file_committee_certificate_proto_rawDescGZIP() []byte {
	file_committee_certificate_proto_rawDescOnce.Do(func() {
		file_committee_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_committee_certificate_proto_rawDescData)
	})
	return file_committee_certificate_proto_rawDescData
}

var file_committee_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_committee_certificate_proto_goTypes = []interface{}{
	(*CommitteeCertificate)(nil), // 0: sonic.scc.cert.proto.CommitteeCertificate
	(*Member)(nil),               // 1: sonic.scc.cert.proto.Member
	(*AggregatedSignature)(nil),  // 2: sonic.scc.cert.proto.AggregatedSignature
}
var file_committee_certificate_proto_depIdxs = []int32{
	2, // 0: sonic.scc.cert.proto.CommitteeCertificate.signature:type_name -> sonic.scc.cert.proto.AggregatedSignature
	1, // 1: sonic.scc.cert.proto.CommitteeCertificate.members:type_name -> sonic.scc.cert.proto.Member
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_committee_certificate_proto_init() }
func file_committee_certificate_proto_init() {
	if File_committee_certificate_proto != nil {
		return
	}
	file_signature_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_committee_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitteeCertificate); i {
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
		file_committee_certificate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
			RawDescriptor: file_committee_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_committee_certificate_proto_goTypes,
		DependencyIndexes: file_committee_certificate_proto_depIdxs,
		MessageInfos:      file_committee_certificate_proto_msgTypes,
	}.Build()
	File_committee_certificate_proto = out.File
	file_committee_certificate_proto_rawDesc = nil
	file_committee_certificate_proto_goTypes = nil
	file_committee_certificate_proto_depIdxs = nil
}
