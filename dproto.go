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
}

// NewProtoFieldMap create a new ProtoFieldMap object.
func NewProtoFieldMap() *ProtoFieldMap {
	var fm = new(ProtoFieldMap)
	fm.Reset()
	return fm
}

// Reset clears the stored associations inside a ProtoFieldMap
func (fm *ProtoFieldMap) Reset() {
	fm.field2type = make(map[FieldNum]descriptor.FieldDescriptorProto_Type)
}

// Add adds a Field-Type association to a ProtoFieldMap
func (fm *ProtoFieldMap) Add(field FieldNum, typ descriptor.FieldDescriptorProto_Type) (ok bool) {
	if _, ok := protoType2WireType[typ]; ok {
		fm.field2type[field] = typ
	}
	return
}

// RemoveByField removes the Field-Type association from a ProtoFieldMap
// that has the specified field number.
// It returns true if the association was found and removed, false otherwise
func (fm *ProtoFieldMap) RemoveByField(field FieldNum) (ok bool) {
	if _, ok := fm.field2type[field]; ok {
		delete(fm.field2type, field)
	}
	return
}

// RemoveByType removes all field Field-Type association from a ProtoFieldMap
// that has the specified type. This will check all map entries.
// It returns true if an association was found and removed, false otherwise
func (fm *ProtoFieldMap) RemoveByType(typ descriptor.FieldDescriptorProto_Type) bool {
	deleteList := make([]FieldNum, 0, len(fm.field2type))
	for k, v := range fm.field2type {
		if v == typ {
			deleteList = append(deleteList, k)
		}
	}

	if len(deleteList) == 0 {
		return false
	}

	for _, k := range deleteList {
		delete(fm.field2type, k)
	}
	return true
}

// Print shows the ProtoFieldMap to the user for debugging purposes.
func (fm *ProtoFieldMap) Print() {
	fmt.Println(fm)
}
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
