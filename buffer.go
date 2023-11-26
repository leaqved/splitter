package splitter

type Buffer struct {
	incoming []rune
	current  []rune
	done     []string
}

func (b *Buffer) isEmpty() bool {
	return len(b.current) == 0
}

func (b *Buffer) isDone() bool {
	return len(b.incoming) == 0
}

func (b *Buffer) clear() {
	b.current = []rune{}
}

func (b *Buffer) store() {
	if !b.isEmpty() {
		b.done = append(b.done, string(b.current))
		b.clear()
	}
}

func (b *Buffer) trim() {
	if !b.isDone() {
		b.incoming = b.incoming[1:]
	}
}

func (b *Buffer) load() {
	if !b.isDone() {
		b.current = append(b.current, b.incoming[0])
		b.trim()
	}
}

func (b *Buffer) GetAll() ([]rune, []rune, []string) {
	return b.incoming, b.current, b.done
}

func (b *Buffer) GetIncoming() []rune {
	return b.incoming
}

func (b *Buffer) GetCurrent() []rune {
	return b.current
}

func (b *Buffer) GetDone() []string {
	return b.done
}

func (b *Buffer) GetLeft() (rune, bool) {
	if b.isEmpty() {
		return 0, false
	}
	return b.current[len(b.current)-1], true
}

func (b *Buffer) GetRight() (rune, bool) {
	if b.isDone() {
		return 0, false
	}
	return b.incoming[0], true
}
