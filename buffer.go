package splitter

type Buffer struct {
	Incoming []rune
	Current  []rune
	Done     []string
}

func (b *Buffer) IsEmpty() bool {
	return len(b.Current) == 0
}

func (b *Buffer) IsDone() bool {
	return len(b.Incoming) == 0
}

func (b *Buffer) Clear() {
	b.Current = []rune{}
}

func (b *Buffer) Store() {
	if !b.IsEmpty() {
		b.Done = append(b.Done, string(b.Current))
		b.Clear()
	}
}

func (b *Buffer) Trim() {
	if !b.IsDone() {
		b.Incoming = b.Incoming[1:]
	}
}

func (b *Buffer) Load() {
	if !b.IsDone() {
		b.Current = append(b.Current, b.Incoming[0])
		b.Trim()
	}
}
