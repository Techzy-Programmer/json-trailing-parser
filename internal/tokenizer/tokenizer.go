package tokenizer

// Tokenization rules for JSON Trailing Parser
// ignore leading and trailing dots

const (
	ObjectAccessor = '.'
	ArrayStart     = '['
	ArrayEnd       = ']'
	EscapeChar     = '\\'
	Wildcard       = '*'
	TemplateStart  = '{' // To define start of the template accessor
	TemplateEnd    = '}' // To define end of the template accessor
)

type TokenAction int

const (
	Started TokenAction = iota
	BufferAdded
	ArrayStarted
	ArrayEnded
	ObjectAccessed
	EscapeStarted
	EscapeEnded
	TemplateStarted
	TemplateEnded
	Ended
)

type Tokenizer struct {
	query      string
	buffer     string
	arrayMode  bool
	escapeMode bool
	history    []TokenAction
	nodes      []Node
}

func NewTokenizer(q string) *Tokenizer {
	return &Tokenizer{
		query:   q,
		history: make([]TokenAction, 0),
	}
}

type Node struct {
	IsArray  bool
	Accessor string
	GetAll   bool
}
