package exercise

import (
	"testing"
)

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
