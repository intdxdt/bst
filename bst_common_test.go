package bst

import (
	"testing"
	"github.com/intdxdt/cmp"
	"github.com/franela/goblin"
)

func TestBSTUtil(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("BST-Node", func() {
		//var left  = NewBranch().AsLeft()
		//var right =    NewBranch().AsRight()
		var node *Node
		var array = []int{1, 2}
		nodes := make([]*Node, len(array))
		for i := range array {
			nodes[i] = NewNode(array[i], cmp.Int)
		}
		node, nodes = ShiftNode(nodes)
		g.Assert(node.Key).Equal(array[0])

		node, nodes = PopNode(nodes)
		g.Assert(node != nil).IsTrue()

		node, nodes = PopNode(nodes)
		g.Assert(node == nil).IsTrue()

		node, nodes = ShiftNode(nodes)
		g.Assert(node == nil).IsTrue()

	})
}