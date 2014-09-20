package util9

import (
	"testing"
)

func TestSysstat(t *testing.T) {
	sw, err := ReadSysstat()
	if err != nil {
		t.Fatal(err)
	}

	if len(sw) < 1 {
		t.Fatal("sysstat read returns nothing")
	}

	for _, s := range sw {
		t.Logf("%+v", s)
	}
}
