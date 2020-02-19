package engine_test

import (
	"iam/core"
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
