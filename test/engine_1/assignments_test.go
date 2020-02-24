package engine_test

import (
	"testing"
)

//TODO make some deepest tests
func TestAssignment(t *testing.T) {
	const (
		roleName   string = "bendTester"
		subjName   string = "Carlos"
		domainName string = "Gandalf"
	)

	_ = iam.AddRole(roleName)
	_ = iam.AddSubjectToRoot(subjName)
	_ = iam.AddDomainToRoot(domainName)

	err := iam.RemoveAssignment(roleName, subjName, domainName)

	if err == nil {
		t.Errorf("remove assignment that doesn't exist shouldn't work")
	}

	err = iam.AddAssignment(roleName, subjName, domainName)

	if err != nil {
		t.Errorf("should have added the assignment")
	}

	err = iam.AddAssignment(roleName, subjName, domainName)

	if err == nil {
		t.Errorf("should have failed cause it's already exists")
	}

	err = iam.RemoveAssignment(roleName, subjName, domainName)

	if err != nil {
		t.Errorf("should have removed the assignment")
	}
}
