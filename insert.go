package bst

//InsertItem inserts value in to bst.
func InsertItem(tree *BST, val interface{}) *Node {
	var newnode = NewNode(val, tree.Cmp)
	//early exit
	if IsNil(tree.Root) {
		tree.Root = newnode
		tree.size += 1 //size
		return tree.Root
	}

	var node = tree.Root
	var found = false
	var left  = NewBranch().AsLeft()
	var right = NewBranch().AsRight()

	for !found {
		if node.keyCmp(val, node.Key) == left.B {
			// Go left
			if IsNil(node.Left) {
				Ptr(node, newnode, left)
				tree.size += 1 //size
				found = true
			}
			node = node.Left
		} else if node.keyCmp(val, node.Key) == right.B {
			// Go right
			if IsNil(node.Right) {
				Ptr(node, newnode, right)
				tree.size += 1 //size
				found = true
			}
			node = node.Right
		} else {
			//value already exist
			newnode = node
			found = true
		}
	}

	return newnode
}
