package bst

import "github.com/intdxdt/stack"

// InOrder - inorder traversal
func InOrder(node *Node, callback func(*Node) bool) {
	var stk = stack.NewStack[*Node]()
	for !stk.IsEmpty() || NotNil(node) {
		if NotNil(node) {
			stk.Push(node)
			node = node.Left
		} else {
			node = stk.Pop()
			if NotNil(node) {
				//process
				if !callback(node) {
					break
				}
				node = node.Right
			}
		}
	}
}
