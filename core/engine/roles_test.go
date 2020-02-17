package engine_test

import (
	"iam/core/model"
	"testing"
)

func searchRole(roles []model.Role, name string) bool {
	for _, elem := range roles {
		if elem.Name == name {
			return true
		}
	}

	return false
}

func areModelsEquals(roles1 []model.Role, roles2 []model.Role) bool {
	for _, elem1 := range roles1 {
		found := false

		for _, elem2 := range roles2 {
			if elem1 == elem2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

//nolint: gocyclo
func TestRole(t *testing.T) {
	const (
		nameRole = "ProductTester"
	)

	e := iam.AddRole("a")

	if e != nil {
		t.Errorf("adding role should have worked")
	}

	_ = iam.AddRole("b")
	_ = iam.AddRole("c")

	beforeRoles, err := iam.GetAllRoles()

	if err != nil {
		t.Errorf("should never happened")
	}

	if searchRole(beforeRoles, nameRole) {
		t.Errorf("this role shouldn't be in the iam")
	}

	errAdd := iam.AddRole(nameRole)

	if errAdd != nil {
		t.Errorf("adding role should have worked")
	}

	roles, err := iam.GetAllRoles()

	if err != nil {
		t.Errorf("should have worked")
	}

	if !searchRole(roles, nameRole) {
		t.Errorf("this role should be in the iam")
	}

	errRemove := iam.RemoveRole(nameRole)

	if errRemove != nil {
		t.Errorf("remove role should have worked")
	}

	afterRoles, err := iam.GetAllRoles()

	if err != nil {
		t.Errorf("should have worked")
	}

	if searchRole(afterRoles, nameRole) {
		t.Errorf("this role shouldn't be in the iam")
	}

	if !areModelsEquals(beforeRoles, afterRoles) {
		t.Errorf("beforeRoles and afterRoles should be the same")
	}
}
