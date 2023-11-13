package bst

import "github.com/intdxdt/list"

// Computes height of binary tree node.
func height(node *Node) int {
	if node == nil {
		return 0 //early exit
	}
	var queue = list.NewList[*Node]()

	queue.Append(node)

	var curLevel = 1
	var depth = -1
	var nextLevel = 0
	for queue.Len() > 0 {
		node = queue.PopLeft()
		curLevel += -1
		if node.Left != nil {
			queue.Append(node.Left)
			nextLevel += 1
		}
		if node.Right != nil {
			queue.Append(node.Right)
			nextLevel += 1
		}
		if curLevel == 0 {
			depth += 1
			curLevel = nextLevel
			nextLevel = 0
		}
	}
	return depth
}
