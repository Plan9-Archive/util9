package util9

import (
	"testing"
)

func TestSwap(t *testing.T) {
	sw, err := ReadSwap()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(sw)
}
