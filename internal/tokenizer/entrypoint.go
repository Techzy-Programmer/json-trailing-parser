package tokenizer

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Techzy-Programmer/json-trailing-parser/jtparser"
)

func (t *Tokenizer) Tokenize() (*[]Node, error) {
	t.logAction(started)
	t.nodes = make([]Node, 0)

	if strings.HasPrefix(t.query, string(arrayStart)) {
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
			if !slices.Contains(escapables, r) {
				return nil, &jtparser.ErrInvalidQuery{
					Query:  t.query,
					Reason: fmt.Sprintf("expecting escape sequence found %c at position %d", r, i),
				}
			}

			t.buffer += str
			t.escapeMode = false
			t.logAction(escapeEnded)

			continue
		}

		switch r {
		case escapeToken:
			t.escapeMode = true
			t.logAction(escapeStarted)

		case arrayStart:
			if asErr := t.validateArrayModeState(true, i); asErr != nil {
				return nil, asErr
			}

			t.arrayMode = true
			t.flushBuffer()
			t.logAction(arrayStarted)

		case arrayEnd:
			if asErr := t.validateArrayModeState(false, i); asErr != nil {
				return nil, asErr
			}

			t.flushBuffer()
			t.arrayMode = false
			t.logAction(arrayEnded)

		case objectAccessor:
			if oaErr := t.validateObjectAccessor(i); oaErr != nil {
				return nil, oaErr
			}

			t.flushBuffer()
			t.logAction(objectAccessed)

		default:
			t.buffer += str
			t.logAction(bufferAdded)
		}
	}

	if eoiErr := t.validateEndOfInput(); eoiErr != nil {
		return nil, eoiErr
	}

	t.flushBuffer()
	t.logAction(ended)
	return &t.nodes, nil
}
