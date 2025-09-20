package tokenizer

import (
	"strings"

	"github.com/Techzy-Programmer/json-trailing-parser/jtparser"
)

func (t *Tokenizer) Tokenize() (*[]Node, error) {
	t.logAction(Started)
	t.nodes = make([]Node, 0)

	if strings.HasPrefix(t.query, string(ArrayStart)) {
		return nil, &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: "expecting root to be of type object found array at position 0",
		}
	}

	for i, r := range t.query {
		str := string(r)

		if avErr := t.validateArrayMode(str, i); avErr != nil {
			return nil, avErr
		}

		if t.escapeMode {
			// ToDo: check if this rune is even escapable?

			t.buffer += str
			t.escapeMode = false
			t.logAction(EscapeEnded)

			continue
		}

		switch r {
		case EscapeChar:
			t.escapeMode = true
			t.logAction(EscapeStarted)

		case ArrayStart:
			if asErr := t.validateArrayModeState(true, i); asErr != nil {
				return nil, asErr
			}

			t.arrayMode = true
			t.flushBuffer()
			t.logAction(ArrayStarted)

		case ArrayEnd:
			if asErr := t.validateArrayModeState(false, i); asErr != nil {
				return nil, asErr
			}

			t.flushBuffer()
			t.arrayMode = false
			t.logAction(ArrayEnded)

		case ObjectAccessor:
			if oaErr := t.validateObjectAccessor(i); oaErr != nil {
				return nil, oaErr
			}

			t.flushBuffer()
			t.logAction(ObjectAccessed)

		default:
			t.buffer += str
			t.logAction(BufferAdded)
		}
	}

	if eoiErr := t.validateEndOfInput(); eoiErr != nil {
		return nil, eoiErr
	}

	t.flushBuffer()
	t.logAction(Ended)
	return &t.nodes, nil
}
