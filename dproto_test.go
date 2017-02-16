package dproto

import (
	"io/ioutil"
	"testing"
)

// message.myint32 = 32423;
// message.myint64 = -98327;
// message.myuint32 = 1;
// message.myuint64 = 962329;
// message.mysint32 = -231;
// message.mysint64 = -3932764127;
// message.mybool = true;
// message.myenum = TestEnum_THIRD;

// message.myfixed64 = 342647260612;
// message.mysfixed64 = -324;
// message.mydouble = 3.1456;

// message.myfixed32 = 445545;
// message.mysfixed32 = -30423;
// message.myfloat = 3.227799;
const protobufBinary = "testprotobuf.bin"

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
	t.Log("Unmarshalled the buffer")

	// Print out decoded values
	myint32, ok := m.DecodeInt32(1)
	if ok == false {
		t.Error("Failed to find myint32")
	}
	t.Logf("myint32 = %d\n", myint32)

	myint64, ok := m.DecodeInt64(2)
	if ok == false {
		t.Error("Failed to find myint64")
	}
	t.Logf("myint64 = %d\n", myint64)

	myuint32, ok := m.DecodeUint32(3)
	if ok == false {
		t.Error("Failed to find myuint32")
	}
	t.Logf("myuint32 = %d\n", myuint32)

	myuint64, ok := m.DecodeUint64(4)
	if ok == false {
		t.Error("Failed to find myuint64")
	}
	t.Logf("myuint64 = %d\n", myuint64)

	mysint32, ok := m.DecodeSint32(5)
	if ok == false {
		t.Error("Failed to find mysint32")
	}
	t.Logf("mysint32 = %d\n", mysint32)

	mysint64, ok := m.DecodeSint64(6)
	if ok == false {
		t.Error("Failed to find mysint64")
	}
	t.Logf("mysint64 = %d\n", mysint64)

	mybool, ok := m.DecodeBool(7)
	if ok == false {
		t.Error("Failed to find mybool")
	}
	t.Logf("mybool = %t\n", mybool)

	// myenum, ok := m.DecodeEnum(8)
	// if ok == false {
	// 	t.Error("Failed to find myenum")
	// }
	// t.Logf("myenum = %c\n", myenum)

	myfixed64, ok := m.DecodeFixed64(9)
	if ok == false {
		t.Error("Failed to find myfixed64")
	}
	t.Logf("myfixed64 = %d\n", myfixed64)

	mysfixed64, ok := m.DecodeSfixed64(10)
	if ok == false {
		t.Error("Failed to find mysfixed64")
	}
	t.Logf("mysfixed64 = %d\n", mysfixed64)

	mydouble, ok := m.DecodeDouble(11)
	if ok == false {
		t.Error("Failed to find mydouble")
	}
	t.Logf("mydouble = %f\n", mydouble)

	myfixed32, ok := m.DecodeFixed32(12)
	if ok == false {
		t.Error("Failed to find myfixed32")
	}
	t.Logf("myfixed32 = %d\n", myfixed32)

	mysfixed32, ok := m.DecodeSfixed32(13)
	if ok == false {
		t.Error("Failed to find mysfixed32")
	}
	t.Logf("mysfixed32 = %d\n", mysfixed32)

	myfloat, ok := m.DecodeFloat(14)
	if ok == false {
		t.Error("Failed to find myfloat")
	}
	t.Logf("myfloat = %f\n", myfloat)
}
