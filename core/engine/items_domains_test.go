package engine_test

import (
	"iam/core/model"
	"testing"
)

func TestGetDomainUnknown(t *testing.T) {
	_, err := iam.GetDomain("alice")
	if err == nil {
		t.Errorf("get unknown shouldn't work")
	}
}

func TestAddDomain(t *testing.T) {
	bobby, _ := model.NewDomain("bobby")

	err := iam.AddDomain(*bobby)

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

	emptySubj := new(model.Item)
	emptySubj.Type = model.ITEM_TYPE_SUBJ
	emptySubj.Name = ""

	err = iam.AddDomain(*emptySubj)

	if err == nil {
		t.Errorf("should not add an empty Domain")
	}

	err = iam.AddDomain(*bobby)

	if err == nil {
		t.Errorf("should not add a Domain that's already existe")
	}

	alice, _ := model.NewDomain("alice")
	carole, _ := model.NewDomain("carole")
	david, _ := model.NewDomain("david")

	_ = iam.AddDomain(*alice)
	_ = iam.AddDomain(*carole)
	_ = iam.AddDomain(*david)

	subj, err = iam.GetDomain("alice")
	if err != nil {
		t.Errorf("get should work : %s", err.Error())
	}

	if subj.Name != "alice" {
		t.Errorf("alice is not alice")
	}
}

func TestRemoveDomain(t *testing.T) {
	elodie, _ := model.NewDomain("elodie")
	engie, _ := model.NewDomain("engie")

	_ = iam.AddDomain(*elodie)

	err := iam.RemoveDomain(*elodie)
	if err != nil {
		t.Errorf("remove should work : %s", err.Error())
	}

	err = iam.RemoveDomain(*engie)
	if err == nil {
		t.Errorf("remove should failed, engie is not in iam")
	}
}

func TestRenameDomain(t *testing.T) {
	fan, _ := model.NewDomain("fan")
	gwen, _ := model.NewDomain("gwen")
	folie, _ := model.NewDomain("folie")

	_ = iam.AddDomain(*fan)
	_ = iam.AddDomain(*gwen)

	fanFromDB1, _ := iam.GetDomain("fan")

	err := iam.RenameDomain(*fan, "new fan")

	if err != nil {
		t.Errorf("rename should work : %s", err.Error())
	}

	fanFromDB2, _ := iam.GetDomain("new fan")

	if fanFromDB1.ID != fanFromDB2.ID {
		t.Errorf("rename should not change IDs")
	}

	err = iam.RenameDomain(*folie, "try try")

	if err == nil {
		t.Errorf("rename non existing Domain should have failed")
	}

	err = iam.RenameDomain(*gwen, "")

	if err == nil {
		t.Errorf("rename should have failed cause new name is empty")
	}
}

func TestAddDomainLink(t *testing.T) {
	helene, _ := model.NewDomain("hélene")
	ismail, _ := model.NewDomain("ismail")
	joseph, _ := model.NewDomain("joseph")

	_ = iam.AddDomain(*helene)
	_ = iam.AddDomain(*ismail)

	err := iam.AddDomainLink(*helene, *ismail)

	if err != nil {
		t.Errorf("add Domain link should work : %s", err.Error())
	}

	err = iam.AddDomainLink(*ismail, *joseph)

	if err == nil {
		t.Errorf("add Domain link shouldn't work cause joseph not in iam")
	}

	err = iam.AddDomainLink(*helene, *ismail)

	if err == nil {
		t.Errorf("add Domain link shouldn't work cause Domain link already exist")
	}
}

func TestRemoveDomainLink(t *testing.T) {
	kevin, _ := model.NewDomain("kevin")
	laure, _ := model.NewDomain("laure")
	maude, _ := model.NewDomain("maude")

	_ = iam.AddDomain(*kevin)
	_ = iam.AddDomain(*laure)

	_ = iam.AddDomainLink(*kevin, *laure)

	err := iam.RemoveDomainLink(*kevin, *laure)

	if err != nil {
		t.Errorf("remove Domain link should have worked")
	}

	err = iam.RemoveDomainLink(*laure, *maude)

	if err == nil {
		t.Errorf("remove Domain link should have failed cause it doesn't exist")
	}
}
