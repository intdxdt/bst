package bst

import (
	"github.com/intdxdt/cmp"
)

//Binary Tree Node
type Node struct {
	Key    interface{}
	Left   *Node
	Right  *Node
	Parent *Node
	Height int
	keyCmp cmp.Compare
}

//NewNode - Create  New Binary Search Tree Node
func NewNode(val interface{}, comparator cmp.Compare) *Node {
	return &Node{
		Key:    val,
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Height: 0,
		keyCmp: comparator,
	}
}

//Compare - item interface
func (node *Node) Compare(other *Node) int {
	if other == nil {
		return 1
	}
	return node.keyCmp(node.Key, other.Key)
}

//Clone node
func (node *Node) Clone(parent ...*Node) *Node {
	var p *Node
	if len(parent) > 0 {
		p = parent[0]
	}
	return &Node{
		Key:    node.Key,
		Left:   nil,
		Right:  nil,
		Parent: p,
		Height: node.Height,
	}
}

//NodeHeight - computes height of node
func (node *Node) NodeHeight() int {
	return height(node)
}

//Disconnect  node
func (node *Node) Disconnect() *Node {
	node.Key = nil
	node.Left = nil
	node.Right = nil
	node.Parent = nil
	return node
}

//SetNode at branch , does nothing if branch is neutral
func (node *Node) SetNode(n *Node, br *Branch) {
	if br.IsLeft() {
		node.Left = n
	} else if br.IsRight() {
		node.Right = n
	}
}

//GetNode on branch
func (node *Node) GetNode(br *Branch) *Node {
	var n *Node
	if br.IsLeft() {
		n = node.Left
	} else if br.IsRight() {
		n = node.Right
	}
	return n
}

//IsBranchable - is node branchable to left or right, is always false on neutral
func (node *Node) IsBranchable(br *Branch) bool {
	var bln bool
	if br.IsLeft() {
		bln = (node.Left != nil )
	} else if br.IsRight() {
		bln = (node.Right != nil )
	}
	return bln
}

//DetachFromParent - detaches node from parent
func (node *Node) DetachFromParent() *Branch {
	var br = NewBranch()
	if node.Parent != nil {
		br = node.Branch()
	}

	//update forward point from parent -> child
	if !br.IsNeutral() {
		node.Parent.SetNode(nil, br)
	}
	// update backward pointer parent <- child
	node.Parent = nil
	return br
}

//Branch - get branch of node with respect to parent
func (node *Node) Branch() *Branch {
	var br = NewBranch()
	if node.Parent != nil {
		if node == node.Parent.Left {
			br.B = Left
		} else if node == node.Parent.Right {
			br.B = Right
		}
	}
	return br
}

//ConjBranch finds conjugate branch of node
func (node *Node) ConjBranch() *Branch {
	return node.Branch().ConjBranch()
}

//GrandParent get node grandparent
func (node *Node) GrandParent() *Node {
	if node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

//ChildCount - the number child count
func (node *Node) ChildCount() int {
	var count int = 0
	if node.Left != nil {
		count += 1
	}
	if node.Right != nil {
		count += 1
	}
	return count
}
