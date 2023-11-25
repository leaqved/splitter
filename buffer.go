package splitter

type Buffer struct {
	Incoming []rune
	Current  []rune
	Done     []string
}

func (b *Buffer) IsEmpty() bool {
	return len(b.Current) == 0
}
