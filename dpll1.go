package main

import (
	"fmt"
)

type literal struct {
	name string
	sig  bool
}

type clause struct {
	ls []literal //literal set
}

type cnf struct {
	cs []clause //clause set
	ls []literal
}

func main() {
	l1 := literal{"l1", true}
	l2 := literal{"l2", true}
	l3 := literal{"l1", false}
	l4 := literal{"l2", false}

	c1 := clause{l1}
	c2 := clause{l2}
	c3 := clause{l3}
	c4 := clause{l4}
	c5 := clause{l1, l2}
	c6 := clause{l3, l4}

	f1 := cnf{{c1, c2}, {}}
	f2 := cnf{{c1, c3}, {}}
	f3 := cnf{{c3, c4, c5}, {}}

	fmt.Println("done.")
}

func dpll(f cnf) bool {
	//consistency checking
	if isConsistent(f) {
		return true
	}
	//empty clasue checking
	if hasEmptyClause(f) {
		return false
	}

	var ucs []literal
	for uc := range f.cs {
		if len(uc.ls) == 1 {
			ucs += uc.ls[0]
		}
	}

	//////here goes on...
	for len(ucs) != 0 {
		for l := range ucs {
			f, e := unitPropogate(l, f)
			if e != nil {
				fmt.Println{"Error:" + e}
			}
		}
		////still not ...
		for uc := range f.cs {
			if len(uc.ls[0] == 1 {
				ucs += uc.ls[0]
			}
		}
	}

}

func unitPropogate(l literal, f cnf) cnf {
	var newcs cnf
	return newcs
}

func pureLiteralAssign(l literal, f cnf) cnf {
	var newcs cnf
	return newcs
}

func isConsistent(f cnf) bool {
	lcs := len(f.cs)
	if lcs > 0 {
		return false
	}
	// for i := 0; i < lcs; i++ {
	// 	if len(f.cs[i].ls) > 1 {
	// 		return false
	// 	}
	// }
	lls := len(f.ls)
	for i := 0; i < lls-1; i++ {
		for j := i + 1; j < lls; j++ {
			if (f.ls[i].sig != f.ls[j].sig) && (f.ls[i].name == f.ls[j].name) {
				return false
			}
		}
	}
	return true
}

func hasEmptyClause(f cnf) bool {
	for c := range f.cs {
		if len(c.ls) == 0 {
			return true
		}
	}
	return false
}
