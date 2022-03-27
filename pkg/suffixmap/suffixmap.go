package suffixmap

import (
	"strings"
)

// TODO: SuffixMap will eventually become Ukkonen's Algorithm
// when we have time to implement
type SuffixMap interface {
	Get(string) (interface{}, bool)
	GetAll() map[string]interface{}
	Put(string, interface{})
	Delete(string) bool
	Size() int
}

type untypedSuffixMap struct {
	m map[string]interface{}
}

func NewSuffixMap(m map[string]interface{}) SuffixMap {
	if m == nil {
		m = make(map[string]interface{})
	}
	return &untypedSuffixMap{
		m: m,
	}
}

func (usm *untypedSuffixMap) Size() int {
	return len(usm.m)
}

func (usm *untypedSuffixMap) Get(key string) (interface{}, bool) {
	rv, ok := usm.m[key]
	if ok {
		return rv, true
	}
	for k, v := range usm.m {
		if suffixMatches(k, key) {
			return v, true
		}
	}
	return nil, false
}

func (usm *untypedSuffixMap) GetAll() map[string]interface{} {
	return usm.m
}

func (usm *untypedSuffixMap) Delete(key string) bool {
	_, ok := usm.m[key]
	if ok {
		delete(usm.m, key)
		return true
	}
	for k := range usm.m {
		if suffixMatches(k, key) {
			delete(usm.m, k)
			return true
		}
	}
	return false
}

func (usm *untypedSuffixMap) Put(k string, v interface{}) {
	usm.m[k] = v
}

func SuffixMatches(s, suffix string) bool {
	return suffixMatches(s, suffix)
}

func suffixMatches(s, suffix string) bool {
	return strings.HasSuffix(s, suffix) && (strings.HasSuffix(s, "."+suffix) || s == suffix)
}
