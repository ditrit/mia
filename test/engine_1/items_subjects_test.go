package engine_test

import (
	"testing"
)

func TestGetSubjectUnknown(t *testing.T) {
	_, err := mia.GetSubject("alice")
	if err == nil {
		t.Errorf("get unknown shouldn't work")
	}
}

func TestAddSubject(t *testing.T) {
	err := mia.AddSubjectToRoot("bobby")

	if err != nil {
		t.Errorf("add should work : %s", err.Error())
	}

	subj, err := mia.GetSubject("bobby")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "bobby" {
		t.Errorf("name is not the same")
	}

	err = mia.AddSubjectToRoot("")

	if err == nil {
		t.Errorf("should not add an empty subject")
	}

	err = mia.AddSubjectToRoot("bobby")

	if err == nil {
		t.Errorf("should not add a subject that's already exists")
	}

	_ = mia.AddSubjectToRoot("alice")
	_ = mia.AddSubjectToRoot("carole")
	_ = mia.AddSubjectToRoot("david")

	subj, err = mia.GetSubject("alice")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "alice" {
		t.Errorf("alice is not alice")
	}
}

func TestRemoveSubject(t *testing.T) {
	_ = mia.AddSubjectToRoot("elodie")

	err := mia.RemoveSubject("elodie")
	if err != nil {
		t.Errorf("remove should work : %s", err.Error())
	}

	err = mia.RemoveSubject("engie")
	if err == nil {
		t.Errorf("remove should failed, engie is not in mia")
	}
}

func TestRenameSubject(t *testing.T) {
	_ = mia.AddSubjectToRoot("fan")
	_ = mia.AddSubjectToRoot("gwen")

	fanFromDB1, _ := mia.GetSubject("fan")

	err := mia.RenameSubject("fan", "new fan")

	if err != nil {
		t.Errorf("rename should work : %s", err.Error())
	}

	fanFromDB2, _ := mia.GetSubject("new fan")

	if fanFromDB1.ID != fanFromDB2.ID {
		t.Errorf("rename should not change IDs")
	}

	err = mia.RenameSubject("folie", "try try")

	if err == nil {
		t.Errorf("rename non existing subject should have failed")
	}

	err = mia.RenameSubject("gwen", "")

	if err == nil {
		t.Errorf("rename should have failed cause new name is empty")
	}

	err = mia.RenameSubject("gwen", "new fan")

	if err == nil {
		t.Errorf("rename should have failed user existed")
	}
}

func TestAddSubjectLink(t *testing.T) {
	_ = mia.AddSubjectToRoot("helene")
	_ = mia.AddSubjectToRoot("ismail")

	err := mia.AddSubjectLink("helene", "ismail")

	if err != nil {
		t.Errorf("add subject link should work : %s", err.Error())
	}

	err = mia.AddSubjectLink("ismail", "joseph")

	if err == nil {
		t.Errorf("add subject link shouldn't work cause joseph not in mia")
	}

	err = mia.AddSubjectLink("helene", "ismail")

	if err == nil {
		t.Errorf("add subject link shouldn't work cause subject link already exist")
	}
}

func TestRemoveSubjectLink(t *testing.T) {
	_ = mia.AddSubjectToRoot("kevin")
	_ = mia.AddSubjectToRoot("laure")

	_ = mia.AddSubjectLink("kevin", "laure")

	err := mia.RemoveSubjectLink("kevin", "laure")

	if err != nil {
		t.Errorf("remove subject link should have worked")
	}

	err = mia.RemoveSubjectLink("laure", "maude")

	if err == nil {
		t.Errorf("remove subject link should have failed cause it doesn't exist")
	}
}
