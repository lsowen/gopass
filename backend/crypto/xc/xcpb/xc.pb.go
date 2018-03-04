// Code generated by protoc-gen-go.
// source: xc.proto
// DO NOT EDIT!

/*
Package xcpb is a generated protocol buffer package.

It is generated from these files:
	xc.proto

It has these top-level messages:
	Header
	Chunk
	Message
	Identity
	PublicKey
	PrivateKey
	Secring
	Pubring
*/
package xcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PublicKeyAlgorithm int32

const (
	PublicKeyAlgorithm_UNKNOWN PublicKeyAlgorithm = 0
	PublicKeyAlgorithm_NACL    PublicKeyAlgorithm = 1
)

var PublicKeyAlgorithm_name = map[int32]string{
	0: "UNKNOWN",
	1: "NACL",
}
var PublicKeyAlgorithm_value = map[string]int32{
	"UNKNOWN": 0,
	"NACL":    1,
}

func (x PublicKeyAlgorithm) String() string {
	return proto.EnumName(PublicKeyAlgorithm_name, int32(x))
}
func (PublicKeyAlgorithm) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Header struct {
	Sender     string            `protobuf:"bytes,1,opt,name=Sender" json:"Sender,omitempty"`
	Recipients map[string][]byte `protobuf:"bytes,3,rep,name=Recipients" json:"Recipients,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metadata   map[string]string `protobuf:"bytes,4,rep,name=Metadata" json:"Metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Header) GetRecipients() map[string][]byte {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *Header) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Chunk struct {
	Body []byte `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Chunk) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Message struct {
	Version    uint32   `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	Header     *Header  `protobuf:"bytes,2,opt,name=Header" json:"Header,omitempty"`
	Chunks     []*Chunk `protobuf:"bytes,3,rep,name=Chunks" json:"Chunks,omitempty"`
	Compressed bool     `protobuf:"varint,4,opt,name=Compressed" json:"Compressed,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Message) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Message) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Message) GetChunks() []*Chunk {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *Message) GetCompressed() bool {
	if m != nil {
		return m.Compressed
	}
	return false
}

type Identity struct {
	Name    string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	Comment string `protobuf:"bytes,2,opt,name=Comment" json:"Comment,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=Email" json:"Email,omitempty"`
}

func (m *Identity) Reset()                    { *m = Identity{} }
func (m *Identity) String() string            { return proto.CompactTextString(m) }
func (*Identity) ProtoMessage()               {}
func (*Identity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Identity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Identity) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *Identity) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type PublicKey struct {
	PubKeyAlgo   PublicKeyAlgorithm `protobuf:"varint,1,opt,name=PubKeyAlgo,enum=xcpb.PublicKeyAlgorithm" json:"PubKeyAlgo,omitempty"`
	CreationTime uint64             `protobuf:"varint,2,opt,name=CreationTime" json:"CreationTime,omitempty"`
	PublicKey    []byte             `protobuf:"bytes,3,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	Identity     *Identity          `protobuf:"bytes,4,opt,name=Identity" json:"Identity,omitempty"`
	Fingerprint  string             `protobuf:"bytes,5,opt,name=Fingerprint" json:"Fingerprint,omitempty"`
}

func (m *PublicKey) Reset()                    { *m = PublicKey{} }
func (m *PublicKey) String() string            { return proto.CompactTextString(m) }
func (*PublicKey) ProtoMessage()               {}
func (*PublicKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PublicKey) GetPubKeyAlgo() PublicKeyAlgorithm {
	if m != nil {
		return m.PubKeyAlgo
	}
	return PublicKeyAlgorithm_UNKNOWN
}

func (m *PublicKey) GetCreationTime() uint64 {
	if m != nil {
		return m.CreationTime
	}
	return 0
}

func (m *PublicKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PublicKey) GetIdentity() *Identity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *PublicKey) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

type PrivateKey struct {
	PublicKey  *PublicKey `protobuf:"bytes,1,opt,name=PublicKey" json:"PublicKey,omitempty"`
	Ciphertext []byte     `protobuf:"bytes,2,opt,name=Ciphertext,proto3" json:"Ciphertext,omitempty"`
	Nonce      []byte     `protobuf:"bytes,3,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	Salt       []byte     `protobuf:"bytes,4,opt,name=Salt,proto3" json:"Salt,omitempty"`
}

func (m *PrivateKey) Reset()                    { *m = PrivateKey{} }
func (m *PrivateKey) String() string            { return proto.CompactTextString(m) }
func (*PrivateKey) ProtoMessage()               {}
func (*PrivateKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PrivateKey) GetPublicKey() *PublicKey {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PrivateKey) GetCiphertext() []byte {
	if m != nil {
		return m.Ciphertext
	}
	return nil
}

func (m *PrivateKey) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *PrivateKey) GetSalt() []byte {
	if m != nil {
		return m.Salt
	}
	return nil
}

type Secring struct {
	Version     uint32        `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	PrivateKeys []*PrivateKey `protobuf:"bytes,2,rep,name=PrivateKeys" json:"PrivateKeys,omitempty"`
}

func (m *Secring) Reset()                    { *m = Secring{} }
func (m *Secring) String() string            { return proto.CompactTextString(m) }
func (*Secring) ProtoMessage()               {}
func (*Secring) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Secring) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Secring) GetPrivateKeys() []*PrivateKey {
	if m != nil {
		return m.PrivateKeys
	}
	return nil
}

type Pubring struct {
	Version    uint32       `protobuf:"varint,1,opt,name=Version" json:"Version,omitempty"`
	PublicKeys []*PublicKey `protobuf:"bytes,2,rep,name=PublicKeys" json:"PublicKeys,omitempty"`
}

func (m *Pubring) Reset()                    { *m = Pubring{} }
func (m *Pubring) String() string            { return proto.CompactTextString(m) }
func (*Pubring) ProtoMessage()               {}
func (*Pubring) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Pubring) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Pubring) GetPublicKeys() []*PublicKey {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "xcpb.Header")
	proto.RegisterType((*Chunk)(nil), "xcpb.Chunk")
	proto.RegisterType((*Message)(nil), "xcpb.Message")
	proto.RegisterType((*Identity)(nil), "xcpb.Identity")
	proto.RegisterType((*PublicKey)(nil), "xcpb.PublicKey")
	proto.RegisterType((*PrivateKey)(nil), "xcpb.PrivateKey")
	proto.RegisterType((*Secring)(nil), "xcpb.Secring")
	proto.RegisterType((*Pubring)(nil), "xcpb.Pubring")
	proto.RegisterEnum("xcpb.PublicKeyAlgorithm", PublicKeyAlgorithm_name, PublicKeyAlgorithm_value)
}

func init() { proto.RegisterFile("xc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xad, 0x6c, 0xc5, 0x1f, 0x23, 0x25, 0x31, 0x4b, 0x29, 0xc2, 0x0d, 0xc5, 0xa8, 0x3d, 0x98,
	0x94, 0xba, 0xe0, 0x42, 0x09, 0xfd, 0x38, 0xa4, 0x26, 0xa5, 0xc5, 0x8d, 0x6a, 0xd6, 0x69, 0x73,
	0x96, 0xa5, 0xc1, 0x5e, 0x62, 0xad, 0xc4, 0x6a, 0x1d, 0xec, 0x1f, 0xd0, 0x6b, 0x0f, 0xfd, 0x6f,
	0xfd, 0x3f, 0x45, 0xbb, 0x92, 0xbd, 0x4e, 0xa0, 0xb9, 0xed, 0x7c, 0xbc, 0x37, 0xef, 0xcd, 0x08,
	0x41, 0x6b, 0x1d, 0x0d, 0x32, 0x91, 0xca, 0x94, 0xd8, 0xeb, 0x28, 0x9b, 0xf9, 0x7f, 0x6a, 0xd0,
	0xf8, 0x82, 0x61, 0x8c, 0x82, 0x3c, 0x81, 0xc6, 0x14, 0x79, 0x8c, 0xc2, 0xb3, 0x7a, 0x56, 0xbf,
	0x4d, 0xcb, 0x88, 0x7c, 0x00, 0xa0, 0x18, 0xb1, 0x8c, 0x21, 0x97, 0xb9, 0x57, 0xef, 0xd5, 0xfb,
	0xce, 0xf0, 0x64, 0x50, 0xa0, 0x07, 0x1a, 0x39, 0xd8, 0x95, 0x2f, 0xb8, 0x14, 0x1b, 0x6a, 0xf4,
	0x93, 0xb7, 0xd0, 0xba, 0x44, 0x19, 0xc6, 0xa1, 0x0c, 0x3d, 0x5b, 0x61, 0xbb, 0x7b, 0xd8, 0xaa,
	0xa8, 0x91, 0xdb, 0xde, 0xee, 0x47, 0x38, 0xbe, 0x43, 0x4b, 0x3a, 0x50, 0xbf, 0xc1, 0x4d, 0xa9,
	0xae, 0x78, 0x92, 0xc7, 0x70, 0x70, 0x1b, 0x2e, 0x57, 0xe8, 0xd5, 0x7a, 0x56, 0xdf, 0xa5, 0x3a,
	0x78, 0x57, 0x3b, 0xb3, 0xba, 0xef, 0xe1, 0x70, 0x8f, 0xf9, 0x21, 0x70, 0xdb, 0x00, 0xfb, 0x4f,
	0xe1, 0x60, 0xb4, 0x58, 0xf1, 0x1b, 0x42, 0xc0, 0xfe, 0x94, 0xc6, 0x1a, 0xe5, 0x52, 0xf5, 0xf6,
	0x7f, 0x5b, 0xd0, 0xbc, 0xc4, 0x3c, 0x0f, 0xe7, 0x48, 0x3c, 0x68, 0xfe, 0x44, 0x91, 0xb3, 0x94,
	0xab, 0x96, 0x43, 0x5a, 0x85, 0xe4, 0x45, 0xb5, 0x56, 0xc5, 0xee, 0x0c, 0x5d, 0xd3, 0x34, 0xad,
	0x56, 0xfe, 0x1c, 0x1a, 0x6a, 0x50, 0xb5, 0x56, 0x47, 0x77, 0xa9, 0x1c, 0x2d, 0x4b, 0xe4, 0x19,
	0xc0, 0x28, 0x4d, 0x32, 0x81, 0x79, 0x8e, 0xb1, 0x67, 0xf7, 0xac, 0x7e, 0x8b, 0x1a, 0x19, 0x3f,
	0x80, 0xd6, 0xd7, 0x18, 0xb9, 0x64, 0x72, 0x53, 0x08, 0x0e, 0xc2, 0x04, 0x4b, 0x9b, 0xea, 0x5d,
	0x88, 0x1c, 0xa5, 0x49, 0x82, 0x5c, 0x96, 0x4e, 0xab, 0xb0, 0xd8, 0xc0, 0x45, 0x12, 0xb2, 0xa5,
	0x57, 0xd7, 0x1b, 0x50, 0x81, 0xff, 0xd7, 0x82, 0xf6, 0x64, 0x35, 0x5b, 0xb2, 0x68, 0x8c, 0x1b,
	0x72, 0x06, 0x30, 0x59, 0xcd, 0xc6, 0xb8, 0x39, 0x5f, 0xce, 0x53, 0xc5, 0x7b, 0x34, 0xf4, 0xb4,
	0xcc, 0x6d, 0x53, 0x51, 0x12, 0x4c, 0x2e, 0x12, 0x6a, 0xf4, 0x12, 0x1f, 0xdc, 0x91, 0xc0, 0x50,
	0xb2, 0x94, 0x5f, 0xb1, 0x44, 0xaf, 0xd9, 0xa6, 0x7b, 0x39, 0x72, 0x62, 0x8c, 0x52, 0x2a, 0x5c,
	0x6a, 0xcc, 0x3e, 0xdd, 0x39, 0x53, 0xbe, 0x9d, 0xe1, 0x91, 0x9e, 0x5c, 0x65, 0xe9, 0xce, 0x79,
	0x0f, 0x9c, 0xcf, 0x8c, 0xcf, 0x51, 0x64, 0x82, 0x71, 0xe9, 0x1d, 0x28, 0x47, 0x66, 0xca, 0xff,
	0x65, 0x01, 0x4c, 0x04, 0xbb, 0x0d, 0x25, 0x16, 0xe4, 0xaf, 0xcc, 0xd1, 0x96, 0x62, 0x3f, 0xbe,
	0xe3, 0xcb, 0xd4, 0x52, 0x5c, 0x81, 0x65, 0x0b, 0x14, 0x12, 0xd7, 0xb2, 0xfc, 0xde, 0x8c, 0x4c,
	0xb1, 0xcb, 0x20, 0xe5, 0x11, 0x96, 0x2e, 0x74, 0x50, 0xdc, 0x63, 0x1a, 0x2e, 0xa5, 0x52, 0xef,
	0x52, 0xf5, 0xf6, 0xaf, 0xa1, 0x39, 0xc5, 0x48, 0x30, 0x3e, 0xff, 0xcf, 0xf7, 0x33, 0x04, 0x67,
	0xa7, 0x35, 0xf7, 0x6a, 0xea, 0xf3, 0xe8, 0x94, 0xfa, 0xb6, 0x05, 0x6a, 0x36, 0xf9, 0x57, 0xd0,
	0x9c, 0xac, 0x66, 0x0f, 0x10, 0xbf, 0x56, 0xf7, 0xd4, 0xa6, 0x2a, 0xde, 0x7b, 0xbe, 0x8d, 0x96,
	0xd3, 0x97, 0x40, 0xee, 0x1f, 0x9a, 0x38, 0xd0, 0xfc, 0x11, 0x8c, 0x83, 0xef, 0xd7, 0x41, 0xe7,
	0x11, 0x69, 0x81, 0x1d, 0x9c, 0x8f, 0xbe, 0x75, 0xac, 0x59, 0x43, 0xfd, 0x5b, 0xde, 0xfc, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0xf4, 0x7d, 0xf6, 0xc9, 0x67, 0x04, 0x00, 0x00,
}