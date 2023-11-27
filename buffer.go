// Package splitter implements a string splitting package.
// It defines a type, [Splitter], with [Buffer] to store and change data,
// and [Rule] functions that control its behavior.
package splitter

// Buffer represents a temporary storage for [Splitter].
type Buffer struct {
	incoming []rune
	current  []rune
	done     []string
}

// isEmpty returns true if the current slice is empty.
func (b *Buffer) isEmpty() bool {
	return len(b.current) == 0
}

// isDone returns true if the incoming slice is empty.
func (b *Buffer) isDone() bool {
	return len(b.incoming) == 0
}

// clear clears the [Buffer].
func (b *Buffer) clear() {
	b.current = []rune{}
}

// store moves the current slice to the done slice.
func (b *Buffer) store() {
	if !b.isEmpty() {
		b.done = append(b.done, string(b.current))
		b.clear()
	}
}

// trim trims the first element of the incoming slice.
func (b *Buffer) trim() {
	if !b.isDone() {
		b.incoming = b.incoming[1:]
	}
}

// load moves the first element of the incoming slice to the current slice.
func (b *Buffer) load() {
	if !b.isDone() {
		b.current = append(b.current, b.incoming[0])
		b.trim()
	}
}

// GetAll returns incoming, current and done slices.
func (b *Buffer) GetAll() ([]rune, []rune, []string) {
	return b.incoming, b.current, b.done
}

// GetIncoming returns incoming slice.
func (b *Buffer) GetIncoming() []rune {
	return b.incoming
}

// GetCurrent returns current slice.
func (b *Buffer) GetCurrent() []rune {
	return b.current
}

// GetDone returns done slice.
func (b *Buffer) GetDone() []string {
	return b.done
}

// GetLeft returns the last element of the current slice.
//
// Returns false if the current slice is empty.
func (b *Buffer) GetLeft() (rune, bool) {
	if b.isEmpty() {
		return 0, false
	}
	return b.current[len(b.current)-1], true
}

// GetRight returns the first element of the incoming slice.
//
// Returns false if the incoming slice is empty.
func (b *Buffer) GetRight() (rune, bool) {
	if b.isDone() {
		return 0, false
	}
	return b.incoming[0], true
}
