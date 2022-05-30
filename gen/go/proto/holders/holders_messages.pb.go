// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/holders/holders_messages.proto

package holders

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Customer definition
type Holder struct {
	Id        string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	FirstName string   `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	LastName  string   `protobuf:"bytes,3,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	Phone     string   `protobuf:"bytes,4,opt,name=phone" json:"phone,omitempty"`
	Email     string   `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	HeldBooks []string `protobuf:"bytes,6,rep,name=held_books,json=heldBooks" json:"held_books,omitempty"`
}

func (m *Holder) Reset()                    { *m = Holder{} }
func (m *Holder) String() string            { return proto.CompactTextString(m) }
func (*Holder) ProtoMessage()               {}
func (*Holder) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Holder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Holder) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Holder) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Holder) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Holder) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Holder) GetHeldBooks() []string {
	if m != nil {
		return m.HeldBooks
	}
	return nil
}

func init() {
	proto.RegisterType((*Holder)(nil), "tutorial.grpc.holders.v1.Holder")
}

func init() { proto.RegisterFile("proto/holders/holders_messages.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x86, 0xb1, 0x5d, 0xbb, 0x95, 0x86, 0x0e, 0xa2, 0x83, 0xa0, 0x14, 0x4c, 0xc9, 0x90, 0x49,
	0x21, 0xe4, 0x0d, 0x3c, 0x65, 0xca, 0x90, 0x31, 0x8b, 0x91, 0xa3, 0x8b, 0x2d, 0x22, 0x59, 0x46,
	0x52, 0xf2, 0x3a, 0x79, 0xd5, 0xa0, 0xb3, 0x3d, 0x1d, 0xff, 0xf7, 0x1d, 0xfc, 0x77, 0x74, 0x33,
	0x79, 0x17, 0xdd, 0x6e, 0x70, 0x46, 0x81, 0x0f, 0xeb, 0x6c, 0x2d, 0x84, 0x20, 0x7b, 0x08, 0x02,
	0x35, 0xe3, 0xf1, 0x11, 0x9d, 0xd7, 0xd2, 0x88, 0xde, 0x4f, 0x57, 0xb1, 0x6c, 0x89, 0xe7, 0xfe,
	0xff, 0x95, 0xd1, 0xea, 0x88, 0x91, 0x7d, 0xd3, 0x5c, 0x2b, 0x9e, 0xd5, 0xd9, 0x96, 0x9c, 0x73,
	0xad, 0xd8, 0x1f, 0xa5, 0x37, 0xed, 0x43, 0x6c, 0x47, 0x69, 0x81, 0xe7, 0xc8, 0x09, 0x92, 0x93,
	0xb4, 0xc0, 0x7e, 0x29, 0x31, 0x72, 0xb5, 0x05, 0xda, 0xaf, 0x04, 0x50, 0xfe, 0xd0, 0x72, 0x1a,
	0xdc, 0x08, 0xfc, 0x03, 0xc5, 0x1c, 0x12, 0x05, 0x2b, 0xb5, 0xe1, 0xe5, 0x4c, 0x31, 0xa4, 0x9e,
	0x01, 0x8c, 0x6a, 0x3b, 0xe7, 0xee, 0x81, 0x57, 0x75, 0x91, 0x7a, 0x12, 0x69, 0x12, 0x68, 0xc8,
	0xe5, 0x73, 0xb9, 0xb7, 0xab, 0xf0, 0x9b, 0xc3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x23, 0xc9,
	0x1b, 0xf5, 0x00, 0x00, 0x00,
}