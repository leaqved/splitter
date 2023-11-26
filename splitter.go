package splitter

type Splitter struct {
	Buffer *Buffer
	Skip   []Rule
	Split  []Rule
	Join   []Rule
}
