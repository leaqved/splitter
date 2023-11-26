package splitter

type Option func(*Splitter)

func AddSkip(rules ...Rule) Option {
	return func(s *Splitter) {
		s.Skip = append(s.Skip, rules...)
	}
}
