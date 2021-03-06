package engine_test

import (
	"testing"

	"mia/core/constant"
	"mia/core/model"
	"mia/core/utils"
)

type link struct {
	parent string
	child  string
}

//nolint: funlen, gocritic, gocyclo, gocognit
func testEquity(
	t *testing.T,
	items []string,
	itemsLinks []link,
	iType model.ItemType,
	vertices []string,
	parentTable map[string][]string,
) {
	rootName, err := model.GetRootNameWithType(iType)

	if err != nil {
		t.Errorf("GetRootNameWithType returns error")
	}

	if len(items)+1 == len(vertices) {
		foundRoot := false

		for _, elem := range vertices {
			if elem == rootName {
				foundRoot = true
				break
			}
		}

		if !foundRoot {
			t.Errorf("root domain not found")
		}
	} else {
		t.Errorf("not the same count of items")
	}

	for _, item := range items {
		found := false

		for _, vertex := range vertices {
			if item == vertex {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("the item %s was not found", item)
		}
	}

	nbLinks := 0

	for child := range parentTable {
		for _, parent := range parentTable[child] {
			found := false

			for _, link := range itemsLinks {
				if link.parent == parent && link.child == child {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("the link wasn't found P: %s,  C: %s", parent, child)
			} else {
				nbLinks++
			}
		}
	}

	if nbLinks != len(itemsLinks) {
		t.Errorf("the number of links are not the same")
	}
}

//nolint: funlen
func TestArchitecture(t *testing.T) {
	subjects := []string{
		"Angel",
		"Barney",
		"Charlie",
		"David",
		"DevTeam",
		"Administration",
	}

	subjectsLinks := []link{
		{parent: constant.ROOT_SUBJECTS, child: "Administration"},
		{parent: constant.ROOT_SUBJECTS, child: "DevTeam"},
		{parent: "Administration", child: "Angel"},
		{parent: "DevTeam", child: "Barney"},
		{parent: "DevTeam", child: "Charlie"},
		{parent: "DevTeam", child: "David"},
	}

	domains := []string{
		"ORNESS",
		"DevTeam",
		"Gandalf",
		"Commercial",
	}

	domainLinks := []link{
		{parent: constant.ROOT_DOMAINS, child: "ORNESS"},
		{parent: "ORNESS", child: "DevTeam"},
		{parent: "ORNESS", child: "Commercial"},
		{parent: "DevTeam", child: "Gandalf"},
	}

	objects := []string{
		"GitLab",
		"GitCommit",
	}

	objectLinks := []link{
		{parent: constant.ROOT_OBJECTS, child: "GitLab"},
		{parent: "GitLab", child: "GitCommit"},
	}

	for _, link := range subjectsLinks {
		if link.parent == constant.ROOT_SUBJECTS {
			_ = mia.AddSubjectToRoot(link.child)
		} else {
			_ = mia.AddSubject(link.child, link.parent)
		}
	}

	for _, link := range domainLinks {
		if link.parent == constant.ROOT_DOMAINS {
			_ = mia.AddDomainToRoot(link.child)
		} else {
			_ = mia.AddDomain(link.child, link.parent)
		}
	}

	for _, link := range objectLinks {
		if link.parent == constant.ROOT_OBJECTS {
			_ = mia.AddObjectToRoot(link.child)
		} else {
			_ = mia.AddObject(link.child, link.parent)
		}
	}

	subjVertices, subjParentTable, err := mia.GetSubjectArchitecture()

	if err != nil {
		t.Errorf("mia.GetSubjectArchitecture should work")
	}

	domainVertices, domainParentTable, err := mia.GetDomainArchitecture()

	if err != nil {
		t.Errorf("mia.GetDomainArchitecture should work")
	}

	objectVertices, objectParentTable, err := mia.GetObjectArchitecture()

	if err != nil {
		t.Errorf("mia.GetObjectArchitecture should work")
	}

	testEquity(t, subjects, subjectsLinks, model.ITEM_TYPE_SUBJ, subjVertices, subjParentTable)
	testEquity(t, domains, domainLinks, model.ITEM_TYPE_DOMAIN, domainVertices, domainParentTable)
	testEquity(t, objects, objectLinks, model.ITEM_TYPE_OBJ, objectVertices, objectParentTable)

	utils.PrintArchi(subjVertices, subjParentTable)
	utils.PrintArchi(domainVertices, domainParentTable)
	utils.PrintArchi(objectVertices, objectParentTable)
}
