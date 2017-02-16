// Craig Hesling <craig@hesling.com>
// Started December 15, 2016
//
// This file houses the highest level interface for dproto. Using these methods
// you will register associations

// Package dproto allows for marshalling and unmarshalling of Protobuf
// messages in a dynamic manor. This means you can interface with new
// Protobuf libraries or clients during runtime, without compiling.
//
// Dproto currently implements the basic abstraction layer between declared
// Protobuf message field/types and the field/wiretypes. These associations
// are expected from the user before marshalling/unmarshalling. It is up to
// the user how to store and translate these associations to the dproto library.
//
// Note that we say some construct is on the "wire" or has a "wiretype", when
// it refers the bits in a marshalled buffer. See the following link for more
// information on wire types:
// https://developers.google.com/protocol-buffers/docs/encoding
package dproto

import (
	"fmt"

	"os"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// FieldNum represents Protobuf's field numbers
type FieldNum uint32

// ProtoFieldMap associates field numbers with it's high-level Protobuf type.
type ProtoFieldMap struct {
	field2type map[FieldNum]descriptor.FieldDescriptorProto_Type
	type2field map[descriptor.FieldDescriptorProto_Type]FieldNum
}

// NewProtoFieldMap create a new ProtoFieldMap object.
func NewProtoFieldMap() *ProtoFieldMap {
	var m = new(ProtoFieldMap)
	m.Reset()
	return m
}

// Reset clears the stored associations inside a ProtoFieldMap
func (m *ProtoFieldMap) Reset() {
	m.field2type = make(map[FieldNum]descriptor.FieldDescriptorProto_Type)
	m.type2field = make(map[descriptor.FieldDescriptorProto_Type]FieldNum)
}

// Add adds a Field-Type association to a ProtoFieldMap
func (m *ProtoFieldMap) Add(field FieldNum, typ descriptor.FieldDescriptorProto_Type) {
	m.field2type[field] = typ
	m.type2field[typ] = field

	if len(m.field2type) != len(m.type2field) {
		fmt.Fprintf(os.Stderr, "Error - ProtoTypeMap is inconsistent")
	}
}

// RemoveByField removes the Field-Type association from a ProtoFieldMap
// that has the specified field number.
// It returns true if the association was found and removed, false otherwise
func (m *ProtoFieldMap) RemoveByField(field FieldNum) (ok bool) {
	if typ, ok := m.field2type[field]; ok {
		delete(m.field2type, field)
		delete(m.type2field, typ)
	}
	return
}

// RemoveByType removes the Field-Type association from a ProtoFieldMap
// that has the specified type.
// It returns true if the association was found and removed, false otherwise
func (m *ProtoFieldMap) RemoveByType(typ descriptor.FieldDescriptorProto_Type) (ok bool) {
	if field, ok := m.type2field[typ]; ok {
		delete(m.type2field, typ)
		delete(m.field2type, field)
	}
	return
}

// Print shows the ProtoFieldMap to the user for debugging purposes.
func (m *ProtoFieldMap) Print() {
	fmt.Println(m)
}

var protoType2WireType = map[descriptor.FieldDescriptorProto_Type]WireType{
	descriptor.FieldDescriptorProto_TYPE_DOUBLE:  proto.WireFixed64,
	descriptor.FieldDescriptorProto_TYPE_FLOAT:   proto.WireFixed32,
	descriptor.FieldDescriptorProto_TYPE_INT64:   proto.WireVarint,
	descriptor.FieldDescriptorProto_TYPE_UINT64:  proto.WireVarint,
	descriptor.FieldDescriptorProto_TYPE_INT32:   proto.WireVarint,
	descriptor.FieldDescriptorProto_TYPE_UINT32:  proto.WireVarint,
	descriptor.FieldDescriptorProto_TYPE_FIXED64: proto.WireFixed64,
	descriptor.FieldDescriptorProto_TYPE_FIXED32: proto.WireFixed32,
	descriptor.FieldDescriptorProto_TYPE_BOOL:    proto.WireVarint,
	descriptor.FieldDescriptorProto_TYPE_STRING:  proto.WireBytes,
	// descriptor.FieldDescriptorProto_TYPE_GROUP: proto.WireStartGroup
	descriptor.FieldDescriptorProto_TYPE_MESSAGE:  proto.WireBytes,
	descriptor.FieldDescriptorProto_TYPE_ENUM:     proto.WireVarint,
	descriptor.FieldDescriptorProto_TYPE_SFIXED32: proto.WireFixed32,
	descriptor.FieldDescriptorProto_TYPE_SFIXED64: proto.WireFixed64,
	descriptor.FieldDescriptorProto_TYPE_SINT32:   proto.WireVarint,
	descriptor.FieldDescriptorProto_TYPE_SINT64:   proto.WireVarint,
}

// Unmarshal will unmarshal a byte array into a WireMessage
func Unmarshal(buf []byte) (*WireMessage, error) {
	m := NewWireMessage()
	if err := m.Unmarshal(buf); err != nil {
		return nil, err
	}
	return m, nil
}
