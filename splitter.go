package splitter

type Splitter struct {
	Buffer *Buffer
	Skip   []Rule
	Split  []Rule
	Join   []Rule
}

func New() *Splitter {
	return &Splitter{
		Buffer: &Buffer{},
	}
}

func (s *Splitter) Check(rules ...Rule) bool {
	for _, rule := range rules {
		if rule(s.Buffer) {
			return true
		}
	}
	return false
}
