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
	s.Buffer.Incoming = []rune(str)
	for !s.Buffer.IsDone() {
		switch {
		case s.Check(s.Skip...):
			s.Buffer.Trim()
		case s.Check(s.Split...):
			s.Buffer.Store()
			fallthrough
		default:
			if s.Check(s.Join...) {
				s.Buffer.Load()
			} else {
				s.Buffer.Trim()
			}
		}
	}
	if !s.Buffer.IsEmpty() {
		s.Buffer.Store()
	}
	return s.Buffer.Done
}
