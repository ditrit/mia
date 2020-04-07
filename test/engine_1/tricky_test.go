package engine_test

import (
	"testing"
)

func TestTrickyAdd(t *testing.T) {
	_ = mia.AddSubjectToRoot("Batman")
	_ = mia.AddSubject("Alfred", "Batman")

	err := mia.AddSubjectLink("Alfred", "Batman")

	if err == nil {
		t.Errorf("should be an error, cycle appears")
	}
}

func TestTrickyRemove(t *testing.T) {
	_ = mia.AddObjectToRoot("Git")
	_ = mia.AddObject("Github", "Git")
	_ = mia.AddObject("Gitlab", "Git")
	_ = mia.AddObject("Jenkins", "Gitlab")

	err := mia.RemoveObjectLink("Gitlab", "Jenkins")

	if err == nil {
		t.Errorf("should be an error, connectivity is broken")
	}

	_ = mia.AddObjectLink("Github", "Jenkins")

	err = mia.RemoveObjectLink("Gitlab", "Jenkins")

	if err != nil {
		t.Errorf("should work now, new route available")
	}
}
