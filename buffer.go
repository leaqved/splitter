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