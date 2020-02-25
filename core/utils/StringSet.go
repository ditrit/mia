// Package utils :
// Contains the implementation of a simple string set
// optimization with the map default implementation
// This struct contains only the minimal code needed by the project
package utils

// StringSet :
type StringSet struct {
	m map[string]struct{}
}

// NewStringSet :
// Constructor
func NewStringSet() StringSet {
	return StringSet{
		m: make(map[string]struct{}),
	}
}

// Add :
// Add a fresh element
func (s StringSet) Add(elem string) {
	s.m[elem] = struct{}{}
}

// Remove :
// Remove an element
func (s StringSet) Remove(elem string) {
	delete(s.m, elem)
}

// Contains :
// Returns if an element is in the set
func (s StringSet) Contains(elem string) bool {
	_, ok := s.m[elem]
	return ok
}

// Pop :
// popping element if not empty
func (s StringSet) Pop() (string, bool) {
	if len(s.m) == 0 {
		return "", false
	}

	var key string

	for k := range s.m {
		key = k
		break
	}

	delete(s.m, key)

	return key, true
}
