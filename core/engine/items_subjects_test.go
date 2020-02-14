package engine_test

import (
	"testing"
)

func TestGetSubjectUnknown(t *testing.T) {
	_, err := iam.GetSubject("alice")
	if err == nil {
		t.Errorf("get unknown shouldn't work")
	}
}

func TestAddSubject(t *testing.T) {
	err := iam.AddSubject("bobby")

	if err != nil {
		t.Errorf("add should work : %s", err.Error())
	}

	subj, err := iam.GetSubject("bobby")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "bobby" {
		t.Errorf("name is not the same")
	}

	err = iam.AddSubject("")

	if err == nil {
		t.Errorf("should not add an empty subject")
	}

	err = iam.AddSubject("bobby")

	if err == nil {
		t.Errorf("should not add a subject that's already exists")
	}

	_ = iam.AddSubject("alice")
	_ = iam.AddSubject("carole")
	_ = iam.AddSubject("david")

	subj, err = iam.GetSubject("alice")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "alice" {
		t.Errorf("alice is not alice")
	}
}

func TestRemoveSubject(t *testing.T) {
	_ = iam.AddSubject("elodie")

	err := iam.RemoveSubject("elodie")
	if err != nil {
		t.Errorf("remove should work : %s", err.Error())
	}

	err = iam.RemoveSubject("engie")
	if err == nil {
		t.Errorf("remove should failed, engie is not in iam")
	}
}

func TestRenameSubject(t *testing.T) {
	_ = iam.AddSubject("fan")
	_ = iam.AddSubject("gwen")

	fanFromDB1, _ := iam.GetSubject("fan")

	err := iam.RenameSubject("fan", "new fan")

	if err != nil {
		t.Errorf("rename should work : %s", err.Error())
	}

	fanFromDB2, _ := iam.GetSubject("new fan")

	if fanFromDB1.ID != fanFromDB2.ID {
		t.Errorf("rename should not change IDs")
	}

	err = iam.RenameSubject("folie", "try try")

	if err == nil {
		t.Errorf("rename non existing subject should have failed")
	}

	err = iam.RenameSubject("gwen", "")

	if err == nil {
		t.Errorf("rename should have failed cause new name is empty")
	}
}

func TestAddSubjectLink(t *testing.T) {
	_ = iam.AddSubject("helene")
	_ = iam.AddSubject("ismail")

	err := iam.AddSubjectLink("helene", "ismail")

	if err != nil {
		t.Errorf("add subject link should work : %s", err.Error())
	}

	err = iam.AddSubjectLink("ismail", "joseph")

	if err == nil {
		t.Errorf("add subject link shouldn't work cause joseph not in iam")
	}

	err = iam.AddSubjectLink("helene", "ismail")

	if err == nil {
		t.Errorf("add subject link shouldn't work cause subject link already exist")
	}
}

func TestRemoveSubjectLink(t *testing.T) {
	_ = iam.AddSubject("kevin")
	_ = iam.AddSubject("laure")

	_ = iam.AddSubjectLink("kevin", "laure")

	err := iam.RemoveSubjectLink("kevin", "laure")

	if err != nil {
		t.Errorf("remove subject link should have worked")
	}

	err = iam.RemoveSubjectLink("laure", "maude")

	if err == nil {
		t.Errorf("remove subject link should have failed cause it doesn't exist")
	}
}
