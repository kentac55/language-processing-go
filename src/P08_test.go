package exercise

import (
	"testing"
)

func TestP08(t *testing.T) {
	const i = "testTEST<>?_"
	if P08(P08(i)) != i {
		t.Errorf("P08(P08()) should return same str but different")
	}
}
