// Package utils :
// Contains the implementation of a simple id set (uint64)
// optimization with the map default implementation
// This struct contains only the minimal code needed by the project
package utils

// IDSet :
type IDSet struct {
	m map[uint64]struct{}
}

// NewIDSet :
// Constructor
func NewIDSet() IDSet {
	return IDSet{
		m: make(map[uint64]struct{}),
	}
}

// NewIDSetFromSlice :
// Constructor
func NewIDSetFromSlice(li []uint64) IDSet {
	res := NewIDSet()

	for _, num := range li {
		res.Add(num)
	}

	return res
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

// Contains :
// Returns if an element is in the set
func (s IDSet) Contains(elem uint64) bool {
	_, ok := s.m[elem]
	return ok
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
