package nodetypes

type ConnectNode struct {
	ParentNodeIndex int
	IsOpen          bool
}

func NewNode() ConnectNode {
	return ConnectNode{}
}

func (c ConnectNode) Open() {
	c.IsOpen = true
}
