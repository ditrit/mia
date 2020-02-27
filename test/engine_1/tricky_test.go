package engine_test

import (
	"testing"
)

func TestTrickyAdd(t *testing.T) {
	_ = iam.AddSubjectToRoot("Batman")
	_ = iam.AddSubject("Alfred", "Batman")

	err := iam.AddSubjectLink("Alfred", "Batman")

	if err == nil {
		t.Errorf("should be an error, cycle appears")
	}
}

func TestTrickyRemove(t *testing.T) {
	_ = iam.AddObjectToRoot("Git")
	_ = iam.AddObject("Github", "Git")
	_ = iam.AddObject("Gitlab", "Git")
	_ = iam.AddObject("Jenkins", "Gitlab")

	err := iam.RemoveObjectLink("Gitlab", "Jenkins")

	if err == nil {
		t.Errorf("should be an error, connectivity is broken")
	}

	_ = iam.AddObjectLink("Github", "Jenkins")

	err = iam.RemoveObjectLink("Gitlab", "Jenkins")

	if err != nil {
		t.Errorf("should work now, new route available")
	}
}
