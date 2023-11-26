package splitter

type Splitter struct {
	buffer *Buffer
	skip   []Rule
	split  []Rule
	join   []Rule
}

func New(opts ...Option) *Splitter {
	s := &Splitter{
		buffer: &Buffer{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Splitter) Check(rules ...Rule) bool {
	for _, rule := range rules {
		if rule(s.buffer) {
			return true
		}
	}
	return false
}

func (s *Splitter) Process(str string) []string {
	s.buffer.incoming = []rune(str)
	for !s.buffer.isDone() {
		switch {
		case s.Check(s.skip...):
			s.buffer.trim()
		case s.Check(s.split...):
			s.buffer.store()
			fallthrough
		default:
			if s.Check(s.join...) {
				s.buffer.load()
			} else {
				s.buffer.trim()
			}
		}
	}
	if !s.buffer.isEmpty() {
		s.buffer.store()
	}
	return s.buffer.done
}
