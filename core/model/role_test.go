package model_test

import (
	"iam/core/model"
	"testing"
)

func TestRole(t *testing.T) {
	r, err := model.NewRole("test")

	if err != nil {
		t.Errorf("NewRole should have worked")
	}

	if r.Name != "test" {
		t.Errorf("NewRole wrong name")
	}
}
