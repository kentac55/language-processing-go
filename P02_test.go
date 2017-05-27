package language_processing_go

import (
	"testing"
)

func TestP02(t *testing.T) {
	const i1, i2, o = "パトカー", "タクシー", "パタトクカシーー"
	if x, _ := P02(i1, i2); x != o {
		t.Errorf("P02(%v, %v) = %v, want %v", i1, i2, x, o)
	}
	const invalidStrA, invalidStrB = "パトカー", "バス"
	if _, e := P02(invalidStrA, invalidStrB); e == nil {
		t.Errorf("P02(%v, %v) expected throw error but nil", invalidStrA, invalidStrB)
	}
}
