package tokenizer

import (
	"fmt"

	"github.com/Techzy-Programmer/json-trailing-parser/jtparser"
)

func (t *Tokenizer) validateObjectAccessor(i int) error {
	if len(t.buffer) == 0 && t.getLastAction() != ArrayEnded {
		return &jtparser.ErrInvalidQuery{
			Query:  t.query,
			Reason: fmt.Sprintf("illegal object accessor at position %d", i),
		}
	}

	return nil
}
