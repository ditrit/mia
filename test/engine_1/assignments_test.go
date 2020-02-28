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

	_ = mia.AddRole(roleName)
	_ = mia.AddSubjectToRoot(subjName)
	_ = mia.AddDomainToRoot(domainName)

	err := mia.RemoveAssignment(roleName, subjName, domainName)

	if err == nil {
		t.Errorf("remove assignment that doesn't exist shouldn't work")
	}

	err = mia.AddAssignment(roleName, subjName, domainName)

	if err != nil {
		t.Errorf("should have added the assignment")
	}

	err = mia.AddAssignment(roleName, subjName, domainName)

	if err == nil {
		t.Errorf("should have failed cause it's already exists")
	}

	err = mia.RemoveAssignment(roleName, subjName, domainName)

	if err != nil {
		t.Errorf("should have removed the assignment")
	}
}
