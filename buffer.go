package splitter

type Buffer struct {
	Incoming []rune
	Current  []rune
	Done     []string
}
