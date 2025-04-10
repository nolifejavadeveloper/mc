package network

import "errors"

type Buffer struct {
	data    []byte
	pointer int
}

func makeBuffer(data []byte) *Buffer {
	return &Buffer{
		data:    data,
		pointer: 0,
	}
}

func (b *Buffer) advancePointer(n int) {
	b.pointer += n
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.pointer < b.Size() {
		return 0, errors.New("out of bounds")
	}

	data := b.data[b.pointer]

	b.advancePointer(1)

	return data, nil
}

func (b *Buffer) ReadBytes(n int) ([]byte, error) {
	if b.pointer+n < b.Size() {
		return nil, errors.New("out of bounds")
	}

	data := b.data[b.pointer : b.pointer+n]

	b.advancePointer(n)

	return data, nil
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
