package model_test

import (
	"mia/core/model"
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

	_, err = model.NewRole("")

	if err == nil {
		t.Errorf("should have failed, role name can't be empty")
	}
}
