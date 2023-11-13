package bst

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/cmp"
	"testing"
)

func TestBSTNode(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("BST-Node", func() {
		var left = NewBranch().AsLeft()
		var right = NewBranch().AsRight()
		var neutral = NewBranch().AsNeutral()

		var array = []int{1, 2, 4, 0}
		datanodes := make([]*Node, len(array))
		for i := range array {
			datanodes[i] = NewNode(array[i], cmp.Int)
		}
		var node1 = datanodes[0]
		var node2 = datanodes[1]
		var node4 = datanodes[2]
		var node0 = datanodes[3]
		var null = NewNode(nil, cmp.Int)
		node1.SetNode(node2, neutral) //nothing should happen
		g.Assert(node1.Left == nil).IsTrue()
		g.Assert(node1.Right == nil).IsTrue()

		//compare node and item equality
		var node4_clone = NewNode(4, cmp.Int)
		g.Assert(node4.Compare(nil) == 1).IsTrue()
		g.Assert(node4.keyCmp(node4.Key, 4) == 0).IsTrue()
		g.Assert(node4.keyCmp(node4.Key, node4_clone.Key) == 0).IsTrue()

		//new node parent is nil
		g.Assert(null.Key == nil).IsTrue()
		g.Assert(null.Left == nil).IsTrue()
		g.Assert(null.Right == nil).IsTrue()
		g.Assert(null.Parent == nil).IsTrue()
		g.Assert(node1.Parent == nil).IsTrue()

		//set child parent nodes node1 --> node2
		Ptr(node1, node2, right)
		Ptr(node2, node4, right)
		Ptr(node1, node0, left)
		/**
		   1
		   / \
		   0  2
		   /\ /\
			4
			/\
		*/
		//--------------------------------------
		g.Assert(node2.Parent == node1).IsTrue() // 2s parent is 1
		g.Assert(node4.Parent == node2).IsTrue() // 4s parent is 1
		g.Assert(node1.Right == node2).IsTrue()  // 1s right is 2
		g.Assert(SizeOfNode(node1)).Equal(4)
		g.Assert(node1.ChildCount()).Equal(2)

		g.Assert(node4.Parent).Equal(node2)
		g.Assert(node1.GrandParent() == nil).IsTrue()
		g.Assert(node4.GrandParent()).Equal(node1)
		g.Assert(node4.GrandParent()).Equal(node1)
		g.Assert(node4.ChildCount()).Equal(0)

		g.Assert(node1.GetNode(right)).Eql(node2)
		g.Assert(node1.GetNode(neutral) == nil).IsTrue()

		g.Assert(node2.Branch()).Eql(right)
		g.Assert(node2.IsBranchable(right)).IsTrue()
		g.Assert(node2.IsBranchable(left)).IsFalse()
		g.Assert(node2.ConjBranch().B).Equal(left.B) //use node branch
		g.Assert(left.ConjBranch()).Eql(right)       //use given branch
		g.Assert(node0.Branch()).Eql(left)
		g.Assert(node0.ConjBranch()).Eql(right) //use node branch
		//detach from parent
		node2.DetachFromParent()
		g.Assert(node2.Parent == nil).IsTrue()
		g.Assert(node2.Key).Eql(2)
		g.Assert(node1.Right == nil).IsTrue()
		//disconnect node
		node1.Disconnect()

		g.Assert(node1.Left == nil).IsTrue()
		g.Assert(node1.Right == nil).IsTrue()
		g.Assert(node1.Key == nil).IsTrue()
		g.Assert(node1.Parent == nil).IsTrue()

		g.Assert(null.Left == nil).IsTrue()
		g.Assert(null.Right == nil).IsTrue()
		g.Assert(null.Key == nil).IsTrue()
		g.Assert(null.Parent == nil).IsTrue()
	})

}
