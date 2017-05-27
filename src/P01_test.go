package exercise

import (
	"testing"
)

func TestP01(t *testing.T) {
	const i, o = "パタトクカシーー", "パトカー"
	if x := P01(i); x != o {
		t.Errorf("P01(%v) = %v, want %v", i, x, o)
	}
}