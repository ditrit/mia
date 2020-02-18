// Package utils :
// Contains the implementation of a simple id set (uint64)
// optimization with the map default implementation
// This struct contains only the minimal code needed by the project
// TODO test this struct
package utils

// IDSet :
type IDSet struct {
	m map[uint64]struct{}
}

// NewIDSet :
// Constructor
func NewIDSet() IDSet {
	return IDSet{}
}

// Add :
// Add a fresh element
func (s IDSet) Add(elem uint64) {
	s.m[elem] = struct{}{}
}

// Remove :
// Remove an element
func (s IDSet) Remove(elem uint64) {
	delete(s.m, elem)
}

// ToSlice :
// convert the set in a slice
func (s IDSet) ToSlice() []uint64 {
	res := make([]uint64, len(s.m))
	index := 0

	for key := range s.m {
		res[index] = key
		index++
	}

	return res
}
