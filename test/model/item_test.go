package model_test

import (
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

	_, err2 := model.NewDomain("")

	if err2 == nil {
		t.Errorf("root domain can't be instanciated with NewDomain")
	}
}

func TestRootDomain(t *testing.T) {
	s := model.GetRootDomain()

	if s.Name != "" {
		t.Errorf("root domain must not have root domain")
	}

	if s.Type != model.ITEM_TYPE_DOMAIN {
		t.Errorf("root domain must be of type domain")
	}
}
