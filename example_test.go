package dproto_test

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/linux4life798/dproto"
)

func ExampleNewWireMessage() {
	m := dproto.NewWireMessage()

	// Add the following bool and int64
	m.EncodeBool(1, true)
	m.EncodeInt64(2, 10)

	bytes, err := m.Marshal()
	if err != nil {
		panic("Error Marshaling: " + err.Error())
	}
	fmt.Printf("ProtobufBinary: [ %# x ]\n", bytes)
	// Output: ProtobufBinary: [ 0x08 0x01 0x10 0x0a ]
}

func ExampleWireMessage_Marshal() {
	m := dproto.NewWireMessage()

	// Add the following bool and int64
	m.EncodeBool(1, true)
	m.EncodeInt64(2, 10)

	bytes, err := m.Marshal()
	if err != nil {
		panic("Error Marshaling: " + err.Error())
	}
	fmt.Printf("ProtobufBinary: [ %# x ]\n", bytes)
	// Output: ProtobufBinary: [ 0x08 0x01 0x10 0x0a ]
}

func ExampleWireMessage_Unmarshal() {
	m := dproto.NewWireMessage()

	// Add the following bool and int64
	m.EncodeBool(1, true)
	m.EncodeInt64(2, 10)

	bytes, err := m.Marshal()
	if err != nil {
		panic("Error Marshaling: " + err.Error())
	}
	fmt.Printf("ProtobufBinary: [ %# x ]\n", bytes)

	// Unmarshal the already marshalled bytes
	m, err = dproto.Unmarshal(bytes)
	if err != nil {
		panic("Error Unmarshaling: " + err.Error())
	}

	if bVal, ok := m.DecodeBool(1); ok {
		fmt.Println(bVal)
	} else {
		fmt.Println("No bool field 1")
	}

	if iVal, ok := m.DecodeInt64(2); ok {
		fmt.Println(iVal)
	} else {
		fmt.Println("No int64 field 2")
	}
	// Output:
	// ProtobufBinary: [ 0x08 0x01 0x10 0x0a ]
	// true
	// 10
}

// This example shows how to abstractly Marshal and Unmarshal
// protobuf messages.
// It should be noted that the ProtoFieldMap class already implements
// this behavior for you.
func ExampleWireMessage_Unmarshal_abstract() {

	types := []descriptor.FieldDescriptorProto_Type{
		descriptor.FieldDescriptorProto_TYPE_BOOL,
		descriptor.FieldDescriptorProto_TYPE_INT64,
	}
	values := []interface{}{
		bool(true),
		int64(10),
	}

	m := dproto.NewWireMessage()

	// Add the following bool and int64
	for index, value := range values {
		err := m.EncodeAs(dproto.FieldNum(index+1), value, types[index])
		if err != nil {
			panic(err)
		}
	}

	// Marshal the message
	bytes, err := m.Marshal()
	if err != nil {
		panic("Error Marshaling: " + err.Error())
	}
	fmt.Printf("ProtobufBinary: [ %# x ]\n", bytes)

	// Unmarshal the already marshalled bytes
	m, err = dproto.Unmarshal(bytes)
	if err != nil {
		panic("Error Unmarshaling: " + err.Error())
	}

	// Decode each field and print
	for index, typ := range types {
		val, err := m.DecodeAs(dproto.FieldNum(index+1), typ)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(val)
		}
	}
	// Output:
	// ProtobufBinary: [ 0x08 0x01 0x10 0x0a ]
	// true
	// 10
}

// Encodes and Decodes a protobuf buffer that would interface with
// a .proto file message containing "bool status =1;" and
// "int64 intensity = 2;".
func ExampleNewProtoFieldMap() {
	// Setup the ProtoFieldMap to interface with the
	// following .proto file:
	//     message LightStatus {
	//         bool  status    = 1;
	//         int64 intensity = 2;
	//     }

	// Setup a FieldMap that holds the type-fieldnumber association
	// This is effectively the information held in a .proto file
	fm := dproto.NewProtoFieldMap()

	// Add Protobuf bool as field 1 ("bool status = 1;")
	if !fm.Add(1, descriptor.FieldDescriptorProto_TYPE_BOOL) {
		panic("Failed to add bool field 1")
	}
	// Add Protobuf int64 as field 2 ("int64 intensity = 2;")
	if !fm.Add(2, descriptor.FieldDescriptorProto_TYPE_INT64) {
		panic("Failed to add bool field 1")
	}

	// Provide some values for our "status" and "intensity"
	values := []dproto.FieldValue{
		dproto.FieldValue{
			Field: 1, // status field number
			Value: bool(true),
		},
		dproto.FieldValue{
			Field: 2, // intensity field number
			Value: int64(10),
		},
	}

	// Encode out values into the protobuf message described in fm
	bytes, err := fm.EncodeBuffer(values)
	if err != nil {
		panic("Error Encoding: " + err.Error())
	}
	fmt.Printf("ProtobufBinary: [ %# x ]\n", bytes)

	// Decode all protobuf fields
	decodedValues, err := fm.DecodeBuffer(bytes)
	if err != nil {
		panic("Error Decoding: " + err.Error())
	}
	for _, val := range decodedValues {
		fmt.Printf("%v: %v\n", val.Field, val.Value)
	}
	// Unordered output:
	// ProtobufBinary: [ 0x08 0x01 0x10 0x0a ]
	// 1: true
	// 2: 10
}
