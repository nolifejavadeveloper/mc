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

func (b *Buffer) readByte() byte {
	data := b.data[b.pointer]

	b.advancePointer(1);

	return data
}

func (b *Buffer) readBytes(n int) []byte {
	data := b.data[b.pointer : b.pointer + n]

	b.advancePointer(n);

	return data
}