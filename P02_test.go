package language_processing_go

import (
	"testing"
)

func TestP02(t *testing.T) {
	const i, o = "パタトクカシーー", "パトカー"
	if x := P02(i); x != o {
		t.Errorf("P02(%v) = %v, want %v", i, x, o)
	}
}