package model_test

import (
	"iam/core/constant"
	"iam/core/model"
	"testing"
)

func TestNewSubject(t *testing.T) {
	s, err := model.NewSubject("alice")

	switch {
	case err != nil:
		t.Errorf("should have created subject")
	case s.Type != model.ITEM_TYPE_SUBJ:
		t.Errorf("should have the subj type")
	case s.Name != "alice":
		t.Errorf("should have the good name")
	}

	_, err2 := model.NewSubject("")

	if err2 == nil {
		t.Errorf("subject can't have empty name")
	}
}

func TestNewObject(t *testing.T) {
	s, err := model.NewObject("bob")

	switch {
	case err != nil:
		t.Errorf("should have created object")
	case s.Type != model.ITEM_TYPE_OBJ:
		t.Errorf("should have the object type")
	case s.Name != "bob":
		t.Errorf("should have the good name")
	}

	_, err2 := model.NewObject("")

	if err2 == nil {
		t.Errorf("object can't have empty name")
	}
}

func TestNewDomain(t *testing.T) {
	s, err := model.NewDomain("charlie")

	switch {
	case err != nil:
		t.Errorf("should have created domain")
	case s.Type != model.ITEM_TYPE_DOMAIN:
		t.Errorf("should have the domain type")
	case s.Name != "charlie":
		t.Errorf("should have the good name")
	}
}

func TestRoots(t *testing.T) {
	if st, _ := model.GetRootNameWithType(model.ITEM_TYPE_DOMAIN); st != constant.ROOT_DOMAINS {
		t.Errorf("domain root hasn't the good name")
	}

	if st, _ := model.GetRootNameWithType(model.ITEM_TYPE_SUBJ); st != constant.ROOT_SUBJECTS {
		t.Errorf("subject root hasn't the good name")
	}

	if st, _ := model.GetRootNameWithType(model.ITEM_TYPE_OBJ); st != constant.ROOT_OBJECTS {
		t.Errorf("object root hasn't the good name")
	}

	if _, err := model.GetRootNameWithType(42); err == nil { //nolint: gomnd
		t.Errorf("GetRootNameWithType should return error on unknown type")
	}
}
