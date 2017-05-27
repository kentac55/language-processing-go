package language_processing_go

import (
	"testing"
)

func TestP02(t *testing.T) {
	const i, o = "パタトクカシーー", "パトカー"
	if x := P01(i); x != o {
		t.Errorf("P01(%v) = %v, want %v", i, x, o)
	}
}
