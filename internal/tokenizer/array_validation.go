package tokenizer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Techzy-Programmer/json-trailing-parser/jtparser"
)

func (t *Tokenizer) validateArrayMode(curr string, i int) error {
	if !t.arrayMode {
		return nil
	}

	if strings.Contains(t.buffer, string(Wildcard)) && curr != string(ArrayEnd) {
		return &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: fmt.Sprintf("expecting %c got %s at position %d", ArrayEnd, curr, i),
		}
	}

	// If buffer is present already, skip wildcard checking
	if (t.getLastAction() == BufferAdded || curr != string(Wildcard)) && curr != string(ArrayEnd) {
		if _, err := strconv.Atoi(curr); err != nil {
			return &jtparser.ErrInvalidQuery{
				Query:  t.query,
				Reason: fmt.Sprintf("expecting integer, got %s at position %d", curr, i),
			}
		}
	}

	return nil
}

func (t *Tokenizer) validateArrayModeState(isStarting bool, i int) error {
	if isStarting {
		if t.getLastAction() == ObjectAccessed {
			return &jtparser.ErrInvalidQuery{
				Query:  t.query,
				Reason: fmt.Sprintf("illegal dot(.) before array index at position %d", i),
			}
		}

		if t.arrayMode {
			return &jtparser.ErrInvalidQuery{
				Query:  t.query,
				Reason: fmt.Sprintf("%c found at position %d expecting accessor or %c", ArrayEnd, i, ArrayStart),
			}
		}

		return nil
	}

	// query is trying to end the array mode

	if !t.arrayMode { // While the array mode was not previously triggered
		return &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: fmt.Sprintf("%c found at position %d expecting %c", ArrayEnd, i, ArrayStart),
		}
	}

	if len(t.buffer) == 0 { // When there was no accessor provided
		return &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: fmt.Sprintf("expecting accessor found end of array at position %d", i),
		}
	}

	return nil
}
