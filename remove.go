package bst

//Empties the bst tree
func Empty(tree *BST) *BST {
	tree.Root = nil
	tree.size = 0
	return tree
}

//RemoveItem - remove node with value at given key
func RemoveItem(tree *BST, val interface{}) (*Node, *Node) {

	var childcount int
	var successor *Node
	var child *Node
	var node *Node
	var left = NewBranch().AsLeft()
	var right = NewBranch().AsRight()

	if tree != nil {
		node = tree.Root
	}

	// if not found and node is valid
	node = SearchItem(node, val)

	if IsNil(node) {
		return nil, nil //early exit
	}

	var sucparent = node.Parent
	childcount = node.ChildCount()
	if childcount == 0 {
		if node == tree.Root {
			//is root
			tree.Root = nil
		} else {
			//is leaf
			node.DetachFromParent()
		}
	} else if childcount == 1 {
		//find 1 child
		if IsNil(node.Right) {
			child = node.Left
		} else {
			child = node.Right
		}

		//make node.parent the childs parent
		Ptr(node.Parent, child, node.Branch())

		if node == tree.Root {
			tree.Root = child //is root
		}
	} else {
		successor = node.Left
		successor = BranchMost(successor, right)
		sucparent = successor.Parent
		if successor.Parent == node {
			//node immediate left -> right child is the same as successor
			sucparent = successor
			//update right branch
			Ptr(successor, node.Right, right)
			//update parent
			Ptr(node.Parent, successor, node.Branch()) //update successor parent
			if node == tree.Root {
				//is root , update new root as successor
				tree.Root = successor
			}
		} else {
			//successor is a deeper left right child
			Ptr(successor.Parent, successor.Left, right)
			//assign children to the successor
			Ptr(successor, node.Right, right)
			Ptr(successor, node.Left, left)
			//update successor and new parent
			Ptr(node.Parent, successor, node.Branch())

			//is root , update new root as successor
			if node == tree.Root {
				tree.Root = successor
			}
		}
	}

	tree.size += -1 //size
	return node, sucparent
}
