package jtparser

import (
	"fmt"

	"github.com/Techzy-Programmer/json-trailing-parser/internal/tokenizer"
)

type MonoParser[T any] struct {
	SearchSpace []any
	nodes       *[]tokenizer.Node
	walker      *walker
	result      T
}

func NewMonoParser[T any](query string, searchSpace []any) (*MonoParser[T], error) {
	tk := tokenizer.NewTokenizer(query)
	nodes, tokErr := tk.Tokenize()
	if tokErr != nil {
		return nil, tokErr
	}

	return &MonoParser[T]{
		walker: &walker{
			searchSpace: &searchSpace,
			path:        "$",
		},

		SearchSpace: searchSpace,
		nodes:       nodes,
	}, nil
}

func (p *MonoParser[T]) ChangeQuery(newQuery string) error {
	tk := tokenizer.NewTokenizer(newQuery)
	nodes, tokErr := tk.Tokenize()
	if tokErr != nil {
		return tokErr
	}

	p.nodes = nodes
	p.walker.path = "$"
	p.walker.previousResults = nil
	p.walker.searchSpace = &p.SearchSpace

	return nil
}

func (p *MonoParser[T]) Parse() (*T, error) {
	for _, node := range *p.nodes {
		if p.walker.previousResults == nil {
			if prevErr := p.walker.scanFind(node.Accessor); prevErr != nil {
				return nil, prevErr
			}

			continue
		}

		result, osErr := p.walker.objectCollect(
			p.walker.previousResults,
			node.Accessor,
			node.IsArray,
		)
		if osErr != nil {
			return nil, osErr
		}

		p.walker.previousResults = &result
	}

	t, typed := (*p.walker.previousResults).(T)
	if !typed {
		return nil, fmt.Errorf("can't parse to requested type '%T' expecting '%T'", p.result, (*p.walker.previousResults))
	}

	return &t, nil
}
