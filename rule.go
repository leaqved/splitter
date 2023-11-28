// Package splitter implements a string splitting package.
// It defines a type, [Splitter], with [Buffer] to store and change data,
// and [Rule] functions that control its behavior.
package splitter

// Rule function gains access to [Buffer] data using its public methods.
//
// Should return true if current [Buffer] state satisfies the rule and false otherwise.
type Rule func(*Buffer) bool

type CharRule func(char rune) bool

func Before(rule CharRule) Rule {
	return func(b *Buffer) bool {
		char, ok := b.GetRight()
		if !ok {
			return false
		}
		return rule(char)
	}
}

func After(rule CharRule) Rule {
	return func(b *Buffer) bool {
		char, ok := b.GetLeft()
		if !ok {
			return false
		}
		return rule(char)
	}
}

func (r Rule) And(another Rule) Rule {
	return func(b *Buffer) bool {
		return r(b) && another(b)
	}
}
