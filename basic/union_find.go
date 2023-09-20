package basic

// Union-Find的应用场景
//・Percolation.
//・Games (Go, Hex).
//・Least common ancestor.
//・Equivalence of finite state automata.
//・Hoshen-Kopelman algorithm in physics.
//・Hinley-Milner polymorphic type inference.
//・Kruskal's minimum spanning tree algorithm.
//・Compiling equivalence statements in Fortran.
//・Morphological attribute openings and closings.
//・Matlab's bwlabel() function in image processing.

type UF interface {
	Union(p, q int)          //add connection between p and q
	Connected(p, q int) bool //are p and q in the same component?
	Find(p int) int          // component identifier for p (0 to N – 1)
	Count() int              //component identifier for p (0 to N – 1)
}

type QuickUnionUF struct {
	id []int
}

func NewQuickUnionUF(N int) *QuickUnionUF {
	id := make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i // self as root
	}
	return &QuickUnionUF{id: id}
}

func (u *QuickUnionUF) root(p int) int {
	for u.id[p] != p {
		u.id[p] = u.id[u.id[p]] // Make every other node in path point to its grandparent (thereby halving path length)
		p = u.id[p]             // get parent node id
	}
	return p
}

func (u *QuickUnionUF) Connected(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u *QuickUnionUF) Find(p int) int {
	return u.id[p]
}

// Count number of components
func (u *QuickUnionUF) Count() int {
	// todo
	return 0
}

func (u *QuickUnionUF) Union(p, q int) {
	i := u.root(p)
	j := u.root(q)
	u.id[i] = j // merge trees
}

type QuickFindUF struct {
	id []int
}

func (u *QuickFindUF) Union(p, q int) {
	pId := u.id[p]
	qId := u.id[q]
	for i, id := range u.id {
		if id == pId {
			u.id[i] = qId
		}
	}
}
func (u *QuickFindUF) Connected(p, q int) bool {
	return u.id[p] == u.id[q]
}
func (u *QuickFindUF) Find(p int) int {
	return u.id[p]
}

// Count number of components
func (u *QuickFindUF) Count() int {
	// todo
	return 0
}
