// Craig Hesling <craig@hesling.com>
// Started December 15, 2016
//

package dproto

import "math"

// Notes Section
//
// # - Wire Field Types : Protobuf abstracted types
// 0 - Varint           : int32, int64, uint32, uint64, sint32, sint64, bool, enum
// 1 - Fixed64          : fixed64, sfixed64, double
// 2 - Length-delimited : string, byte, embedded messages, packet repeated fields
// 3 - Start group      : groups (depreciated)
// 4 - End group        : groups (depreciated)
// 5 - Fixed32          : fixed32, sfixed32, float
//
// Protobuf types int32 and int64 are pos/neg capable non-zigzag varints
// Protobuf types sint32 and sint64 are pos/neg capable zigzag encoded varints
// These types are the only ones that use the zig-zag encoding. After zig-zag
// encoding, they are then saved as a varint.

// WireType represents Protobuf's 3 bit wiretype identifier
// used in a marshalled packet.
type WireType uint8 // only uses 3 bits on wire

// WireVarint represents the Protobuf varint wire type
type WireVarint uint64

// AsInt32 returns the wiretype interpreted as a Protobuf int32
func (v WireVarint) AsInt32() int32 {
	return int32(v)
}

// AsInt64 returns the wiretype interpreted as a Protobuf int64
func (v WireVarint) AsInt64() int64 {
	return int64(v)
}

// AsUint32 returns the wiretype interpreted as a Protobuf uint32
func (v WireVarint) AsUint32() uint32 {
	return uint32(v)
}

// AsUint64 returns the wiretype interpreted as a Protobuf uint64
func (v WireVarint) AsUint64() uint64 {
	return uint64(v)
}

// AsSint32 returns the wiretype interpreted as a Protobuf sint32
func (v WireVarint) AsSint32() int32 {
	// Computation taken from Protobuf's decode.go
	return int32(uint64((uint32(v) >> 1) ^ uint32((int32(uint64(v)&1)<<31)>>31)))
}

// AsSint64 returns the wiretype interpreted as a Protobuf sint64
func (v WireVarint) AsSint64() int64 {
	// Computation taken from Protobuf's decode.go
	return int64(uint64((uint64(v) >> 1) ^ uint64((int64(uint64(v)&1)<<63)>>63)))
}

// AsBool returns the wiretype interpreted as a Protobuf bool
func (v WireVarint) AsBool() bool {
	//TODO: Confirm correctness
	return uint64(v) != 0
}

// AsEnum returns the wiretype interpreted as a Protobuf enum (just a uint)
func (v WireVarint) AsEnum() uint64 {
	//TODO: Confirm correctness
	return uint64(v)
}

// FromInt32 sets the wiretype from a Protobuf int32
func (v *WireVarint) FromInt32(i int32) {
	*v = WireVarint(uint32(i))
}

// FromInt64 sets the wiretype from a Protobuf int64
func (v *WireVarint) FromInt64(i int64) {
	*v = WireVarint(i)
}

// FromUint32 sets the wiretype from a Protobuf uint32
func (v *WireVarint) FromUint32(i uint32) {
	*v = WireVarint(i)
}

// FromUint64 sets the wiretype from a Protobuf uint64
func (v *WireVarint) FromUint64(i uint64) {
	*v = WireVarint(i)
}

// FromSint32 sets the wiretype from a Protobuf sint32
func (v *WireVarint) FromSint32(i int32) {
	//TODO: Confirm correctness
	// Taken from protobuff web page on encoding
	*v = WireVarint((i << 1) ^ (i >> 31))
}

// FromSint64 sets the wiretype from a Protobuf sint64
func (v *WireVarint) FromSint64(i int64) {
	//TODO: Confirm correctness
	// Taken from protobuff web page on encoding
	*v = WireVarint((i << 1) ^ (i >> 63))
}

// FromBool sets the wiretype from a Protobuf bool
func (v *WireVarint) FromBool(i bool) {
	//TODO: Confirm correctness
	if i {
		*v = WireVarint(1)
	} else {
		*v = WireVarint(0)
	}
}

// FromEnum sets the wiretype from a Protobuf enum
func (v *WireVarint) FromEnum(i uint64) {
	//TODO: Confirm correctness
	*v = WireVarint(i)
}

// WireFixed32 represents the Protobuf fixed32 wire type
type WireFixed32 uint64 // DecodeFixed32 gives uint64

// AsFixed32 returns the wiretype interpreted as a Protobuf fixed32
func (v WireFixed32) AsFixed32() uint32 {
	return uint32(v)
}

// AsSfixed32 returns the wiretype interpreted as a Protobuf sfixed32
func (v WireFixed32) AsSfixed32() int32 {
	// This should be sufficient, since any saved negative value should have
	// the 63rd bit set, which would look like a very large uint64.
	// This value is larger than what an int64 can hold, so it lets it
	// overflow and wrap around modulo. This works in the twos-complement
	// system to recreate the negative saved value.
	return int32(v)
}

// AsFloat returns the wiretype interpreted as a Protobuf float
func (v WireFixed32) AsFloat() float32 {
	return math.Float32frombits(uint32(v))
}

// FromFixed32 sets the wiretype from a Protobuf fixed32
func (v *WireFixed32) FromFixed32(i uint32) {
	*v = WireFixed32(i)
}

// FromSfixed32 sets the wiretype from a Protobuf fixed32
func (v *WireFixed32) FromSfixed32(i int32) {
	*v = WireFixed32(uint32(i))
}

// FromFloat sets the wiretype from a Protobuf float
func (v *WireFixed32) FromFloat(i float32) {
	*v = WireFixed32(math.Float32bits(i))
}

// WireFixed64 represents the Protobuf fixed64 wire type
type WireFixed64 uint64

// AsFixed64 returns the wiretype interpreted as a Protobuf fixed64
func (v WireFixed64) AsFixed64() uint64 {
	return uint64(v)
}

// AsSfixed64 returns the wiretype interpreted as a Protobuf sfixed64
func (v WireFixed64) AsSfixed64() int64 {
	// This should be sufficient, since any saved negative value should have
	// the 63rd bit set, which would look like a very large uint64.
	// This value is larger than what an int64 can hold, so it lets it
	// overflow and wrap around modulo. This works in the twos-complement
	// system to recreate the negative saved value.
	return int64(v)
}

// AsDouble returns the wiretype interpreted as a Protobuf double
func (v WireFixed64) AsDouble() float64 {
	return math.Float64frombits(uint64(v))
}

// FromFixed64 sets the wiretype from a Protobuf fixed64
func (v *WireFixed64) FromFixed64(i uint64) {
	*v = WireFixed64(i)
}

// FromSfixed64 sets the wiretype from a Protobuf fixed64
func (v *WireFixed64) FromSfixed64(i int64) {
	*v = WireFixed64(uint64(i))
}

// FromDouble sets the wiretype from a Protobuf double
func (v *WireFixed64) FromDouble(i float64) {
	*v = WireFixed64(math.Float64bits(i))
}
