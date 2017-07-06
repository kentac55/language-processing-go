package exercise

import (
	"testing"
)

func TestP04(t *testing.T) {
	const i = "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	o := map[string]int{"Cl": 17, "N": 7, "Be": 4, "F": 9, "Mi": 12, "Na": 11, "Ne": 10, "Ar": 18, "B": 5, "Li": 3, "P": 15, "He": 2, "Si": 14, "C": 6, "H": 1, "Ca": 20, "K": 19, "Al": 13, "O": 8, "S": 16}
	x := P04(i)
	for k, v := range o {
		if x[k] != v {
			t.Errorf("P04(%v) = %v, want %v", i, x, o)
		}
	}
}
