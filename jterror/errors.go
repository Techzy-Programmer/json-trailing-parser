package jterror

import "fmt"

// ErrInvalidQuery is returned when the query string is malformed.
type ErrInvalidQuery struct {
	Query  string
	Reason string
}

func (e *ErrInvalidQuery) Error() string {
	return fmt.Sprintf("invalid query %q: %s", e.Query, e.Reason)
}

// ErrOutOfBounds is returned when internally referenced index is out of range.
type ErrOutOfBounds struct {
	Index int
	Path  string
}

func (e *ErrOutOfBounds) Error() string {
	return fmt.Sprintf("index %d at path %q is out of bounds", e.Index, e.Path)
}

type ParentType string

const (
	ParentTypeObject ParentType = "object"
	ParentTypeArray  ParentType = "array"
)

// ErrKeyNotFound is returned when a specified key is not found in an object or array.
type ErrKeyNotFound struct {
	Key    string
	Path   string
	Parent ParentType
}

func (e *ErrKeyNotFound) Error() string {
	if e.Parent == ParentTypeObject {
		return fmt.Sprintf("key %q not found in object at path %q", e.Key, e.Path)
	}

	return fmt.Sprintf("key %q not found in array at path %q", e.Key, e.Path)
}

type ErrTypeMismatch struct {
	Path   string
	Key    string
	Object any
}

func (e *ErrTypeMismatch) Error() string {
	return fmt.Sprintf("type mismatch in search space at path %s for key %s\nerrored object >> %v", e.Path, e.Key, e.Object)
}
