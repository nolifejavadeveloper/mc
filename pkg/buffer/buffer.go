package network

import "errors"

type Buffer struct {
	data    []byte
	pointer int
}

func MakeBuffer1(data []byte) *Buffer {
	return &Buffer{
		data:    data,
		pointer: 0,
	}
}

func (b *Buffer) advancePointer(n int) {
	b.pointer += n
}

func (b *Buffer) Read(n int) ([]byte, error) {
	if b.pointer+n < b.Size() {
		return nil, errors.New("out of bounds")
	}

	data := b.data[b.pointer : b.pointer+n]

	b.advancePointer(n)

	return data, nil
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.pointer < b.Size() {
		return 0, errors.New("out of bounds")
	}

	data := b.data[b.pointer]

	b.advancePointer(1)
	return data, nil
}

func (b *Buffer) ReadUInt8() (uint8, error) {
	val, err := b.ReadByte()
	return uint8(val), err
}

func (b *Buffer) ReadUInt16() (uint16, error) {
	bytes, err := b.Read(2)
	if err != nil {
		return 0, err
	}

	return uint16(bytes[0])<<8 | uint16(bytes[1]), nil
}

func (b *Buffer) ReadInt16() (int16, error) {
	val, err := b.ReadUInt16()
	return int16(val), err
}

func (b *Buffer) ReadUInt32() (uint32, error) {
	bytes, err := b.Read(4)
	if err != nil {
		return 0, err
	}

	return uint32(bytes[0])<<24 | uint32(bytes[1])<<16 | uint32(bytes[2])<<8 | uint32(bytes[3]), nil
}

func (b *Buffer) ReadInt32() (int32, error) {
	val, err := b.ReadUInt32()
	return int32(val), err
}

func (b *Buffer) ReadUInt64() (uint64, error) {
	bytes, err := b.Read(4)
	if err != nil {
		return 0, err
	}

	return uint64(bytes[0])<<56 | uint64(bytes[1])<<48 | uint64(bytes[2])<<40 | uint64(bytes[3])<<32 | uint64(bytes[4])<<24 | uint64(bytes[5])<<16 | uint64(bytes[6])<<8 | uint64(bytes[7]), nil
}

func (b *Buffer) ReadInt64() (int64, error) {
	val, err := b.ReadUInt64()
	return int64(val), err
}

func (b *Buffer) ReadVarInt() (int32, error) {
	var pos int = 0

	var val int32 = 0
	for {
		byt, err := b.ReadByte()
		if err != nil {
			return 0, err
		}

		val |= int32(byt&0x7F) << pos
		pos += 7
		if pos >= 32 {
			// varint too large
			return 0, errors.New("VarInt too large")
		}

		if byt&0x80 == 0 {
			break
		}
	}

	return val, nil
}

func (b *Buffer) Size() int {
	return len(b.data)
}
