package circular

import "errors"

const testVersion = 4

// Buffer type
type Buffer struct {
	size       int
	startPos   int
	endPos     int
	overwrited int
	store      []byte
}

// NewBuffer initializes new buffer
func NewBuffer(size int) *Buffer {
	b := new(Buffer)
	b.size = size
	b.startPos = -1
	b.endPos = -1
	b.overwrited = -1
	b.store = make([]byte, size, size)
	return b
}

// ReadByte deletes oldest byte from buffer
func (b *Buffer) ReadByte() (byte, error) {
	if b.startPos > -1 {
		element := b.store[b.startPos]
		if b.startPos == b.endPos {
			if b.overwrited > 0 {
				if b.startPos < b.size-1 {
					b.startPos++
					b.endPos++
				} else {
					b.startPos = 0
					b.endPos = 0
				}
				b.overwrited--
			} else {
				b.startPos = -1
				b.endPos = -1
			}
		} else {
			if b.startPos < b.size-1 {
				b.startPos++
			} else {
				b.startPos = 0
			}
		}

		return element, nil
	}
	return 0, errors.New("No elements")
}

// WriteByte writes new byte to buffer
func (b *Buffer) WriteByte(c byte) error {
	if b.endPos == -1 {
		b.store[0] = c
		b.endPos = 0
		b.startPos = 0
		return nil
	}
	if b.endPos < b.size-1 {
		b.endPos++
	} else {
		b.endPos = 0
	}
	if b.endPos != b.startPos {
		b.store[b.endPos] = c
		return nil
	}

	return errors.New("Full buffer")
}

// Overwrite overwrites oldest byte if buffer is full
func (b *Buffer) Overwrite(c byte) {
	if b.startPos == b.endPos {
		b.WriteByte(c)
		b.overwrited++
		return
	}
	b.store[b.startPos] = c
	if b.startPos < b.size-1 {
		b.startPos++
		if b.endPos < b.size-1 {
			b.endPos++
		} else {
			b.endPos--
		}
	} else {
		b.startPos = 0
	}
	b.overwrited++
}

// Reset clears buffer
func (b *Buffer) Reset() {
	b.endPos = -1
	b.startPos = -1
}
