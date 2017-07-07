package exercise

import (
	"reflect"
	"testing"
)

func TestP05(t *testing.T) {
	i := "I am an NLPer"
	o := biGram{
		[]wordBiGram{
			{"I", "am"}, {"am", "an"}, {"an", "NLPer"},
		},
		[]letterBiGram{
			{"I "}, {" a"}, {"am"}, {"m "}, {" a"}, {"an"},
			{"n "}, {" N"}, {"NL"}, {"LP"}, {"Pe"}, {"er"},
		},
	}
	if x := P05(i); !reflect.DeepEqual(x, o) {
		t.Errorf("P03(%v) = %v, want %v", i, x, o)
	}

}
