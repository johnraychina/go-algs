package basic

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	//g := NewGraphFromFile(file)

	g := NewUndirectedGraph(13)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)

	g.AdjOf(0)
	g.Degree(0)
}

func TestPathsOf(t *testing.T) {

	g := NewUndirectedGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(0, 5)

	paths := DepthFirstPathsOf(g, 0)
	fmt.Println("EdgeTo[]")
	for i, v := range paths.edgeTo {
		fmt.Printf("%d|%d\n", i, v)
	}
	for i := 0; i < g.V(); i++ {
		fmt.Println("path from 0 to ", i, " is:", paths.PathTo(i))
	}
}

func TestBFS(t *testing.T) {

	g := NewUndirectedGraph(6)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(0, 5)

	paths := BreadthFirstPathsOf(g, 0)
	fmt.Println("EdgeTo[]")
	for i, v := range paths.edgeTo {
		fmt.Printf("%d|%d\n", i, v)
	}
	for i := 0; i < g.V(); i++ {
		fmt.Println("path from 0 to ", i, " is:", paths.PathTo(i))
	}
}

func TestConnectedComponent(t *testing.T) {

	g := NewUndirectedGraph(10)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(0, 5)

	g.AddEdge(6, 7)
	g.AddEdge(6, 8)

	cc := NewConnectedComponent(g)
	fmt.Println("count:", cc.count)
	fmt.Println("Id[v]")
	for v, cId := range cc.id {
		fmt.Printf("%d|%d\n", v, cId)
	}
}
