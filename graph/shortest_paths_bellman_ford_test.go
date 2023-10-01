package graph

import (
	"fmt"
	"testing"
)

func TestFindCycle(t *testing.T) {

	edges := []*DirectedEdge{
		{0, 1, 1},
		{1, 2, 1},
		{2, 0, -3}}
	g := NewEdgeWeightedDiGraph(3)
	for _, e := range edges {
		g.AddDirectedEdge(e)
	}

	sp := NewBellmanFordSP(g, 0)

	fmt.Println("has cycle?  ", sp.hasNegativeCycle())
	if sp.hasNegativeCycle() {
		for !sp.cycle.IsEmpty() {
			fmt.Println(sp.cycle.Pop())
		}
	}
}
