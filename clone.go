package bst

import "github.com/intdxdt/stack"


//Clones BST.
func clone(tree *BST) *BST {
	if tree == nil {
		return tree
	}
	node := tree.Root
	other := NewBST(tree.Cmp)

	if node == nil {
		return other
	}

	other.Root = node.Clone(node.Parent)
	_node := other.Root

	var _stack = stack.NewStack()
	for !_stack.IsEmpty() || !IsNil(node) {
		if !IsNil(node) {
			//pack stack _node, node
			pack_2_stack(_stack, _node, node)
			//copy ptr to left branch
			if !IsNil(node.Left) {
				_node.Left = node.Left.Clone(_node)
			}
			node, _node = node.Left, _node.Left
		} else {
			//unpack stack  node, _node
			node, _node = unpack_2_stack(_stack)
			if !IsNil(node) {
				//process
				if !IsNil(node.Right) {
					_node.Right = node.Right.Clone(_node)
				}
				node, _node = node.Right, _node.Right
			}
		}
	}
	return other
}

//Packs two nodes onto stack.
func pack_2_stack(s *stack.Stack, _n, n *Node) {
	s.Push(_n)
	s.Push(n)
}

//Unpacks two nodes from stack.
func unpack_2_stack(s *stack.Stack) (*Node, *Node) {
	n := s.Pop().(*Node)
	_n := s.Pop().(*Node)
	return n, _n
}
