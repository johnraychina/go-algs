package graph

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestPathTo(t *testing.T) {
	// build graph
	g := buildGraph()

	sp := NewShortestPaths(g, 0)
	for i := 0; i < g.V(); i++ {
		fmt.Printf("[%d]%v:, distance:%f, path:%v\n", i, sp.HasPathTo(i), sp.DistTo(i), fmtPath(sp.PathTo(i)))
	}
}

func fmtPath(path []*DirectedEdge) string {
	var out bytes.Buffer
	for i := len(path) - 1; i >= 0; i-- {
		e := path[i]
		out.WriteString(strconv.Itoa(e.v))
		out.WriteString("->")
		out.WriteString(strconv.Itoa(e.w))
		out.WriteString(",")
		//out.WriteString(":")
		//out.WriteString(strconv.FormatFloat(e.weight, 'g', 2, 64))
		//out.WriteString(fmt.Sprint(e.weight))
	}
	return out.String()
}

func buildGraph() *EdgeWeightedDiGraph {
	//edges := []*DirectedEdge{
	//	{0, 7, 0.16},
	//	{0, 4, 0.38},
	//	{0, 2, 0.26},
	//	{2, 3, 0.17},
	//	{2, 7, 0.34},
	//	{1, 7, 0.19},
	//	{1, 3, 0.29},
	//	{1, 5, 0.32},
	//	{1, 2, 0.36},
	//	{4, 5, 0.35},
	//	{4, 7, 0.37},
	//	{3, 6, 0.52},
	//	{5, 7, 0.28},
	//	{6, 2, 0.40},
	//	{6, 0, 0.58},
	//	{6, 4, 0.93}}
	edges := []*DirectedEdge{
		{0, 1, 0.5},
		{1, 2, 0.5},
		{0, 2, 1.5}}
	g := NewEdgeWeightedDiGraph(3)
	for _, e := range edges {
		g.AddDirectedEdge(e)
	}
	return g
}
