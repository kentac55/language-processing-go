package exercise

import (
	"reflect"
	"testing"
)

func TestP05(t *testing.T) {
	i := "I am an NLPer"
	o := biGram{
		[][]string{
			{"I", "am"}, {"am", "an"}, {"an", "NLPer"},
		},
		[]string{
			"am", "an", "NL", "LP", "Pe", "er",
		},
	}
	if x := P05(i); !reflect.DeepEqual(x, o) {
		t.Errorf("P05(%v) = %v, want %v", i, x, o)
	}

}
