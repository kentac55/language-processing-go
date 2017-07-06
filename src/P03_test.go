package exercise

import (
	"reflect"
	"testing"
)

func TestP03(t *testing.T) {
	const i = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	o := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}
	if x := P03(i); !reflect.DeepEqual(x, o) {
		t.Errorf("P03(%v) = %v, want %v", i, x, o)
	}
}
