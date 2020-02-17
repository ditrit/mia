package engine_test

import (
	"testing"
)

func searchRole(roles []string, name string) bool {
	for _, elem := range roles {
		if elem == name {
			return true
		}
	}

	return false
}

func areModelsEquals(roles1 []string, roles2 []string) bool {
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

func TestRoleGetAdd(t *testing.T) {
	const roleName = "coucou"

	found, err := iam.IsRoleExists(roleName)
	_, err2 := iam.GetRole(roleName)

	if err != nil {
		t.Errorf("shoud have worked")
	}

	if err2 == nil {
		t.Errorf("should have failed")
	}

	if found {
		t.Errorf("coucou shouldn't be in the iam")
	}

	e := iam.AddRole(roleName)

	if e != nil {
		t.Errorf("adding role should have worked")
	}

	found, err = iam.IsRoleExists(roleName)
	role, err2 := iam.GetRole(roleName)

	if err != nil || err2 != nil {
		t.Errorf("shoud have worked")
	}

	if role.Name != roleName {
		t.Errorf("names should be the same")
	}

	if !found {
		t.Errorf("coucou should be in the iam")
	}
}

func TestRoleRemove(t *testing.T) {
	const (
		nameRole = "ProductTester"
	)

	_ = iam.AddRole("a")
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
