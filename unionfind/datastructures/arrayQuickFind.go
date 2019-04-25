// to determine if nodes are connected, check if the node is the same value
// to connect the node, update the node and all other related nodes to the value of the connected node
// this is a quick find since getting the index value of an array is O(1)
// connecting is O(n) for 1 node.  for m nodes O(mn)

package datastructures

type ArrayQuickFind struct {
	nodes []int
}

func NewQuickFind(length int) ArrayQuickFind {
	a := ArrayQuickFind{
		nodes: make([]int, length),
	}

	for i := 0; i < length; i++ {
		a.nodes[i] = i
	}

	return a
}

func (a ArrayQuickFind) IsConnected(p int, q int) bool {
	var pVal = a.nodes[p]
	var qVal = a.nodes[q]
	return pVal == qVal
}

func (a ArrayQuickFind) Connect(p int, q int) {
	var pVal = a.nodes[p]
	var qVal = a.nodes[q]

	for i := 0; i < len(a.nodes); i++ {
		if a.nodes[i] == pVal {
			a.nodes[i] = qVal
		}
	}
}
