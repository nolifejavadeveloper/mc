package network

type Buffer struct {
	data    []byte
	pointer int
}

func makeBuffer(data []byte) *Buffer {
	return &Buffer{
		data: data,
		pointer: 0,
	}
}

func (b *Buffer) advancePointer(n int) {
	b.pointer += n
}

func (b *Buffer) ReadByte() byte {
	data := b.data[b.pointer]

	b.advancePointer(1)

	return data
}

func (b *Buffer) ReadBytes(n int) []byte {
	data := b.data[b.pointer : b.pointer + n]

	b.advancePointer(n)

	return data
}

func (b *Buffer) ReadVarInt() int32 {
	int pos = 0
	int32 val = 0;
	for {
		val |= (byt & 0x7F) << pos
		pos += 7
		if (pos >= 32) {
			// varint too large
			//TODO: add error handling 
		}

		if byt & 0x80 == 0 {
			break
		}
	}

	return val;
}

