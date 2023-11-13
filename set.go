package bst

import (
	"github.com/intdxdt/stack"
)

const (
	setUnion = iota
	setIntersect
	setDiff
	setSymDiff
)

// Performs set operation given two trees.
func set_operation(n1, n2 *Node, mode int) []interface{} {
	vals := make([]interface{}, 0)
	// Create two stacks for two inorder traversals
	var s1, s2, empty = stack.NewStack[*Node](), stack.NewStack[*Node](), stack.NewStack[*Node]()

	var iter = func() bool {
		return (NotNil(n1) || Not(s1.IsEmpty()) ||
			NotNil(n2) || Not(s2.IsEmpty()))
	}

	for iter() {
		if NotNil(n1) {
			s1.Push(n1)
			n1 = n1.Left
		} else if NotNil(n2) {
			s2.Push(n2)
			n2 = n2.Left
		} else if Not(s1.IsEmpty()) || Not(s2.IsEmpty()) {
			// If we reach a nil node and either of the stacks is empty,
			// then one tree is exhausted, ptint the other tree
			if s1.IsEmpty() || s2.IsEmpty() {
				//if not diff : symdiff, union, intersect,
				if mode != setDiff {
					empty = s2
				}

				//if not intersect:  diff, symdiff, union
				if mode != setIntersect {
					vals = empty_other_stack(s1, empty, vals)
				}
				break
			}

			// Both root1 and root2 are NULL here
			n1 = s1.Top()
			n2 = s2.Top()

			// If current keys in two trees are same
			if n1.Compare(n2) == 0 {

				if mode == setUnion || mode == setIntersect {
					vals = append(vals, n1.Key)
				}

				s1.Pop()
				s2.Pop()
				// move to the inorder successor
				n1 = n1.Right
				n2 = n2.Right

			} else if n1.Compare(n2) < 0 {

				if mode != setIntersect {
					vals = append(vals, n1.Key)
				}

				s1.Pop()
				n1 = n1.Right
				n2 = nil //make n2 nil , advance n1

			} else if n1.Compare(n2) > 0 {

				//if not diff or intersect : symdiff, union
				if mode != setDiff && mode != setIntersect {
					vals = append(vals, n2.Key)
				}

				s2.Pop()
				n2 = n2.Right
				n1 = nil //make n1 nil , advance n2
			}
		}
	}
	return vals
}

// Empties all stack values.
func empty_other_stack(s1, s2 *stack.Stack[*Node], vals []interface{}) []interface{} {
	for s1.IsEmpty() && Not(s2.IsEmpty()) {
		all_node_right_vals(s2, &vals)
	}
	//
	for s2.IsEmpty() && Not(s1.IsEmpty()) {
		all_node_right_vals(s1, &vals)
	}
	return vals
}

// Gets all right branches
func all_node_right_vals(s *stack.Stack[*Node], vals *[]interface{}) {
	var n, nL *Node
	for Not(s.IsEmpty()) {
		n = s.Pop()
		//prevent going left
		nL, n.Left = n.Left, nil
		InOrder(n, func(n *Node) bool {
			*vals = append(*vals, n.Key)
			return true
		})
		n.Left = nL
	}
}
