package main

import (
	"fmt"

	solver "github.com/frrad/boolform/bfgosat"
	"github.com/frrad/boolform/smt"
)

func main() {
	edges := [][2]string{
		{"a", "b"},
		{"a", "c"},
		{"a", "e"},
		{"a", "f"},
		{"b", "d"},
		{"b", "f"},
		{"b", "g"},
		{"c", "d"},
		{"c", "e"},
		{"c", "f"},
		{"d", "f"},
		{"d", "g"},
		{"e", "g"},
	}

	prob := smt.NewProb()
	colorMap := map[string]*smt.UInt8{}
	colorMax := prob.NewUInt8Const(4)

	for _, e := range edges {
		f, t := e[0], e[1]
		if _, ok := colorMap[f]; !ok {
			x := prob.NewUInt8()
			prob.Assert(x.Lt(colorMax))
			colorMap[f] = x
		}
		if _, ok := colorMap[t]; !ok {
			x := prob.NewUInt8()
			prob.Assert(x.Lt(colorMax))
			colorMap[t] = x
		}
		prob.Assert(colorMap[f].Neq(colorMap[t]))
	}

	solveable := prob.Solve(solver.Solve)
	if !solveable {
		panic("unsat")
	}

	for k, v := range colorMap {
		fmt.Println(k, v.SolVal())
	}
}
