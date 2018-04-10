# dproto
[![GoDoc](https://godoc.org/github.com/linux4life798/dproto?status.png)](https://godoc.org/github.com/linux4life798/dproto)

# Description
This library allows you to use Protobufs (Google Protocol Buffers) without needing to
generate protobuf code and compile it into your project.
This means you don't need to compile .proto files.

This library is designed for marshaling and unmarshalling protobufs in a *dynamic* way.
It is designed to interface with any standard protobuf library and is tested against
the official protobuf C++ and Golang bindings, in addition to the C
[Nanopdb](https://github.com/nanopb/nanopb) library.

The intent of this library was to allow creating long running services that can
interpret and interface with clients using new/unknown protobuf messages.

The basic idea is that you construct a `ProtoFieldMap` that contains any protobuf
field number to protobuf type associations you are interested in and then you
call `DecodeBUffer` on the protobuf payload.
`DecodeBuffer` returns an array of `FieldValue`s that it decoded from the payload.
Each `FieldValue` specifies the protobuf field number and value decoded as a Golang
primitive(must be inside a `interface{}`).

# Status
Dproto supports all ProtoBuf primitive types. The following is a complete list:
* int32, int64
* uint32, uint64
* sint32, sint64
* fixed32, fixed64
* sfixed32, sfixed64
* float, double
* bool, string, bytes

# Name Explanation
Since we are marshalling and unmarshalling Protobuf messages in a dynamic way,
the project is called *dproto*.

# Examples

To give you a little taste of what dproto can do, check out the following
example of marshalling and unmarshalling a protobuf.

Say you want to interface with some other program using the following
protobuf description:
```protobuf
message LightStatus {
    bool  status    = 1;
    int64 intensity = 2;
}
```

The following example encodes a new message, prints out the byte stream,
and then immediately decodes it.


```go
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
```

This and other examples can be found in [example_test.go](example_test.go).