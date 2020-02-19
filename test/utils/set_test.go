package utils_test

import (
	"iam/core/utils"
	"testing"
)

func TestStringSet(t *testing.T) {
	s := utils.NewStringSet()
	elems := []string{"je", "m'appelle", "Joseph", "et", "je", "fais", "des", "tests"}

	for _, elem := range elems {
		s.Add(elem)
		s.Remove(elem)
		s.Add(elem)
	}

	for {
		elem, ok := s.Pop()

		if !ok {
			break
		}

		found := false

		for _, st := range elems {
			if elem == st {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("popping unknown element")
		}
	}

	for _, elem := range elems {
		s.Add(elem)

		st, ok := s.Pop()

		if !ok {
			t.Errorf("the set is not empty")
		}

		if st != elem {
			t.Errorf("the element is not the same")
		}
	}
}

func TestIDSet(t *testing.T) {
	elems := []uint64{0, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	s := utils.NewIDSet()
	s2 := utils.NewIDSetFromSlice(elems)

	for _, elem := range elems {
		if s.Contains(elem) {
			t.Errorf("idset cannot contain this number yet")
		}

		s.Add(elem)

		if !s.Contains(elem) {
			t.Errorf("idset actually should contain elem %d :", elem)
		}

		s.Remove(elem)

		if s.Contains(elem) {
			t.Errorf("idset shouldn't contain this number anymore")
		}

		if !s2.Contains(elem) {
			t.Errorf("should contain after NewFromSlice")
		}
	}

	ft := elems[6] + elems[4] + elems[2] //42

	s2.Add(ft)

	elems = append(elems, ft)

	sliceS2 := s2.ToSlice()

	for _, elem1 := range elems {
		found := false

		for _, elem2 := range sliceS2 {
			if elem1 == elem2 {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("not the same")
			break
		}
	}
}
