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
}

func main() {
	fmt.Println("done.")
}

func dpll(cs []clause) bool {

	return true
}

func unitPropogate(l literal, cs cnf) cnf {
	var newcs cnf
	return newcs
}

func pureLiteralAssign(l literal, cs cnf) cnf {
	var newcs cnf
	return newcs
}

func isConsistent(cs []clause) bool {

	return true
}
