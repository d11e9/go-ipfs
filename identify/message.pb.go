// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package identify is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Identify
*/
package identify

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Identify struct {
	Id               []byte `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Pubkey           []byte `protobuf:"bytes,2,req,name=pubkey" json:"pubkey,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Identify) Reset()         { *m = Identify{} }
func (m *Identify) String() string { return proto.CompactTextString(m) }
func (*Identify) ProtoMessage()    {}

func (m *Identify) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Identify) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func init() {
}