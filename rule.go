// Package splitter implements a string splitting package.
// It defines a type, [Splitter], with [Buffer] to store and change data,
// and [Rule] functions that control its behavior.
package splitter

// Rule function gains access to [Buffer] data using its public methods.
//
// Should return true if current [Buffer] state satisfies the rule and false otherwise.
type Rule func(*Buffer) bool

// CharRule function should return true if given character satisfies the rule and false otherwise.
type CharRule func(char rune) bool

// Before returns a Rule that checks if the next character incoming in [Buffer] satisfies the CharRule.
func Before(rule CharRule) Rule {
	return func(b *Buffer) bool {
		char, ok := b.GetRight()
		if !ok {
			return false
		}
		return rule(char)
	}
}

// After returns a Rule that checks if the last character in [Buffer] satisfies the CharRule.
func After(rule CharRule) Rule {
	return func(b *Buffer) bool {
		char, ok := b.GetLeft()
		if !ok {
			return false
		}
		return rule(char)
	}
}

// And returns a logical conjunction of two Rules.
func (r Rule) And(another Rule) Rule {
	return func(b *Buffer) bool {
		return r(b) && another(b)
	}
}

// And returns a logical conjunction of two CharRules.
func (r CharRule) And(another CharRule) CharRule {
	return func(char rune) bool {
		return r(char) && another(char)
	}
}
