package mod_test

import (
	"os"
	"testing"

	. "github.com/izumin5210/gex/pkg/manager/mod"
)

func TestManagerImpl_Vendor(t *testing.T) {
	current := os.Getenv(GoFlagsEnv)
	defer func() {
		if err := os.Setenv(GoFlagsEnv, current); err != nil {
			t.Log(err)
			t.FailNow()
		}
	}()
	manager := NewManager(nil).(interface{ Vendor() bool })

	if err := os.Setenv(GoFlagsEnv, " -mod=vendor -flag=value "); err != nil {
		t.Log(err)
		t.FailNow()
	}
	if !manager.Vendor() {
		t.Log("vendor mode is expected")
		t.FailNow()
	}

	if err := os.Setenv(GoFlagsEnv, " -mod= -flag=value "); err != nil {
		t.Log(err)
		t.FailNow()
	}
	if manager.Vendor() {
		t.Log("vendor mode is not expected")
		t.FailNow()
	}
}
