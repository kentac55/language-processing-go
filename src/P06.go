package exercise

type Set map[string]struct{}

type SetFunc interface {
	Union(that Set) Set
	Intersect(that Set) Set
	Diff(that Set) Set
	Exists(that Set) bool
	Equal(that Set) bool
}

type P06Res struct {
	X         Set
	Y         Set
	union     Set
	intersect Set
	diff      Set
}

func NewSet(s []string) Set {
	set := Set{}
	if len(s) > 0 {
		for _, i := range s {
			set[i] = struct{}{}
		}
	}
	return set
}

func (s Set) Union(that Set) Set {
	set := Set{}
	for k := range s {
		set[k] = struct{}{}
	}
	for k := range that {
		set[k] = struct{}{}
	}
	return set
}

func (s Set) Intersect(that Set) Set {
	set := Set{}
	for a := range s {
		for b := range that {
			if a == b {
				set[a] = struct{}{}
			}
		}
	}
	return set
}

func (s Set) Diff(that Set) Set {
	set := Set{}
	for a := range s {
		set[a] = struct{}{}
		for b := range that {
			if a == b {
				delete(set, b)
			}
		}
	}
	return set
}

func (s Set) Exists(key string) bool {
	for k := range s {
		if k == key {
			return true
		}
	}
	return false
}

func (s Set) equal(that Set) bool {
	if len(s) != len(that) {
		return false
	}
	for a := range s {
		if _, ok := that[a]; !ok {
			return false
		}
	}
	return true
}

func P06(s1 []string, s2 []string) P06Res {
	X := NewSet(s1)
	Y := NewSet(s2)
	union := X.Union(Y)
	intersect := X.Intersect(Y)
	diff := X.Diff(Y)
	return P06Res{X, Y, union, intersect, diff}
}
