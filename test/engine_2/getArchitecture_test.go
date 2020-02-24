package engine_test

import (
	"testing"

	"iam/core/constant"
	"iam/core/model"
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
		"all",
	}

	subjectsLinks := []link{
		{parent: "DevTeam", child: "Barney"},
		{parent: "DevTeam", child: "Charlie"},
		{parent: "DevTeam", child: "David"},
		{parent: "Administration", child: "Angel"},
		{parent: "all", child: "Administration"},
		{parent: "all", child: "DevTeam"},
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
		{parent: "DevTeam", child: "Gandalf"},
		{parent: "ORNESS", child: "Commercial"},
	}

	objects := []string{
		"GitLab",
		"GitCommit",
	}

	objectLinks := []link{
		{parent: "GitLab", child: "GitCommit"},
	}

	for _, subj := range subjects {
		_ = iam.AddSubject(subj)
	}

	for _, link := range subjectsLinks {
		_ = iam.AddSubjectLink(link.parent, link.child)
	}

	for _, domain := range domains {
		_ = iam.AddDomain(domain)
	}

	for _, link := range domainLinks {
		_ = iam.AddDomainLink(link.parent, link.child)
	}

	for _, object := range objects {
		_ = iam.AddObject(object)
	}

	for _, link := range objectLinks {
		_ = iam.AddObjectLink(link.parent, link.child)
	}

	subjVertices, subjParentTable, err := iam.GetSubjectArchitecture()

	if err != nil {
		t.Errorf("iam.GetSubjectArchitecture should work")
	}

	domainVertices, domainParentTable, err := iam.GetDomainArchitecture()

	if err != nil {
		t.Errorf("iam.GetDomainArchitecture should work")
	}

	objectVertices, objectParentTable, err := iam.GetObjectArchitecture()

	if err != nil {
		t.Errorf("iam.GetObjectArchitecture should work")
	}

	testEquity(t, subjects, subjectsLinks, model.ITEM_TYPE_SUBJ, subjVertices, subjParentTable)
	testEquity(t, domains, domainLinks, model.ITEM_TYPE_DOMAIN, domainVertices, domainParentTable)
	testEquity(t, objects, objectLinks, model.ITEM_TYPE_OBJ, objectVertices, objectParentTable)
}
