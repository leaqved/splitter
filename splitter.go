// Package splitter implements a string splitting package.
// It defines a type, [Splitter], with [Buffer] to store and change data,
// and [Rule] functions that control its behavior.
package splitter

// A Splitter represents an active object that
// splits strings according to [Rule] functions.
type Splitter struct {
	buffer *Buffer
	skip   []Rule
	split  []Rule
	join   []Rule
}

// New creates a new [Splitter] and applies [Option] functions to it.
func New(opts ...Option) *Splitter {
	s := &Splitter{
		buffer: &Buffer{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// check applies [Rule] functions to the [Splitter]'s buffer.
//
// Returns true if any of [Rule] functions returns true and false otherwise.
func (s *Splitter) check(rules ...Rule) bool {
	for _, rule := range rules {
		if rule(s.buffer) {
			return true
		}
	}
	return false
}

// Split splits the string according to [Rule] functions of the [Splitter].
//
// Returns the resulting slice.
func (s *Splitter) Split(str string) []string {
	s.buffer.incoming = []rune(str)
	for !s.buffer.isDone() {
		switch {
		case s.check(s.skip...):
			s.buffer.trim()
			s.buffer.store()
		case s.check(s.split...):
			s.buffer.store()
			fallthrough
		default:
			if s.check(s.join...) {
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
