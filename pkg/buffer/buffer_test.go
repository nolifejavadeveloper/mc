package buffer

import (
	"errors"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestByte(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := byte(235)
	readWriteTest(func() { buf.WriteByte(data) }, func() (interface{}, error) { val, err := buf.ReadByte(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestBytes(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := []byte{2, 30, 9, 2, 6, 4, 3, 30, 20, 60, 99, 101, 100, 200, 182, 254, 255, 0}
	readWriteTest(func() { buf.Write(data) }, func() (interface{}, error) { val, err := buf.Read(len(data)); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestUInt16(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := uint16(50345)
	readWriteTest(func() { buf.WriteUInt16(data) }, func() (interface{}, error) { val, err := buf.ReadUInt16(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestUInt32(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := uint32(402938478)
	readWriteTest(func() { buf.WriteUInt32(data) }, func() (interface{}, error) { val, err := buf.ReadUInt32(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestUInt64(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := uint64(18039487574673633333)
	readWriteTest(func() { buf.WriteUInt64(data) }, func() (interface{}, error) { val, err := buf.ReadUInt64(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestInt8(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := int8(116)
	readWriteTest(func() { buf.WriteInt8(data) }, func() (interface{}, error) { val, err := buf.ReadInt8(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestInt16(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := int16(30295)
	readWriteTest(func() { buf.WriteInt16(data) }, func() (interface{}, error) { val, err := buf.ReadInt16(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestInt32(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := int32(2039485765)
	readWriteTest(func() { buf.WriteInt32(data) }, func() (interface{}, error) { val, err := buf.ReadInt32(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestInt64(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := int64(8418984723987492834)
	readWriteTest(func() { buf.WriteInt64(data) }, func() (interface{}, error) { val, err := buf.ReadInt64(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestVarInt(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := int32(2019283840)
	readWriteTest(func() { buf.WriteVarInt(data) }, func() (interface{}, error) { val, err := buf.ReadVarInt(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestVarLong(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	data := int64(8444038472213948273)
	readWriteTest(func() { buf.WriteVarLong(data) }, func() (interface{}, error) { val, err := buf.ReadVarLong(); return val, err }, data, nil, t)
	testBufferAdvance(buf, t)
}

func TestUUID(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	uid, _ := uuid.NewUUID()
	readWriteTest(func() { buf.WriteUUID(uid) }, func() (interface{}, error) { val, err := buf.ReadUUID(); return val, err }, uid, nil, t)
	testBufferAdvance(buf, t)
}

func TestVarIntOverflow(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	readWriteTest(func() { buf.WriteVarLong(3434343432132312312) }, func() (interface{}, error) { val, err := buf.ReadVarInt(); return val, err }, 0, errors.New("VarInt too large"), t)
}

func TestVarLongOverflow(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	readWriteTest(func() {
		buf.Write([]byte{
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
			0xFF, 0xFF, 0x7F,
		})
	}, func() (interface{}, error) { val, err := buf.ReadVarLong(); return val, err }, 0, errors.New("VarLong too large"), t)
}

func TestOutOfBounds(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0))
	readWriteTest(func() {}, func() (interface{}, error) { val, err := buf.ReadByte(); return val, err }, 12, errors.New("out of bounds"), t)
}

func readWriteTest(write func(), read func() (interface{}, error), expectedRes interface{}, expectedErr error, t *testing.T) {
	write()
	res, err := read()
	if expectedErr != nil {
		if !(err != nil && err.Error() == expectedErr.Error()) {
			t.Errorf("expecting error: %v, got error: %v", expectedErr, err)
		}

		return
	}

	if err != nil {
		t.Errorf("error occured: %v", err)
		return
	}

	if !reflect.DeepEqual(res, expectedRes) {
		t.Errorf("expecting %d, got %d", expectedRes, res)
	}
}

func testBufferAdvance(buf *Buffer, t *testing.T) {
	data := byte(29)
	readWriteTest(func() { buf.WriteByte(data) }, func() (interface{}, error) { val, err := buf.ReadByte(); return val, err }, data, nil, t)
}
