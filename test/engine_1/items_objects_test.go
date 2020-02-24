package engine_test

import (
	"testing"
)

func TestGetObjectUnknown(t *testing.T) {
	_, err := iam.GetObject("alice")
	if err == nil {
		t.Errorf("get unknown shouldn't work")
	}
}

func TestAddObject(t *testing.T) {
	err := iam.AddObjectToRoot("bobby")

	if err != nil {
		t.Errorf("add should work : %s", err.Error())
	}

	subj, err := iam.GetObject("bobby")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "bobby" {
		t.Errorf("name is not the same")
	}

	err = iam.AddObjectToRoot("")

	if err == nil {
		t.Errorf("should not add an empty Object")
	}

	err = iam.AddObjectToRoot("bobby")

	if err == nil {
		t.Errorf("should not add a Object that's already existe")
	}

	_ = iam.AddObjectToRoot("alice")
	_ = iam.AddObjectToRoot("carole")
	_ = iam.AddObjectToRoot("david")

	subj, err = iam.GetObject("alice")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "alice" {
		t.Errorf("alice is not alice")
	}
}

func TestRemoveObject(t *testing.T) {
	_ = iam.AddObjectToRoot("elodie")

	err := iam.RemoveObject("elodie")
	if err != nil {
		t.Errorf("remove should work : %s", err.Error())
	}

	err = iam.RemoveObject("engie")
	if err == nil {
		t.Errorf("remove should failed, engie is not in iam")
	}
}

func TestRenameObject(t *testing.T) {
	_ = iam.AddObjectToRoot("fan")
	_ = iam.AddObjectToRoot("gwen")

	fanFromDB1, _ := iam.GetObject("fan")

	err := iam.RenameObject("fan", "new fan")

	if err != nil {
		t.Errorf("rename should work : %s", err.Error())
	}

	fanFromDB2, _ := iam.GetObject("new fan")

	if fanFromDB1.ID != fanFromDB2.ID {
		t.Errorf("rename should not change IDs")
	}

	err = iam.RenameObject("folie", "try try")

	if err == nil {
		t.Errorf("rename non existing Object should have failed")
	}

	err = iam.RenameObject("gwen", "")

	if err == nil {
		t.Errorf("rename should have failed cause new name is empty")
	}
}

func TestAddObjectLink(t *testing.T) {
	_ = iam.AddObjectToRoot("helene")
	_ = iam.AddObjectToRoot("ismail")

	err := iam.AddObjectLink("helene", "ismail")

	if err != nil {
		t.Errorf("add Object link should work : %s", err.Error())
	}

	err = iam.AddObjectLink("ismail", "joseph")

	if err == nil {
		t.Errorf("add Object link shouldn't work cause joseph not in iam")
	}

	err = iam.AddObjectLink("helene", "ismail")

	if err == nil {
		t.Errorf("add Object link shouldn't work cause Object link already exist")
	}
}

func TestRemoveObjectLink(t *testing.T) {
	_ = iam.AddObjectToRoot("kevin")
	_ = iam.AddObjectToRoot("laure")

	_ = iam.AddObjectLink("kevin", "laure")

	err := iam.RemoveObjectLink("kevin", "laure")

	if err != nil {
		t.Errorf("remove Object link should have worked")
	}

	err = iam.RemoveObjectLink("laure", "maude")

	if err == nil {
		t.Errorf("remove Object link should have failed cause it doesn't exist")
	}
}
