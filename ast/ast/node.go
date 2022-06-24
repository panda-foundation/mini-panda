package ast

type NodeBase struct {
	position int
}

func (n *NodeBase) GetPosition() int {
	return n.position
}

func (n *NodeBase) SetPosition(position int) {
	n.position = position
}
