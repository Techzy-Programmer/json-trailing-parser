package tokenizer

import "strings"

// Tokenization rules for JSON Trailing Parser
// Escape sequence is not valid inside array

type Node struct {
	IsArray  bool
	Accessor string
	GetAll   bool
}

type Tokenizer struct {
	query      string
	buffer     string
	arrayMode  bool
	escapeMode bool
	history    []tokenAction
	nodes      []Node
}

func NewTokenizer(q string) *Tokenizer {
	return &Tokenizer{
		query:   strings.TrimSpace(q),
		history: make([]tokenAction, 0),
	}
}
