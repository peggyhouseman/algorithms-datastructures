package unionfind

import "cs/algo-ds/utils"

type ArrayQuickUnion struct {
	nodes []int
	sizes []int
}

func NewQuickUnion(length int) ArrayQuickUnion {
	a := ArrayQuickUnion{
		nodes: utils.NewIntArray(length),
		sizes: utils.NewIntArray(length),
	}

	return a
}

func (a ArrayQuickUnion) IsConnected(p int, q int) bool {
	pRoot := a.findRoot(p)
	qRoot := a.findRoot(q)
	return pRoot == qRoot
}

func (a ArrayQuickUnion) Connect(p int, q int) {
	pRoot := a.findRoot(p)
	qRoot := a.findRoot(q)

	pTreeSize := a.sizes[pRoot]
	qTreeSize := a.sizes[qRoot]

	if pTreeSize >= qTreeSize {
		a.nodes[qRoot] = pRoot
		a.sizes[pRoot] += qTreeSize
	} else {
		a.nodes[pRoot] = qRoot
		a.sizes[qRoot] += pTreeSize
	}
}

func (a ArrayQuickUnion) findRoot(p int) int {
	pIndex := a.nodes[p]
	for a.nodes[pIndex] != pIndex {
		pIndex = a.nodes[pIndex]
	}
	return pIndex
}
