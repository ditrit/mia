package engine_test

import (
	"iam/core"
	"iam/core/model"
	"os"
	"testing"
)

//nolint: gochecknoglobals
var iam core.IAM

func TestMain(m *testing.M) {
	iam = core.NewIAM("test.db", true)

	ret := m.Run()

	os.Remove("test.db") //nolint: errcheck
	os.Exit(ret)
}

func TestGetUnknown(t *testing.T) {
	_, err := iam.GetSubject("alice")
	if err == nil {
		t.Errorf("get unknown shouldn't work")
	}
}

func TestGetAfterAdd(t *testing.T) {
	bobby, err := model.NewSubject("bobby")
	if err != nil {
		t.Errorf("should be okay")
	}

	err = iam.AddSubject(*bobby)

	if err != nil {
		t.Errorf("add should work %s", err.Error())
	}

	_, err = iam.GetSubject("bobby")
	if err != nil {
		t.Errorf("get unknown shouldn't work")
	}
}
