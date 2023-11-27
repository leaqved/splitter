// Package splitter implements a string splitting package.
// It defines a type, [Splitter], with [Buffer] to store and change data,
// and [Rule] functions that control its behavior.
package splitter

// Option function is applied in constructor function of the [Splitter].
type Option func(*Splitter)

// AddSkip returns an [Option] that adds [Rule] functions to the skip section of the [Splitter].
func AddSkip(rules ...Rule) Option {
	return func(s *Splitter) {
		s.skip = append(s.skip, rules...)
	}
}

// AddSplit returns an [Option] that adds [Rule] functions to the split section of the [Splitter].
func AddSplit(rules ...Rule) Option {
	return func(s *Splitter) {
		s.split = append(s.split, rules...)
	}
}

// AddJoin returns an [Option] that adds [Rule] functions to the join section of the [Splitter].
func AddJoin(rules ...Rule) Option {
	return func(s *Splitter) {
		s.join = append(s.join, rules...)
	}
}
