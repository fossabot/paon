package tree

import (
	"fmt"

	"github.com/negrel/debuggo/pkg/log"
)

// Node define an element in a Node tree.
type Node interface {
	fmt.Stringer

	ID() NodeID
	Name() string

	IsSame(Node) bool

	// Next sibling.
	NextNode() Node
	setNextNode(Node)

	// Previous sibling.
	PreviousNode() Node
	setPreviousNode(Node)

	// ParentNode is the direct parentNode of the Node.
	ParentNode() ParentNode
	setParentNode(ParentNode)

	// RootNode define the root of the Node tree.
	RootNode() Root

	IsDescendantOf(node Node) bool
}

var _ Node = &node{}

type node struct {
	name string
	id   NodeID

	next     Node
	previous Node
	parent   ParentNode
}

func NewNode(name string) Node {
	return newNode(name)
}

func newNode(name string) *node {
	log.Debugln("creating a", name, "node")

	return &node{
		name: name,
		id:   nodeID(),
	}
}

func (n *node) String() string {
	return fmt.Sprintf("%v-%v", n.name, n.id)
}

func (n *node) ID() NodeID {
	return n.id
}

func (n *node) Name() string {
	return n.name
}

func (n *node) IsSame(other Node) bool {
	return n.ID() == other.ID()
}

func (n *node) IsDescendantOf(parent Node) bool {
	if parent == nil {
		return false
	}

	var node Node = n
	for node != nil {
		if node.IsSame(parent) {
			return true
		}

		node = node.ParentNode()
	}

	return false
}

func (n *node) NextNode() Node {
	return n.next
}

func (n *node) setNextNode(next Node) {
	n.next = next
}

func (n *node) PreviousNode() Node {
	return n.previous
}

func (n *node) setPreviousNode(previous Node) {
	n.previous = previous
}

func (n *node) ParentNode() ParentNode {
	return n.parent
}

func (n *node) setParentNode(parent ParentNode) {
	n.parent = parent
}

func (n *node) RootNode() Root {
	if n.parent != nil {
		return n.parent.RootNode()
	}

	return nil
}
