package engine_test

import (
	"testing"
)

func TestGetDomainUnknown(t *testing.T) {
	_, err := mia.GetDomain("alice")
	if err == nil {
		t.Errorf("get unknown shouldn't work")
	}
}

//nolint: goconst
func TestAddDomain(t *testing.T) {
	err := mia.AddDomainToRoot("bobby")

	if err != nil {
		t.Errorf("add should work : %s", err.Error())
	}

	subj, err := mia.GetDomain("bobby")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "bobby" {
		t.Errorf("name is not the same")
	}

	err = mia.AddDomainToRoot("")

	if err == nil {
		t.Errorf("should not add an empty Domain")
	}

	err = mia.AddDomainToRoot("bobby")

	if err == nil {
		t.Errorf("should not add a Domain that's already existe")
	}

	_ = mia.AddDomainToRoot("alice")
	_ = mia.AddDomainToRoot("carole")
	_ = mia.AddDomainToRoot("david")

	subj, err = mia.GetDomain("alice")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "alice" {
		t.Errorf("alice is not alice")
	}
}

func TestRemoveDomain(t *testing.T) {
	_ = mia.AddDomainToRoot("elodie")

	err := mia.RemoveDomain("elodie")
	if err != nil {
		t.Errorf("remove should work : %s", err.Error())
	}

	err = mia.RemoveDomain("engie")
	if err == nil {
		t.Errorf("remove should failed, engie is not in mia")
	}
}

func TestRenameDomain(t *testing.T) {
	_ = mia.AddDomainToRoot("fan")
	_ = mia.AddDomainToRoot("gwen")

	fanFromDB1, _ := mia.GetDomain("fan")

	err := mia.RenameDomain("fan", "new fan")

	if err != nil {
		t.Errorf("rename should work : %s", err.Error())
	}

	fanFromDB2, _ := mia.GetDomain("new fan")

	if fanFromDB1.ID != fanFromDB2.ID {
		t.Errorf("rename should not change IDs")
	}

	err = mia.RenameDomain("folie", "try try")

	if err == nil {
		t.Errorf("rename non existing Domain should have failed")
	}

	err = mia.RenameDomain("gwen", "")

	if err == nil {
		t.Errorf("rename should have failed cause new name is empty")
	}

	err = mia.RenameDomain("gwen", "new fan")

	if err == nil {
		t.Errorf("rename should have failed user existed")
	}
}

func TestAddDomainLink(t *testing.T) {
	_ = mia.AddDomainToRoot("helene")
	_ = mia.AddDomainToRoot("ismail")

	err := mia.AddDomainLink("helene", "ismail")

	if err != nil {
		t.Errorf("add Domain link should work : %s", err.Error())
	}

	err = mia.AddDomainLink("ismail", "joseph")

	if err == nil {
		t.Errorf("add Domain link shouldn't work cause joseph not in mia")
	}

	err = mia.AddDomainLink("helene", "ismail")

	if err == nil {
		t.Errorf("add Domain link shouldn't work cause Domain link already exist")
	}
}

func TestRemoveDomainLink(t *testing.T) {
	_ = mia.AddDomainToRoot("kevin")
	_ = mia.AddDomainToRoot("laure")

	_ = mia.AddDomainLink("kevin", "laure")

	err := mia.RemoveDomainLink("kevin", "laure")

	if err != nil {
		t.Errorf("remove Domain link should have worked")
	}

	err = mia.RemoveDomainLink("laure", "maude")

	if err == nil {
		t.Errorf("remove Domain link should have failed cause it doesn't exist")
	}
}
