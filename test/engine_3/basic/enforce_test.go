package engine_test

import (
	"iam/core/constant"
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

	_ = iam.AddSubject(user1)
	_ = iam.AddSubject(user2)
	_ = iam.AddDomain(domain)
	_ = iam.AddObject(objGit)
	_ = iam.AddObject(objGitPull)
	_ = iam.AddObject(objGitPush)
	_ = iam.AddObject(objGitCommit)

	_ = iam.AddDomainLink(constant.ROOT_DOMAIN, domain)
	_ = iam.AddObjectLink(objGit, objGitPull)
	_ = iam.AddObjectLink(objGit, objGitPush)
	_ = iam.AddObjectLink(objGit, objGitCommit)

	_ = iam.AddRole(role)
	_ = iam.AddAssignment(role, user1, domain)
	_ = iam.AddPermission(role, constant.ROOT_DOMAIN, objGit, constant.ACTION_EXECUTE, true)

	eff, err := iam.Enforce(user1, domain, objGitCommit, constant.ACTION_EXECUTE)

	if err != nil {
		t.Errorf("Something went wrong")
	}

	if !eff {
		t.Errorf("Wrong effect")
	}

	_ = iam.AddPermission(role, domain, objGitPull, constant.ACTION_EXECUTE, false)

	eff, err = iam.Enforce(user1, domain, objGitPull, constant.ACTION_EXECUTE)

	if err != nil {
		t.Errorf("Something went wrong")
	}

	if eff {
		t.Errorf("Wrong effect")
	}
}
