package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_New(t *testing.T) {
	node := newNode("test_node")
	assert.NotNil(t, node)
	assert.Equal(t, "test_node", node.Name())
}

func TestNode_String(t *testing.T) {
	node := newNode("test_node")
	other := newNode("test_node")

	assert.NotEqual(t, node, other)
	assert.Equal(t, node.Name(), other.Name())
	assert.NotEqual(t, node.String(), other.String())
}

func TestNode_isDescendantOf(t *testing.T) {
	parent := newParent("test_parent")
	node := newNode("test_child")

	err := parent.AppendChildNode(node)
	assert.Nil(t, err)
	assert.True(t, node.IsDescendantOf(parent))
	err = parent.RemoveChildNode(node)
	assert.Nil(t, err)
	assert.False(t, node.IsDescendantOf(parent))

	assert.False(t, newNode("test_node").IsDescendantOf(parent))
}

func TestNode_isDescendantOf_NilParent(t *testing.T) {
	node := newNode("test_child")
	assert.False(t, node.IsDescendantOf(nil))
}
