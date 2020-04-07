package engine_test

import (
	"testing"
)

func TestGetObjectUnknown(t *testing.T) {
	_, err := mia.GetObject("alice")
	if err == nil {
		t.Errorf("get unknown shouldn't work")
	}
}

func TestAddObject(t *testing.T) {
	err := mia.AddObjectToRoot("bobby")

	if err != nil {
		t.Errorf("add should work : %s", err.Error())
	}

	subj, err := mia.GetObject("bobby")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "bobby" {
		t.Errorf("name is not the same")
	}

	err = mia.AddObjectToRoot("")

	if err == nil {
		t.Errorf("should not add an empty Object")
	}

	err = mia.AddObjectToRoot("bobby")

	if err == nil {
		t.Errorf("should not add a Object that's already existe")
	}

	_ = mia.AddObjectToRoot("alice")
	_ = mia.AddObjectToRoot("carole")
	_ = mia.AddObjectToRoot("david")

	subj, err = mia.GetObject("alice")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "alice" {
		t.Errorf("alice is not alice")
	}
}

func TestRemoveObject(t *testing.T) {
	_ = mia.AddObjectToRoot("elodie")

	err := mia.RemoveObject("elodie")
	if err != nil {
		t.Errorf("remove should work : %s", err.Error())
	}

	err = mia.RemoveObject("engie")
	if err == nil {
		t.Errorf("remove should failed, engie is not in mia")
	}
}

func TestRenameObject(t *testing.T) {
	_ = mia.AddObjectToRoot("fan")
	_ = mia.AddObjectToRoot("gwen")

	fanFromDB1, _ := mia.GetObject("fan")

	err := mia.RenameObject("fan", "new fan")

	if err != nil {
		t.Errorf("rename should work : %s", err.Error())
	}

	fanFromDB2, _ := mia.GetObject("new fan")

	if fanFromDB1.ID != fanFromDB2.ID {
		t.Errorf("rename should not change IDs")
	}

	err = mia.RenameObject("folie", "try try")

	if err == nil {
		t.Errorf("rename non existing Object should have failed")
	}

	err = mia.RenameObject("gwen", "")

	if err == nil {
		t.Errorf("rename should have failed cause new name is empty")
	}

	err = mia.RenameObject("gwen", "new fan")

	if err == nil {
		t.Errorf("rename should have failed user existed")
	}
}

func TestAddObjectLink(t *testing.T) {
	_ = mia.AddObjectToRoot("helene")
	_ = mia.AddObjectToRoot("ismail")

	err := mia.AddObjectLink("helene", "ismail")

	if err != nil {
		t.Errorf("add Object link should work : %s", err.Error())
	}

	err = mia.AddObjectLink("ismail", "joseph")

	if err == nil {
		t.Errorf("add Object link shouldn't work cause joseph not in mia")
	}

	err = mia.AddObjectLink("helene", "ismail")

	if err == nil {
		t.Errorf("add Object link shouldn't work cause Object link already exist")
	}
}

func TestRemoveObjectLink(t *testing.T) {
	_ = mia.AddObjectToRoot("kevin")
	_ = mia.AddObjectToRoot("laure")

	_ = mia.AddObjectLink("kevin", "laure")

	err := mia.RemoveObjectLink("kevin", "laure")

	if err != nil {
		t.Errorf("remove Object link should have worked")
	}

	err = mia.RemoveObjectLink("laure", "maude")

	if err == nil {
		t.Errorf("remove Object link should have failed cause it doesn't exist")
	}
}
