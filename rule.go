// Package splitter implements a string splitting package.
// It defines a type, [Splitter], with [Buffer] to store and change data,
// and [Rule] functions that control its behavior.
package splitter

// Rule function gains access to [Buffer] data using its public methods.
//
// Should return true if current [Buffer] state satisfies the rule and false otherwise.
type Rule func(*Buffer) bool
