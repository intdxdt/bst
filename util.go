package bst

// SearchItem - find node with given value
func SearchItem(node *Node, val interface{}) *Node {
	var found = false
	if val == nil {
		return nil
	}

	for !found && node != nil {
		var v = node.keyCmp(val, node.Key)
		if v < 0 {
			node = node.Left
		} else if v > 0 {
			node = node.Right
		} else {
			found = true
		}
	}
	if found {
		return node
	}
	return nil
}

// SizeOfNode - computes Size of Node
func SizeOfNode(node *Node) int {
	var n = 0
	InOrder(node, func(o *Node) bool {
		n += 1
		return true
	})
	return n
}

// Ptr - update parent and child bi-directional pointers
func Ptr(parent, child *Node, branch *Branch) {
	if NotNil(parent) && branch != nil {
		parent.SetNode(child, branch)
	}

	if NotNil(child) {
		child.Parent = parent
	}
}

// Not boolean
func Not(b bool) bool {
	return !b
}

// IsNil - checks if node is nil
func IsNil(n *Node) bool {
	return n == nil
}

// NotNil - checks if not is not nil
func NotNil(n *Node) bool {
	return n != nil
}

// ShiftNode from the beginning of list
func ShiftNode(a []*Node) (*Node, []*Node) {
	if len(a) == 0 {
		return nil, a
	}
	x := a[0]
	a[0] = nil
	return x, a[1:]
}

// PopNode pops a node from a slice of nodes.
func PopNode(a []*Node) (*Node, []*Node) {
	var v *Node
	var n int
	if len(a) == 0 {
		return nil, a
	}
	n = len(a) - 1
	v, a[n] = a[n], nil
	return v, a[:n]
}

// Next - find the next value
func Next(tree *BST, val interface{}) *Node {
	var n, p *Node

	if tree != nil {
		n = tree.Root
	}

	// if not found and node is valid
	n = SearchItem(n, val)

	if IsNil(n) {
		return nil //early exit
	}
	//-> right , min value on left branch
	if NotNil(n.Right) {
		n = BranchMost(n.Right, NewBranch().AsLeft())
		return n
	}
	//the above algorithm
	p = n.Parent
	for NotNil(p) && n == p.Right {
		n, p = p, p.Parent
	}
	return p
}

// Prev - find previous value
func Prev(tree *BST, val interface{}) *Node {
	var n, p *Node

	if tree != nil {
		n = tree.Root
	}
	// if not found and node is valid
	n = SearchItem(n, val)

	if IsNil(n) {
		return nil //early exit
	}
	//-> left , max value on right branch
	if NotNil(n.Left) {
		n = BranchMost(n.Left, NewBranch().AsRight())
		return n
	}
	//the above algorithm
	p = n.Parent
	for p != nil && n == p.Left {
		n, p = p, p.Parent
	}
	return p
}
