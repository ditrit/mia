package engine_test

import (
	"iam/core/constant"
	"testing"
)

//TODO make some deepest tests
func TestPermission(t *testing.T) {
	const (
		roleName   string = "bendTester"
		domainName string = "Gandalf"
		objName    string = "Git Commit"
	)

	_ = iam.AddRole(roleName)
	_ = iam.AddDomainToRoot(domainName)
	_ = iam.AddObjectToRoot(objName)

	err := iam.RemovePermission(roleName, domainName, objName, constant.ACTION_EXECUTE)

	if err == nil {
		t.Errorf("remove permission that doesn't exist shouldn't work")
	}

	err = iam.AddPermission(roleName, domainName, objName, constant.ACTION_EXECUTE, true)

	if err != nil {
		t.Errorf("should have added the permission")
	}

	err = iam.AddPermission(roleName, domainName, objName, constant.ACTION_EXECUTE, true)

	if err == nil {
		t.Errorf("should have failed cause it's already exists")
	}

	err = iam.RemovePermission(roleName, domainName, objName, constant.ACTION_EXECUTE)

	if err != nil {
		t.Errorf("should have removed the permission")
	}
}
