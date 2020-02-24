package engine_test

import (
	"testing"
)

func TestGetDomainUnknown(t *testing.T) {
	_, err := iam.GetDomain("alice")
	if err == nil {
		t.Errorf("get unknown shouldn't work")
	}
}

//nolint: goconst
func TestAddDomain(t *testing.T) {
	err := iam.AddDomainToRoot("bobby")

	if err != nil {
		t.Errorf("add should work : %s", err.Error())
	}

	subj, err := iam.GetDomain("bobby")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "bobby" {
		t.Errorf("name is not the same")
	}

	err = iam.AddDomainToRoot("")

	if err == nil {
		t.Errorf("should not add an empty Domain")
	}

	err = iam.AddDomainToRoot("bobby")

	if err == nil {
		t.Errorf("should not add a Domain that's already existe")
	}

	_ = iam.AddDomainToRoot("alice")
	_ = iam.AddDomainToRoot("carole")
	_ = iam.AddDomainToRoot("david")

	subj, err = iam.GetDomain("alice")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "alice" {
		t.Errorf("alice is not alice")
	}
}

func TestRemoveDomain(t *testing.T) {
	_ = iam.AddDomainToRoot("elodie")

	err := iam.RemoveDomain("elodie")
	if err != nil {
		t.Errorf("remove should work : %s", err.Error())
	}

	err = iam.RemoveDomain("engie")
	if err == nil {
		t.Errorf("remove should failed, engie is not in iam")
	}
}

func TestRenameDomain(t *testing.T) {
	_ = iam.AddDomainToRoot("fan")
	_ = iam.AddDomainToRoot("gwen")

	fanFromDB1, _ := iam.GetDomain("fan")

	err := iam.RenameDomain("fan", "new fan")

	if err != nil {
		t.Errorf("rename should work : %s", err.Error())
	}

	fanFromDB2, _ := iam.GetDomain("new fan")

	if fanFromDB1.ID != fanFromDB2.ID {
		t.Errorf("rename should not change IDs")
	}

	err = iam.RenameDomain("folie", "try try")

	if err == nil {
		t.Errorf("rename non existing Domain should have failed")
	}

	err = iam.RenameDomain("gwen", "")

	if err == nil {
		t.Errorf("rename should have failed cause new name is empty")
	}
}

func TestAddDomainLink(t *testing.T) {
	_ = iam.AddDomainToRoot("helene")
	_ = iam.AddDomainToRoot("ismail")

	err := iam.AddDomainLink("helene", "ismail")

	if err != nil {
		t.Errorf("add Domain link should work : %s", err.Error())
	}

	err = iam.AddDomainLink("ismail", "joseph")

	if err == nil {
		t.Errorf("add Domain link shouldn't work cause joseph not in iam")
	}

	err = iam.AddDomainLink("helene", "ismail")

	if err == nil {
		t.Errorf("add Domain link shouldn't work cause Domain link already exist")
	}
}

func TestRemoveDomainLink(t *testing.T) {
	_ = iam.AddDomainToRoot("kevin")
	_ = iam.AddDomainToRoot("laure")

	_ = iam.AddDomainLink("kevin", "laure")

	err := iam.RemoveDomainLink("kevin", "laure")

	if err != nil {
		t.Errorf("remove Domain link should have worked")
	}

	err = iam.RemoveDomainLink("laure", "maude")

	if err == nil {
		t.Errorf("remove Domain link should have failed cause it doesn't exist")
	}
}
