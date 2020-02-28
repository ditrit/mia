package engine_test

import (
	"mia/core/constant"
	"testing"
)

func TestBasicEnforce(t *testing.T) {
	const (
		user1        = "user1"
		user2        = "user2"
		role         = "DEV"
		domain       = "gandalf"
		objGit       = "git"
		objGitPull   = "gitpull"
		objGitPush   = "gitpush"
		objGitCommit = "gitcommit"
	)

	_ = mia.AddSubjectToRoot(user1)
	_ = mia.AddSubjectToRoot(user2)
	_ = mia.AddDomainToRoot(domain)
	_ = mia.AddObjectToRoot(objGit)

	_ = mia.AddObject(objGitPull, objGit)
	_ = mia.AddObject(objGitPush, objGit)
	_ = mia.AddObject(objGitCommit, objGit)

	_ = mia.AddRole(role)
	_ = mia.AddAssignment(role, user1, domain)
	_ = mia.AddPermission(role, constant.ROOT_DOMAINS, objGit, constant.ACTION_EXECUTE, true)

	eff, err := mia.Enforce(user1, domain, objGitCommit, constant.ACTION_EXECUTE)

	if err != nil {
		t.Errorf("Something went wrong")
	}

	if !eff {
		t.Errorf("Wrong effect")
	}

	_ = mia.AddPermission(role, domain, objGitPull, constant.ACTION_EXECUTE, false)

	eff, err = mia.Enforce(user1, domain, objGitPull, constant.ACTION_EXECUTE)

	if err != nil {
		t.Errorf("Something went wrong")
	}

	if eff {
		t.Errorf("Wrong effect")
	}
}
