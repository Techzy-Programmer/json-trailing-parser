package jtparser

import (
	"fmt"

	"github.com/Techzy-Programmer/json-trailing-parser/jterror"
)

const wildcard = "*"

type walker struct {
	searchSpace     *[]any
	previousResults *any
	path            string
}

func (w *walker) scanFind(key string) error {
	for _, element := range *w.searchSpace {
		elIndex := safeObjectIndexLookup(element, key)
		if elIndex < 0 {
			continue
		}

		res, rErr := w.resolveElAt(elIndex, key)
		w.previousResults = &res
		w.path += "." + key

		return rErr
	}

	return &jterror.ErrKeyNotFound{
		Parent: jterror.ParentTypeObject,
		Path:   w.path,
		Key:    key,
	}
}

func (w *walker) objectCollect(root *any, key string, isArray bool) (any, error) {
	results, wasArray := (*root).([]any)
	if wasArray {
		for i, childRoot := range results {
			if isArray && key != wildcard && fmt.Sprintf("%d", i) != key {
				results[i] = nil
				continue
			}

			object, err := w.objectCollect(&childRoot, key, isArray)
			if err != nil {
				return nil, err
			}

			results[i] = object
		}

		return compactNil(results), nil
	}

	var err error
	var object any

	if !isArray {
		object, err = w._searchObject(*root, key)
	} else {
		index, isFloat := (*root).(float64)
		if !isFloat {
			return "", &jterror.ErrTypeMismatch{
				Object: *root,
				Path:   w.path,
				Key:    key,
			}
		}

		object, err = w.resolveElAt(int(index), key)
	}

	if err != nil {
		return "", err
	}

	return object, nil
}

func (w *walker) _searchObject(prevResult any, key string) (any, error) {
	elIndex := safeObjectIndexLookup(prevResult, key)
	if elIndex == -1 {
		return "", &jterror.ErrTypeMismatch{
			Object: prevResult,
			Path:   w.path,
			Key:    key,
		}
	}

	if elIndex < -1 {
		return "", &jterror.ErrKeyNotFound{
			Parent: jterror.ParentTypeObject,
			Path:   w.path,
			Key:    key,
		}
	}

	return w.resolveElAt(elIndex, key)
}

func (w *walker) resolveElAt(index int, path string) (any, error) {
	if index >= len(*w.searchSpace) {
		return nil, &jterror.ErrOutOfBounds{
			Index: index,
			Path:  path,
		}
	}

	return (*w.searchSpace)[index], nil
}
