package exercise

import (
	"reflect"
	"testing"
)

func TestP00(t *testing.T) {
	const i, o = "stressed", "desserts"
	if x := P00(i); x != o {
		t.Errorf("P00(%v) = %v, want %v", i, x, o)
	}
}

func TestP01(t *testing.T) {
	const i, o = "パタトクカシーー", "パトカー"
	if x := P01(i); x != o {
		t.Errorf("P01(%v) = %v, want %v", i, x, o)
	}
}

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

func TestP03(t *testing.T) {
	const i = "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	o := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}
	if x := P03(i); !reflect.DeepEqual(x, o) {
		t.Errorf("P03(%v) = %v, want %v", i, x, o)
	}
}

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
func TestP05(t *testing.T) {
	i := "I am an NLPer"
	o := Bigrams{
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

func TestP06(t *testing.T) {
	i1 := P05("paraparaparadise").Letter
	i2 := P05("paragraph").Letter
	o := P06Res{
		Set{
			"se": {}, "ar": {}, "is": {}, "pa": {}, "ap": {}, "di": {}, "ra": {}, "ad": {},
		},
		Set{
			"ar": {}, "pa": {}, "gr": {}, "ap": {}, "ra": {}, "ag": {}, "ph": {},
		},
		Set{
			"se": {}, "ar": {}, "is": {}, "pa": {}, "gr": {}, "ap": {}, "di": {}, "ra": {}, "ag": {}, "ph": {}, "ad": {},
		},
		Set{
			"ar": {}, "pa": {}, "ap": {}, "ra": {},
		},
		Set{
			"se": {}, "is": {}, "di": {}, "ad": {},
		},
	}

	x := P06(i1, i2)

	setA := Set{"a": {}, "b": {}, "c": {}}
	setB := Set{"c": {}, "b": {}, "a": {}}
	setC := Set{"a": {}, "b": {}}
	setD := Set{"z": {}, "y": {}, "x": {}}

	if !setA.equal(setB) {
		t.Errorf("equal method broken\n expect: %v == %v", setA, setB)
	}
	if setA.equal(setC) {
		t.Errorf("equal method broken\n expect: %v != %v", setA, setC)
	}
	if setA.equal(setD) {
		t.Errorf("equal method broken\n expect: %v != %v", setA, setD)
	}
	if !x.X.equal(o.X) {
		t.Errorf("expected: \n%v,\n actually: \n%v", o.X, x.X)
	}
	if !x.Y.equal(o.Y) {
		t.Errorf("expected: \n%v,\n actually: \n%v", o.Y, x.Y)
	}
	if !x.union.equal(o.union) {
		t.Errorf("expected: \n%v,\n actually: \n%v", o.union, x.union)
	}
	if !x.intersect.equal(o.intersect) {
		t.Errorf("expected: \n%v,\n actually: \n%v", o.intersect, x.intersect)
	}
	if !x.diff.equal(o.diff) {
		t.Errorf("expected: \n%v,\n actually: \n%v", o.diff, x.diff)
	}
	if x.X.Exists("se") != true {
		t.Errorf("X should have %v key", "se")
	}
	if x.Y.Exists("se") != false {
		t.Errorf("Y should not have %v key", "se")
	}
}

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

func TestP08(t *testing.T) {
	const i = "testTEST<>?_"
	if P08(P08(i)) != i {
		t.Errorf("P08(P08()) should return same str but different")
	}
}

func TestP09(t *testing.T) {
	const i = "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind ."
	const o = "I c'nlduot beielve that I cuold aualctly unaserntdd what I was ridneag : the phmnenaoel peowr of the huamn mind ."
	const seed = 1
	if x := P09(i, seed); x != o {
		t.Errorf("P09() may be broken... \ni: %v\ne: %v\no: %v", i, o, x)
	}
}
