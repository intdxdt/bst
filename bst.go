// Package bst is a binary search tree,
// items should implement the item interface.
package bst

import (
	"fmt"
	"github.com/intdxdt/cmp"
	"github.com/intdxdt/stack"
)

type BST struct {
	Root *Node
	size int
	Cmp  cmp.Compare
}

// NewBST - New Binary  Search Tree
func NewBST(comparator cmp.Compare) *BST {
	return &BST{Root: nil, size: 0, Cmp: comparator}
}

func (bst *BST) Clone() *BST {
	return clone(bst)
}

// Search item
func (bst *BST) Search(val interface{}) *Node {
	return SearchItem(bst.Root, val)
}

// Insert - inserts value by key
func (bst *BST) Insert(val interface{}) *Node {
	return InsertItem(bst, val)
}

// Remove node with value at given key
func (bst *BST) Remove(val interface{}) (*Node, *Node) {
	return RemoveItem(bst, val)
}

// Empty the BST tree
func (bst *BST) Empty() *BST {
	return Empty(bst)
}

// Traverse each node in the BST
func (bst *BST) Traverse(callback func(*Node) bool) {
	InOrder(bst.Root, callback)
}

// EachItem - Iterates over each item in the BST
func (bst *BST) EachItem(fn func(interface{}) bool) {
	InOrder(bst.Root, func(n *Node) bool { return fn(n.Key) })
}

// Size of bst node, note this is O(n)
func (bst *BST) Size() int {
	return bst.size
}

// Height - computes height of bst tree
func (bst *BST) Height() int {
	return bst.Root.NodeHeight()
}

// Union - computes union of binary search trees
func (bst *BST) Union(other *BST) []interface{} {
	return set_operation(bst.Root, other.Root, setUnion)
}

// UnionTree - union tree
func (bst *BST) UnionTree(other *BST) *BST {
	var tree = NewBST(bst.Cmp)
	var items = bst.Union(other)

	var start, mid, end int
	var rng = [2]int{0, len(items) - 1}
	var _stack = stack.NewStack[[2]int]()

	_stack.Push(rng)

	for !_stack.IsEmpty() {
		rng = _stack.Pop()
		start, end = rng[0], rng[1]
		if start > end {
			continue
		}

		mid = (start + end) / 2

		tree.Insert(items[mid]) //process

		_stack.Push([2]int{mid + 1, end})
		_stack.Push([2]int{start, mid - 1})
	}

	return tree
}

// Intersection - computes intersection of two binary search trees
func (bst *BST) Intersection(other *BST) []interface{} {
	return set_operation(bst.Root, other.Root, setIntersect)
}

// Difference - computes the difference between bst and other binary search trees
func (bst *BST) Difference(other *BST) []interface{} {
	return set_operation(bst.Root, other.Root, setDiff)
}

// XOR  - computes the difference between bst and other binary search trees
func (bst *BST) XOR(other *BST) []interface{} {
	return set_operation(bst.Root, other.Root, setSymDiff)
}

// ToArray- tree as array
func (bst *BST) ToArray() []interface{} {
	var result = make([]interface{}, 0)
	bst.Traverse(func(n *Node) bool {
		result = append(result, n.Key)
		return true
	})

	return result
}

// First item in the Tree
func (bst *BST) First() interface{} {
	var node = BranchMost(bst.Root, NewBranch().AsLeft())
	if node == nil {
		return nil
	}
	return node.Key
}

// Last item in the Tree
func (bst *BST) Last() interface{} {
	var node = BranchMost(bst.Root, NewBranch().AsRight())
	if node == nil {
		return nil
	}
	return node.Key
}

// NextItem - Next item to a given item
func (bst *BST) NextItem(val interface{}) interface{} {
	var node = Next(bst, val)
	if node == nil {
		return nil
	}
	return node.Key
}

// PrevItem - Prev item to a given item
func (bst *BST) PrevItem(val interface{}) interface{} {
	var node = Prev(bst, val)
	if node == nil {
		return nil
	}
	return node.Key
}

func (bst *BST) String() string {
	return bst.Print(bst.KeyPrinter)
}

// Print - print tree structure as string
func (bst *BST) Print(strfn func(interface{}) string) string {
	return PrintBST(bst.Root, strfn)
}

// KeyPrinter - converts key item to string
func (bst *BST) KeyPrinter(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
