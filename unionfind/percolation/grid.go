package percolation

import (
	"cs/algo-ds/unionfind/datastructures"
	"math/rand"
	"time"
)

type UnionFindDsType int

const (
	ArrayQuickFind  UnionFindDsType = 1
	ArrayQuickUnion UnionFindDsType = 2
)

// datastructures interface
type UnionFind interface {
	IsConnected(p int, q int) bool
	Connect(p int, q int)
}

// create new Grid using UnionFind
type NewUnionFindGridFunc func(height int, width int) UnionFindGrid

type UnionFindGrid struct {
	uf     UnionFind
	height int
	width  int
	total  int
}

// interestingly golang ways of creating an object can return nonpointer or pointer
// new(type) = *type
// make(type) = type
// := type{} = type
// := &type{} = *type
func NewUnionFindGrid(unionFindType UnionFindDsType) NewUnionFindGridFunc {
	return func(height int, width int) UnionFindGrid {
		// 0 = top node
		// height * width + 1 = bottom node
		total := height*width + 2

		// can i pass this in instead?
		newUf := createNewUnionFind(unionFindType, total)

		grid := UnionFindGrid{
			uf:     newUf,
			height: height,
			width:  width,
			total:  total,
		}

		// connect top row to top node where top node = 0
		for i := 1; i < width+1; i++ {
			grid.connect(0, i)
		}

		// connect bottom row to bottom node where bottom node = height * width + 1
		bottomIndex := total - 1
		for i := total - width - 1; i < total-1; i++ {
			grid.connect(bottomIndex, i)
		}

		return grid
	}
}

func createNewUnionFind(unionFindType UnionFindDsType, total int) UnionFind {
	switch unionFindType {
	case ArrayQuickFind:
		return datastructures.NewQuickFind(total)
	default:
		return datastructures.NewQuickUnion(total)
	}
}

func (g UnionFindGrid) getRandomIndex() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(g.total)
}

func (g UnionFindGrid) connect(p int, q int) {
	g.uf.Connect(p, q)
}

func (g UnionFindGrid) isConnected(p int, q int) bool {
	return g.uf.IsConnected(p, q)
}

func (g UnionFindGrid) processOneTick() bool {
	p := g.getRandomIndex()
	q := g.getRandomIndex()
	for q == p {
		q = g.getRandomIndex()
	}
	g.connect(p, q)

	return g.isConnected(0, g.total-1)
}

func (g UnionFindGrid) Process() int {
	iterations := 1
	for !g.processOneTick() {
		iterations++
	}
	return iterations
}
