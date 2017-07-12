package exercise

import (
	"testing"
)

func TestP07(t *testing.T) {
	const (
		x = 12
		y = "気温"
		z = 22.4
		o = "12時の気温は22.4"
	)
	if res := P07(x, y, z); res != o {
		t.Errorf("P07(%v, %v, %v) = %v, want %v", x, y, z, res, o)
	}
}
