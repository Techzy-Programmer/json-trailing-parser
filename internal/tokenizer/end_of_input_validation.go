package tokenizer

import (
	"fmt"

	"github.com/Techzy-Programmer/json-trailing-parser/jtparser"
)

func (t *Tokenizer) validateEndOfInput() error {
	pos := fmt.Sprintf("%d", len(t.query)-1)

	if len(t.nodes) == 0 {
		return &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: "expecting token found end of input at position " + pos,
		}
	}

	if t.escapeMode {
		return &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: "expecting a character after escape sequence found end of input at position " + pos,
		}
	}

	if t.arrayMode {
		return &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: fmt.Sprintf("expecting accessor or %c found end of input at position %s", ArrayEnd, pos),
		}
	}

	if t.getLastAction() == ObjectAccessed {
		return &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: "expecting property name found end of input at position " + pos,
		}
	}

	return nil
}
