package splitter

type Option func(*Splitter)

func AddSkip(rules ...Rule) Option {
	return func(s *Splitter) {
		s.skip = append(s.skip, rules...)
	}
}

func AddSplit(rules ...Rule) Option {
	return func(s *Splitter) {
		s.split = append(s.split, rules...)
	}
}

func AddJoin(rules ...Rule) Option {
	return func(s *Splitter) {
		s.join = append(s.join, rules...)
	}
}
