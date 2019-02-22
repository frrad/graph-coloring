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

	colorMap, solveable := solveGraph(edges, 4)
	if !solveable {
		panic("unsat")
	}
	fmt.Println(colorMap)
}

func solveGraph(edges [][2]string, nColors uint8) (map[string]uint8, bool) {
	prob := smt.NewProb()
	colorMap := map[string]*smt.UInt8{}
	colorMax := prob.NewUInt8Const(nColors)

	for _, e := range edges {
		f, t := e[0], e[1]
		if _, ok := colorMap[f]; !ok {
			colorMap[f] = newColor(colorMax, prob)
		}
		if _, ok := colorMap[t]; !ok {
			colorMap[t] = newColor(colorMax, prob)
		}
		prob.Assert(colorMap[f].Neq(colorMap[t]))
	}

	solveable := prob.Solve(solver.Solve)
	if !solveable {
		return nil, false
	}

	ans := map[string]uint8{}
	for k, v := range colorMap {
		ans[k] = v.SolVal()
	}

	return ans, true
}

func newColor(colorMax *smt.UInt8, prob *smt.Problem) *smt.UInt8 {
	x := prob.NewUInt8()
	prob.Assert(x.Lt(colorMax))
	return x
}
