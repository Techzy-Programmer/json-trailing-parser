package tokenizer

func (t *Tokenizer) logAction(act TokenAction) {
	t.history = append(t.history, act)
}

func (t *Tokenizer) getLastAction() TokenAction {
	return t.history[len(t.history)-1]
}

func (t *Tokenizer) flushBuffer() {
	if t.buffer == "" {
		return
	}

	t.nodes = append(t.nodes, Node{
		IsArray:  t.arrayMode,
		Accessor: t.buffer,
	})

	t.buffer = ""
}
