package regex

import (
	"go-algs/basic"
	"go-algs/graph"
)

type NFA struct {
	re string // regex
	M  int    // length of regex
	G  *graph.DiGraph
}

func (nfa *NFA) Recognize(text string) bool {
	states := make(map[int]bool) // current states of NFA
	epsilon := graph.NewDirectedDFS(nfa.G, []int{0})

	// initial states that is reachable from state 0
	for v := 0; v < nfa.G.V(); v++ {
		if epsilon.Marked(v) {
			states[v] = true
		}
	}

	// 计算text[i+1]匹配NFA能达到的所有状态
	for i := range text {
		match := make([]int, 0)
		for v := range states {
			// todo 这里想想为什么？
			if v < nfa.M {
				if nfa.re[v] == text[i] || nfa.re[v] == '.' {
					match = append(match, v+1)
				}
			}
		}

		// update states
		states = make(map[int]bool)
		epsilon = graph.NewDirectedDFS(nfa.G, match)
		for v := 0; v < nfa.G.V(); v++ {
			if epsilon.Marked(v) {
				states[v] = true
			}
		}
	}

	for v := range states {
		if v == nfa.M {
			return true
		}
	}
	return false
}

// NewNFA re - regex
func NewNFA(re string) *NFA {
	M := len(re)
	G := graph.NewDiGraph(M + 1) // M 个字符 + 1个accept status

	// 构造epsilon graph
	// 需要用到stack解决 or匹配和*匹配的跳线问题
	ops := basic.NewStack[int]()
	for i := 0; i < M; i++ {
		lp := i // left parenthesis

		// match (|)
		c := re[i]
		if c == '(' || c == '|' { //后续遇到)，需要匹配(,|
			ops.Push(i)
		} else if re[i] == ')' {
			or, err := ops.Pop()
			if err != nil {
				panic(err)
			}
			if re[or] == '|' {
				lp, _ = ops.Pop() // pop out left parenthesis index
				G.AddEdge(lp, or+1)
				G.AddEdge(or, i)
			} else {
				lp = or
			}
		}

		// match *, 向右多看一个字符
		if i+1 < M && re[i+1] == '*' {
			G.AddEdge(lp, i+1)
			G.AddEdge(i+1, lp)
		}

		// epsilon转换 1个字符
		if re[i] == '(' || re[i] == '*' || re[i] == ')' {
			G.AddEdge(i, i+1)
		}
	}

	return &NFA{re: re, M: len(re), G: G}
}
