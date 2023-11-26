package splitter

type Buffer struct {
	incoming []rune
	current  []rune
	done     []string
}

func (b *Buffer) IsEmpty() bool {
	return len(b.current) == 0
}

func (b *Buffer) IsDone() bool {
	return len(b.incoming) == 0
}

func (b *Buffer) Clear() {
	b.current = []rune{}
}

func (b *Buffer) Store() {
	if !b.IsEmpty() {
		b.done = append(b.done, string(b.current))
		b.Clear()
	}
}

func (b *Buffer) Trim() {
	if !b.IsDone() {
		b.incoming = b.incoming[1:]
	}
}

func (b *Buffer) Load() {
	if !b.IsDone() {
		b.current = append(b.current, b.incoming[0])
		b.Trim()
	}
}
