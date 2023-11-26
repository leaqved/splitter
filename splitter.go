package splitter

type Splitter struct {
	Buffer *Buffer
	Skip   []Rule
	Split  []Rule
	Join   []Rule
}

func New(opts ...Option) *Splitter {
	s := &Splitter{
		Buffer: &Buffer{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Splitter) Check(rules ...Rule) bool {
	for _, rule := range rules {
		if rule(s.Buffer) {
			return true
		}
	}
	return false
}

func (s *Splitter) Process(str string) []string {
	s.Buffer.incoming = []rune(str)
	for !s.Buffer.isDone() {
		switch {
		case s.Check(s.Skip...):
			s.Buffer.trim()
		case s.Check(s.Split...):
			s.Buffer.store()
			fallthrough
		default:
			if s.Check(s.Join...) {
				s.Buffer.load()
			} else {
				s.Buffer.trim()
			}
		}
	}
	if !s.Buffer.isEmpty() {
		s.Buffer.store()
	}
	return s.Buffer.done
}
