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

// Field represents Protobuf's field numbers
type FieldNum uint32

// High Level Combiner Interface Goes Here

// Unmarshal will unmarshal a byte array into a WireMessage
func Unmarshal(buf []byte) (*WireMessage, error) {
	m := NewWireMessage()
	if err := m.Unmarshal(buf); err != nil {
		return nil, err
	}
	return m, nil
}
