package splitter

type Option func(*Splitter)

func AddSkip(rules ...Rule) Option {
	return func(s *Splitter) {
		s.Skip = append(s.Skip, rules...)
	}
}

func AddSplit(rules ...Rule) Option {
	return func(s *Splitter) {
		s.Split = append(s.Split, rules...)
	}
}
