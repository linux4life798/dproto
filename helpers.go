// Craig Hesling <craig@hesling.com>
// Started March 13, 2017
//
// This file holds helpful functions and constants for working with Protobuf
// in the dynamic setting. These are primarily for external use.

package dproto

import "github.com/golang/protobuf/protoc-gen-go/descriptor"
import "strconv"
import "encoding/base64"

// typeName2ProtoType maps the Protobuf type identifier to
// it's usable enumeration type.
var typeName2ProtoType = map[string]descriptor.FieldDescriptorProto_Type{
	"double":   descriptor.FieldDescriptorProto_TYPE_DOUBLE,
	"float":    descriptor.FieldDescriptorProto_TYPE_FLOAT,
	"int32":    descriptor.FieldDescriptorProto_TYPE_INT32,
	"int64":    descriptor.FieldDescriptorProto_TYPE_INT64,
	"uint32":   descriptor.FieldDescriptorProto_TYPE_UINT32,
	"uint64":   descriptor.FieldDescriptorProto_TYPE_UINT64,
	"sint32":   descriptor.FieldDescriptorProto_TYPE_SINT32,
	"sint64":   descriptor.FieldDescriptorProto_TYPE_SINT64,
	"fixed32":  descriptor.FieldDescriptorProto_TYPE_FIXED32,
	"fixed64":  descriptor.FieldDescriptorProto_TYPE_FIXED64,
	"sfixed32": descriptor.FieldDescriptorProto_TYPE_SFIXED32,
	"sfixed64": descriptor.FieldDescriptorProto_TYPE_SFIXED64,
	"bool":     descriptor.FieldDescriptorProto_TYPE_BOOL,
	"string":   descriptor.FieldDescriptorProto_TYPE_STRING,
	"bytes":    descriptor.FieldDescriptorProto_TYPE_BYTES,
}

// ParseProtobufType returns the Protobuf type enumeration of the
// given string
func ParseProtobufType(typ string) (descriptor.FieldDescriptorProto_Type, bool) {
	t, ok := typeName2ProtoType[typ]
	return t, ok
}

// ParseAs parses the string as the specified Protobuf type
func ParseAs(value string, pbtype descriptor.FieldDescriptorProto_Type, base int) (interface{}, error) {
	var v interface{}
	var err = ErrInvalidProtoBufType

	switch pbtype {

	case descriptor.FieldDescriptorProto_TYPE_INT32:
		fallthrough
	case descriptor.FieldDescriptorProto_TYPE_SINT32:
		fallthrough
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32:
		v, err = strconv.ParseInt(value, base, 32)
		v = int32(v.(int64))

	case descriptor.FieldDescriptorProto_TYPE_INT64:
		fallthrough
	case descriptor.FieldDescriptorProto_TYPE_SINT64:
		fallthrough
	case descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		v, err = strconv.ParseInt(value, base, 64)
		v = int64(v.(int64))

	case descriptor.FieldDescriptorProto_TYPE_UINT32:
		fallthrough
	case descriptor.FieldDescriptorProto_TYPE_FIXED32:
		v, err = strconv.ParseUint(value, base, 32)
		v = uint32(v.(uint64))

	case descriptor.FieldDescriptorProto_TYPE_UINT64:
		fallthrough
	case descriptor.FieldDescriptorProto_TYPE_FIXED64:
		v, err = strconv.ParseUint(value, base, 64)
		v = uint64(v.(uint64))

	case descriptor.FieldDescriptorProto_TYPE_BOOL:
		v, err = strconv.ParseBool(value)

	case descriptor.FieldDescriptorProto_TYPE_FLOAT:
		v, err = strconv.ParseFloat(value, 32)
		v = float32(v.(float64))

	case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		v, err = strconv.ParseFloat(value, 64)
		v = float64(v.(float64))

	case descriptor.FieldDescriptorProto_TYPE_STRING:
		v, err = value, nil

	case descriptor.FieldDescriptorProto_TYPE_BYTES:
		// we currently only support base64 encoding
		if base == 64 {
			v, err = base64.StdEncoding.DecodeString(value)
		}

	case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
		// Currently do not implement the Protobuf JSON format
		err = ErrInvalidProtoBufType
	}
	return v, err
}
