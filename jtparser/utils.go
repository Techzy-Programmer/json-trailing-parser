package jtparser

type kv = map[string]any

func safeObjectIndexLookup(from any, find string) int {
	elMap, ok := from.(kv)
	if !ok {
		return -1
	}

	elmMatch, found := elMap[find]
	if !found {
		return -2
	}

	elNum, isNum := elmMatch.(float64)
	if !isNum {
		return -3
	}

	return int(elNum)
}

func compactNil[T any](in []T) []T {
	out := in[:0] // reuse underlying array
	for _, v := range in {
		if any(v) != nil {
			out = append(out, v)
		}
	}
	return out
}
