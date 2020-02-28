package engine_test

import (
	"mia/core/constant"
	"testing"
)

//TODO make some deepest tests
func TestPermission(t *testing.T) {
	const (
		roleName   string = "bendTester"
		domainName string = "Gandalf"
		objName    string = "Git Commit"
	)

	_ = mia.AddRole(roleName)
	_ = mia.AddDomainToRoot(domainName)
	_ = mia.AddObjectToRoot(objName)

	err := mia.RemovePermission(roleName, domainName, objName, constant.ACTION_EXECUTE)

	if err == nil {
		t.Errorf("remove permission that doesn't exist shouldn't work")
	}

	err = mia.AddPermission(roleName, domainName, objName, constant.ACTION_EXECUTE, true)

	if err != nil {
		t.Errorf("should have added the permission")
	}

	err = mia.AddPermission(roleName, domainName, objName, constant.ACTION_EXECUTE, true)

	if err == nil {
		t.Errorf("should have failed cause it's already exists")
	}

	err = mia.RemovePermission(roleName, domainName, objName, constant.ACTION_EXECUTE)

	if err != nil {
		t.Errorf("should have removed the permission")
	}
}
