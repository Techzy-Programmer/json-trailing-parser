package tokenizer

const (
	objectAccessor = '.'
	arrayStart     = '['
	arrayEnd       = ']'
	escapeToken    = ':'
	wildcard       = '*'
	templateStart  = '{'
	templateEnd    = '}'
)

var escapables = []rune{
	objectAccessor,
	arrayStart,
	arrayEnd,
	escapeToken,
	templateStart,
	templateEnd,
}

type tokenAction int

const (
	started tokenAction = iota
	bufferAdded
	arrayStarted
	arrayEnded
	objectAccessed
	escapeStarted
	escapeEnded
	templateStarted
	templateEnded
	ended
)
