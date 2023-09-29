package graph

import (
	"fmt"
	"testing"
)

func TestTopoloySort(t *testing.T) {
	g := NewDiGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(0, 5)

	dfsOrder := NewDepthFirstOrder(g)
	for !dfsOrder.reversePost.IsEmpty() {
		fmt.Println(dfsOrder.reversePost.Pop())
	}
}
