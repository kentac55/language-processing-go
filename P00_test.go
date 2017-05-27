package language_processing_go

import (
	"testing"
)

func TestP01(t *testing.T) {
	const i, o = "stressed", "desserts"
	if x := P00(i); x != o {
		t.Errorf("P00(%v) = %v, want %v", i, x, o)
	}
}
