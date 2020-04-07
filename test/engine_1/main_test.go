package engine_test

import (
	"mia/core"
	"os"
	"testing"
)

//nolint: gochecknoglobals
var mia core.MIA

func TestMain(m *testing.M) {
	mia = core.NewMIA("test.db", true)

	ret := m.Run()

	os.Remove("test.db") //nolint: errcheck
	os.Exit(ret)
}
