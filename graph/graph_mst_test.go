package graph

import (
	"fmt"
	"testing"
)

func createEdgeWeightedGraph() *EdgeWeightedGraph {

	//0-7  0.16
	//2-3  0.17
	//1-7  0.19
	//0-2  0.26
	//5-7  0.28
	//1-3  0.29
	//1-5  0.32
	//2-7  0.34
	//4-5  0.35
	//1-2  0.36
	//4-7  0.37
	//0-4  0.38
	//6-2  0.40
	//3-6  0.52
	//6-0  0.58
	//6-4  0.93
	edges := []*Edge{
		{0, 7, 0.16},
		{2, 3, 0.17},
		{1, 7, 0.19},
		{0, 2, 0.26},
		{5, 7, 0.28},
		{1, 3, 0.29},
		{1, 5, 0.32},
		{2, 7, 0.34},
		{4, 5, 0.35},
		{1, 2, 0.36},
		{4, 7, 0.37},
		{0, 4, 0.38},
		{6, 2, 0.40},
		{3, 6, 0.52},
		{6, 0, 0.58},
		{6, 4, 0.93}}
	g := NewEdgeWeightedGraph(8)
	for _, e := range edges {
		g.AddEdge(e)
	}
	return g
}

func TestKruskalMST(t *testing.T) {
	g := createEdgeWeightedGraph()
	k := NewKruskalMST(g)
	for !k.mst.IsEmpty() {
		fmt.Println(k.mst.Dequeue())
	}
}

func TestPrimMST(t *testing.T) {
	g := createEdgeWeightedGraph()
	k := NewPrimMST(g)
	for !k.mst.IsEmpty() {
		fmt.Println(k.mst.Dequeue())
	}
	//&{0 7 0.16}
	//&{1 7 0.19}
	//&{0 2 0.26}
	//&{2 3 0.17}
	//&{5 7 0.28}
	//&{4 5 0.35}
	//&{6 2 0.4}
}
