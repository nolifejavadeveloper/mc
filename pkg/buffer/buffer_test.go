package buffer

import (
	"errors"
	"testing"
)

func TestBuffer(t *testing.T) {
	buf := MakeBuffer(make([]byte, 0, 2048))

	readWriteTest(func() { buf.WriteByte(12) }, func() (int64, error) { val, err := buf.ReadByte(); return int64(val), err }, 12, nil, "byte", t)
	readWriteTest(func() { buf.WriteUInt16(13) }, func() (int64, error) { val, err := buf.ReadUInt16(); return int64(val), err }, 13, nil, "uint16", t)
	readWriteTest(func() { buf.WriteUInt32(14) }, func() (int64, error) { val, err := buf.ReadUInt32(); return int64(val), err }, 14, nil, "uint32", t)
	readWriteTest(func() { buf.WriteUInt64(15) }, func() (int64, error) { val, err := buf.ReadUInt64(); return int64(val), err }, 15, nil, "uint64", t)

	readWriteTest(func() { buf.WriteInt8(22) }, func() (int64, error) { val, err := buf.ReadInt8(); return int64(val), err }, 22, nil, "int8", t)
	readWriteTest(func() { buf.WriteInt16(23) }, func() (int64, error) { val, err := buf.ReadInt16(); return int64(val), err }, 23, nil, "int16", t)
	readWriteTest(func() { buf.WriteInt32(24) }, func() (int64, error) { val, err := buf.ReadInt32(); return int64(val), err }, 24, nil, "int32", t)
	readWriteTest(func() { buf.WriteInt64(25) }, func() (int64, error) { val, err := buf.ReadInt64(); return int64(val), err }, 25, nil, "int64", t)

	readWriteTest(func() { buf.WriteVarInt(38374745) }, func() (int64, error) { val, err := buf.ReadVarInt(); return int64(val), err }, 38374745, nil, "varint", t)
	readWriteTest(func() { buf.WriteVarLong(97878676878) }, func() (int64, error) { val, err := buf.ReadVarLong(); return int64(val), err }, 97878676878, nil, "varlong", t)
}

func readWriteTest(write func(), read func() (int64, error), expectedRes int64, expectedErr error, name string, t *testing.T) {
	write()
	res, err := read()
	if expectedErr != nil {
		if err != nil && errors.Is(err, expectedErr) {
			t.Logf("%s expecting error: %v, got error: %v", name, expectedErr, err)
		} else {
			t.Errorf("%s expecting error: %v, got error: %v", name, expectedErr, err)
		}

		return
	}

	if res != expectedRes {
		t.Errorf("%s expecting %d, got %d", name, expectedRes, res)
	} else {
		t.Logf("%s expecting %d, got %d", name, expectedRes, res)
	}
}
