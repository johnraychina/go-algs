package basic

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	//g := NewGraphFromFile(file)

	g := NewGraph(13)
	g.addEdge(0, 1)
	g.addEdge(0, 2)
	g.addEdge(1, 2)
	g.addEdge(2, 3)

	g.adjOf(0)
	g.degree()
}

func TestPathsOf(t *testing.T) {

	g := NewGraphFromFile(file)
	paths := PathsOf(g, 0)
	for i := 0; i < g.V(); i++ {
		if paths.hasPathTo(i) {
			fmt.Println(i)
		}
	}
}
