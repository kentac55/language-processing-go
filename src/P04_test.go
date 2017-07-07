package exercise

import (
	"testing"
)

func TestP04(t *testing.T) {
	const i = "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	o := map[string]int{"H": 1, "He": 2, "Li": 3, "Be": 4, "B": 5, "C": 6, "N": 7, "O": 8, "F": 9, "Ne": 10, "Na": 11, "Mi": 12, "Al": 13, "Si": 14, "P": 15, "S": 16, "Cl": 17, "Ar": 18, "K": 19, "Ca": 20}
	x := P04(i)
	if len(x) != len(o) {
		t.Errorf("expected length: %v, actually: %v\n%v", len(o), len(x), x)
	}
	for k, v := range o {
		if x[k] != v {
			t.Errorf("expected: {%v, %v}, actually: {%v, %v)", k, v, k, x[k])
		}
	}
}
