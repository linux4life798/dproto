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

