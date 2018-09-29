package dproto

import (
	"io/ioutil"
	"testing"
)

/* The following fields were set in a nanopb C program */

// message.myint32 = 32423;
// message.myint64 = -98327;
// message.myuint32 = 1;
// message.myuint64 = 962329;
// message.mysint32 = -231;
// message.mysint64 = -3932764127;
// message.mybool = true;
// This Enum field is in nthere, but not yet parsible
// message.myenum = TestEnum_THIRD;

// message.myfixed64 = 342647260612;
// message.mysfixed64 = -324;
// message.mydouble = 3.1456;

// message.myfixed32 = 445545;
// message.mysfixed32 = -30423;
// message.myfloat = 3.227799;

type Answer interface{}

var ans = []Answer{
	int32(32423),
	int64(-98327),
	uint32(1),
	uint64(962329),
	int32(-231),
	int64(-3932764127),
	bool(true),
	//
	uint64(342647260612),
	int64(-324),
	float64(3.1456),
	//
	uint32(445545),
	int32(-30423),
	float32(3.227799),
}

// protobufBinary The reference file generated from nanopb
const protobufBinary = "testprotobuf.bin"

func testAgainstTestMessage(t *testing.T, m *WireMessage) {
	// Print out decoded values
	myint32, ok := m.DecodeInt32(1)
	if !ok {
		t.Error("Failed to find myint32")
	}
	t.Logf("myint32 = %d\n", myint32)
	if myint32 != ans[0] {
		t.Error("myint32 did not match expected value")
	}

	myint64, ok := m.DecodeInt64(2)
	if !ok {
		t.Error("Failed to find myint64")
	}
	t.Logf("myint64 = %d\n", myint64)
	if myint64 != ans[1] {
		t.Error("myint64 did not match expected value")
	}

	myuint32, ok := m.DecodeUint32(3)
	if !ok {
		t.Error("Failed to find myuint32")
	}
	t.Logf("myuint32 = %d\n", myuint32)
	if myuint32 != ans[2] {
		t.Error("myuint32 did not match expected value")
	}

	myuint64, ok := m.DecodeUint64(4)
	if !ok {
		t.Error("Failed to find myuint64")
	}
	t.Logf("myuint64 = %d\n", myuint64)
	if myuint64 != ans[3] {
		t.Error("myuint64 did not match expected value")
	}

	mysint32, ok := m.DecodeSint32(5)
	if !ok {
		t.Error("Failed to find mysint32")
	}
	t.Logf("mysint32 = %d\n", mysint32)
	if mysint32 != ans[4] {
		t.Error("mysint32 did not match expected value")
	}

	mysint64, ok := m.DecodeSint64(6)
	if !ok {
		t.Error("Failed to find mysint64")
	}
	t.Logf("mysint64 = %d\n", mysint64)
	if mysint64 != ans[5] {
		t.Error("mysint64 did not match expected value")
	}

	mybool, ok := m.DecodeBool(7)
	if !ok {
		t.Error("Failed to find mybool")
	}
	t.Logf("mybool = %t\n", mybool)
	if mybool != ans[6] {
		t.Error("mybool did not match expected value")
	}

	// myenum, ok := m.DecodeEnum(8)
	// if !ok {
	// 	t.Error("Failed to find myenum")
	// }
	// t.Logf("myenum = %c\n", myenum)

	myfixed64, ok := m.DecodeFixed64(9)
	if !ok {
		t.Error("Failed to find myfixed64")
	}
	t.Logf("myfixed64 = %d\n", myfixed64)
	if myfixed64 != ans[7] {
		t.Error("myfixed64 did not match expected value")
	}

	mysfixed64, ok := m.DecodeSfixed64(10)
	if !ok {
		t.Error("Failed to find mysfixed64")
	}
	t.Logf("mysfixed64 = %d\n", mysfixed64)
	if mysfixed64 != ans[8] {
		t.Error("mysfixed64 did not match expected value")
	}

	mydouble, ok := m.DecodeDouble(11)
	if !ok {
		t.Error("Failed to find mydouble")
	}
	t.Logf("mydouble = %f\n", mydouble)
	if mydouble != ans[9] {
		t.Error("mydouble did not match expected value")
	}

	myfixed32, ok := m.DecodeFixed32(12)
	if !ok {
		t.Error("Failed to find myfixed32")
	}
	t.Logf("myfixed32 = %d\n", myfixed32)
	if myfixed32 != ans[10] {
		t.Error("myfixed32 did not match expected value")
	}

	mysfixed32, ok := m.DecodeSfixed32(13)
	if !ok {
		t.Error("Failed to find mysfixed32")
	}
	t.Logf("mysfixed32 = %d\n", mysfixed32)
	if mysfixed32 != ans[11] {
		t.Error("mysfixed32 did not match expected value")
	}

	myfloat, ok := m.DecodeFloat(14)
	if !ok {
		t.Error("Failed to find myfloat")
	}
	t.Logf("myfloat = %f\n", myfloat)
	if myfloat != ans[12] {
		t.Error("myfloat did not match expected value")
	}
}

// go test -v
func TestUnmarshal(t *testing.T) {
	// Read in file to buffer
	buf, err := ioutil.ReadFile(protobufBinary)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("Read " + protobufBinary)

	// Get unmarshalling under way
	m, err := Unmarshal(buf)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("Unmarshalled the buffer from file")

	testAgainstTestMessage(t, m)
}

// TestMarshal1 reads the reference protobuf binary, Unmarshals it,
// then re-Marshals it, and then Unmarshals and verifies the values
// The unused Enum field should be transfered along for the ride
func TestMarshal1(t *testing.T) {
	// Read in file to buffer
	buf, err := ioutil.ReadFile(protobufBinary)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("Read " + protobufBinary)

	// Get unmarshalling under way
	m, err := Unmarshal(buf)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("Unmarshalled the buffer from file")

	bytes, err := m.Marshal()
	if err != nil {
		t.Error("Error Marshaling: " + err.Error())
	}
	t.Log("Marshalled the buffer")

	// Save so the user can actually inspect the bits
	err = ioutil.WriteFile("TestMarshal1.bin", bytes, 0644)
	if err != nil {
		t.Error("Error writing TestMarshal1.bin:", err.Error())
	}

	// Unmarshal the already marshalled bytes
	m, err = Unmarshal(bytes)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("Unmarshalled the buffer that we just unmarshalled")

	testAgainstTestMessage(t, m)
}

// TestMarshal2 builds the reference protobuf message manually, Marshals
// the bytes, then Unmarshals the bytes and verifies the data matches
func TestMarshal2(t *testing.T) {
	m := NewWireMessage()

	// message.myint32 = 32423;
	m.EncodeInt32(1, 32423)
	// message.myint64 = -98327;
	m.EncodeInt64(2, -98327)
	// message.myuint32 = 1;
	m.EncodeUint32(3, 1)
	// message.myuint64 = 962329;
	m.EncodeUint64(4, 962329)
	// message.mysint32 = -231;
	m.EncodeSint32(5, -231)
	// message.mysint64 = -3932764127;
	m.EncodeSint64(6, -3932764127)
	// message.mybool = true;
	m.EncodeBool(7, true)
	// message.myenum = TestEnum_THIRD;
	// m.EncodeEnum(8, 3)
	// BIG NOTE: This does not yet incluse the enum, so two bytes won't match

	// message.myfixed64 = 342647260612;
	m.EncodeFixed64(9, 342647260612)
	// message.mysfixed64 = -324;
	m.EncodeSfixed64(10, -324)
	// message.mydouble = 3.1456;
	m.EncodeDouble(11, 3.1456)

	// message.myfixed32 = 445545;
	m.EncodeFixed32(12, 445545)
	// message.mysfixed32 = -30423;
	m.EncodeSfixed32(13, -30423)
	// message.myfloat = 3.227799;
	m.EncodeFloat(14, 3.227799)

	bytes, err := m.Marshal()
	if err != nil {
		t.Error("Error Marshaling: " + err.Error())
	}
	t.Log("Marshalled the buffer")

	// Save so the user can actually inspect the bits
	err = ioutil.WriteFile("TestMarshal2.bin", bytes, 0644)
	if err != nil {
		t.Error("Error writing TestMarshal2.bin:", err.Error())
	}

	// Unmarshal the already marshalled bytes
	m, err = Unmarshal(bytes)
	if err != nil {
		t.Error("Error Unmarshaling: " + err.Error())
	}
	t.Log("Unmarshalled the buffer that we just unmarshalled")

	testAgainstTestMessage(t, m)
}

func BenchmarkWireMessageMarshal(b *testing.B) {
	m := NewWireMessage()

	// message.myint32 = 32423;
	m.EncodeInt32(1, 32423)
	// message.myint64 = -98327;
	m.EncodeInt64(2, -98327)
	// message.myuint32 = 1;
	m.EncodeUint32(3, 1)
	// message.myuint64 = 962329;
	m.EncodeUint64(4, 962329)
	// message.mysint32 = -231;
	m.EncodeSint32(5, -231)
	// message.mysint64 = -3932764127;
	m.EncodeSint64(6, -3932764127)
	// message.mybool = true;
	m.EncodeBool(7, true)
	// message.myenum = TestEnum_THIRD;
	// m.EncodeEnum(8, 3)
	// BIG NOTE: This does not yet incluse the enum, so two bytes won't match

	// message.myfixed64 = 342647260612;
	m.EncodeFixed64(9, 342647260612)
	// message.mysfixed64 = -324;
	m.EncodeSfixed64(10, -324)
	// message.mydouble = 3.1456;
	m.EncodeDouble(11, 3.1456)

	// message.myfixed32 = 445545;
	m.EncodeFixed32(12, 445545)
	// message.mysfixed32 = -30423;
	m.EncodeSfixed32(13, -30423)
	// message.myfloat = 3.227799;
	m.EncodeFloat(14, 3.227799)

	for i := 0; i < b.N; i++ {
		_, err := m.Marshal()
		if err != nil {
			b.Error("Error Marshaling: " + err.Error())
		}
	}
}

func BenchmarkWireMessageUnmarshal(b *testing.B) {
	m := NewWireMessage()

	// message.myint32 = 32423;
	m.EncodeInt32(1, 32423)
	// message.myint64 = -98327;
	m.EncodeInt64(2, -98327)
	// message.myuint32 = 1;
	m.EncodeUint32(3, 1)
	// message.myuint64 = 962329;
	m.EncodeUint64(4, 962329)
	// message.mysint32 = -231;
	m.EncodeSint32(5, -231)
	// message.mysint64 = -3932764127;
	m.EncodeSint64(6, -3932764127)
	// message.mybool = true;
	m.EncodeBool(7, true)
	// message.myenum = TestEnum_THIRD;
	// m.EncodeEnum(8, 3)
	// BIG NOTE: This does not yet incluse the enum, so two bytes won't match

	// message.myfixed64 = 342647260612;
	m.EncodeFixed64(9, 342647260612)
	// message.mysfixed64 = -324;
	m.EncodeSfixed64(10, -324)
	// message.mydouble = 3.1456;
	m.EncodeDouble(11, 3.1456)

	// message.myfixed32 = 445545;
	m.EncodeFixed32(12, 445545)
	// message.mysfixed32 = -30423;
	m.EncodeSfixed32(13, -30423)
	// message.myfloat = 3.227799;
	m.EncodeFloat(14, 3.227799)

	bytes, err := m.Marshal()
	if err != nil {
		b.Error("Error Marshaling: " + err.Error())
	}

	for i := 0; i < b.N; i++ {
		// Unmarshal the already marshalled bytes
		if _, err := Unmarshal(bytes); err != nil {
			b.Error("Error Unmarshaling: " + err.Error())
		}
	}
}
