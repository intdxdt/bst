package bst

import (
	"fmt"
	"sort"
	"testing"
	"github.com/intdxdt/cmp"
	"github.com/franela/goblin"
)

var valPrinter = func(n interface{}) string {
	return fmt.Sprintf("%v", n)
}

func TestBST(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("BST - Binary Search Tree", func() {
		var t0 *BST
		var t1 = t0.Clone()
		g.Assert(t0 == nil).IsTrue()
		g.Assert(t1 == nil).IsTrue()

		var left = NewBranch().AsLeft()
		var right = NewBranch().AsRight()

		var array = []float64{
			1, 0, 2, 3, 4, 3, 3.3, 9, 29, 3.1, 0.1, 1.1,
			1.81, 0.91, 0.81, 0.71, 0.88, 0.82, 0.81,
		}

		var uniqMap = make(map[float64]bool, 0)
		var uniqArray = make([]float64, 0, 0)

		var tree = NewBST(cmp.F64)
		g.Assert(tree.Root.NodeHeight()).Equal(0)
		g.Assert(tree.size).Equal(0)

		tree.Insert(0.0)
		g.Assert(tree.Height()).Equal(0)
		g.Assert(tree.size).Equal(1)

		for i := range array {
			var a = array[i]
			tree.Insert(a)
			if ok, _ := uniqMap[a]; !ok {
				uniqMap[a] = true
			}
		}

		//uniq array
		for a := range uniqMap {
			uniqArray = append(uniqArray, a)
		}
		sort.Float64s(uniqArray)

		g.Assert(len(uniqArray)).Equal(tree.size)

		var eachItem = make([]interface{}, 0)
		var eachUpto1 = make([]interface{}, 0)

		tree.EachItem(func(i interface{}) bool {
			eachItem = append(eachItem, i)
			return true
		})
		tree.EachItem(func(i interface{}) bool {
			if cmp.F64(i, 1.0) <= 0 {
				eachUpto1 = append(eachUpto1, i)
				return true
			}
			return false
		})
		treeClone := tree.Clone()

		treeArray := tree.ToArray()
		treeCloneArray := treeClone.ToArray()

		g.Assert(len(eachUpto1)).Equal(8)
		g.Assert(eachUpto1[0]).Equal(0.)
		g.Assert(eachUpto1[len(eachUpto1)-1]).Equal(1.)
		g.Assert(len(treeArray)).Equal(len(uniqArray))
		g.Assert(len(treeCloneArray)).Equal(len(uniqArray))

		g.Assert(len(eachItem)).Eql(len(uniqArray))
		for i := range treeArray {
			g.Assert(treeArray[i]).Equal(uniqArray[i])
			g.Assert(eachItem[i]).Equal(uniqArray[i])
		}

		var rm, parent = tree.Remove(10000.)
		g.Assert(rm == nil).IsTrue()
		g.Assert(len(uniqArray)).Equal(tree.size)

		g.Assert(BranchMost(rm, left) == nil).IsTrue()
		g.Assert(BranchMost(rm, right) == nil).IsTrue()
		g.Assert(BranchMost(tree.Root, left).Key).Equal(0.0)
		fmt.Println("branchmost : -->>", BranchMost(tree.Root, left).Key)

		g.Assert(parent == nil).IsTrue()
		g.Assert(tree.Height()).Equal(6)
		var n = tree.Search(1.81).Key
		g.Assert(n).Equal(1.81)
		g.Assert(tree.Search(nil) == nil).IsTrue()
		g.Assert(tree.Search(1.777) == nil).IsTrue()

		fmt.Println(tree.String(), "\n\n")

		rm, parent = tree.Remove(0.0)
		g.Assert(len(uniqArray) - 1).Equal(tree.size)
		g.Assert(SizeOfNode(tree.Root)).Equal(tree.size)

		g.Assert(rm.Key).Equal(0.0)
		g.Assert(parent == nil).IsTrue()
		g.Assert(tree.Root.Key).Equal(1.0)

		rm, parent = tree.Remove(0.81)
		rm, parent = tree.Remove(3.0)
		rm, parent = tree.Remove(3.3)
		rm, parent = tree.Remove(9.0)

		g.Assert(len(uniqArray) - 5).Equal(tree.size)
		g.Assert(SizeOfNode(tree.Root)).Equal(tree.size)

		g.Assert(BranchMost(tree.Root, left).Key == 0.1).IsTrue()
		g.Assert(BranchMost(tree.Root, right).Key == 29.).IsTrue()
		g.Assert(BranchMost(tree.Root.Right, left).Key == 1.1).IsTrue()

		g.Assert(tree.Height()).Equal(5)

		g.Assert(tree.size).Equal(12)
		g.Assert(SizeOfNode(tree.Root)).Equal(tree.size)

		rm, parent = tree.Remove(29.)
		g.Assert(rm.Key).Equal(29.)
		g.Assert(parent.Key).Eql(4.)
		rm, parent = tree.Remove(29.)
		rm, parent = tree.Remove(0.88)
		rm, parent = tree.Remove(1.11)

		g.Assert(SizeOfNode(tree.Root)).Equal(tree.size)

		var bst = NewBST(cmp.F64)
		bst.Insert(1.5)
		bst.Insert(2.5)

		g.Assert(bst.Root != nil).IsTrue()
		g.Assert(bst.Root.Key).Equal(1.5)

		rm, parent = bst.Remove(2.5)

		g.Assert(bst.Root != nil).IsTrue()
		g.Assert(bst.Root.ChildCount() == 0).IsTrue()
		rm, parent = bst.Remove(1.5)
		g.Assert(bst.Root == nil).IsTrue()
		g.Assert(bst.size).Equal(0)

		bst.Insert(1.5)
		bst.Insert(2.5)
		bst.Insert(0.5)
		rm, parent = bst.Remove(1.5)
		bst.Insert(1.5)
		bst.Insert(3.5)
		bst.Insert(8.4)
		bst.Insert(3.2)
		bst.Insert(1.0)
		bst.Insert(1.7)
		bst.Insert(1.9)
		bst.Insert(1.98)
		bst.Insert(1.8)

		rm, parent = bst.Remove(2.5)
		g.Assert(bst.Root.Left == nil).IsTrue()
		g.Assert(bst.Root.Right.Key == (1.98)).IsTrue()

		rm, parent = bst.Remove(0.5)
		g.Assert(bst.Root.Key == (1.98)).IsTrue()

		rm, parent = bst.Remove(1.98)
		g.Assert(bst.Root.Key == (1.9)).IsTrue()

		bst = NewBST(cmp.F64)
		var printVal = bst.Print(valPrinter)

		g.Assert(printVal).Equal("")

		fmt.Println(tree.Print(valPrinter))

		fmt.Println("\nWorse Case -- BST as List\n")

		bst = NewBST(cmp.F64)
		bst.Insert(1.1)
		bst.Insert(1.3)
		bst.Insert(1.5)
		bst.Insert(1.9)
		bst.Insert(2.1)
		bst.Insert(2.5)
		bst.Insert(2.7)
		fmt.Println(bst.Print(valPrinter))

		g.Assert(bst.Root.Key).Equal(1.1)
		g.Assert(BranchMost(bst.Root, left).Key).Equal(1.1)
		g.Assert(BranchMost(bst.Root, right).Key).Equal(2.7)
		g.Assert(bst.Size()).Equal(7)

		fmt.Println("\nTest Empty\n")
		g.Assert(bst.Empty().Root == nil).IsTrue()
		g.Assert(bst.size == 0).IsTrue()

		g.Assert(bst.Size()).Equal(0)
		g.Assert(bst.First() == nil).IsTrue()
		g.Assert(bst.Last() == nil).IsTrue()

		bst.Insert(1.5)
		bst.Insert(1.1)
		bst.Insert(2.5)
		g.Assert(bst.Root.Key).Equal(1.5)
		g.Assert(bst.First()).Equal(1.1)
		g.Assert(bst.Last()).Equal(2.5)

		printVal = bst.Print(valPrinter)
		fmt.Println(printVal)

		//Full Next Prev Test
		tree = NewBST(cmp.Int)
		vals := []int{
			66, 62, 70, 60, 50, 61, 64, 63,
			65, 68, 67, 69, 89, 71, 72, 101,}
		for _, f := range vals {
			tree.Insert(f)
		}

		//NextPrev
		g.Assert(tree.PrevItem(10) == nil).IsTrue()
		g.Assert(tree.NextItem(73) == nil).IsTrue()

		g.Assert(tree.PrevItem(50) == nil).IsTrue()
		g.Assert(tree.NextItem(101) == nil).IsTrue()

		g.Assert(tree.NextItem(66)).Eql(67)
		g.Assert(tree.PrevItem(66)).Eql(65)

		g.Assert(tree.NextItem(65)).Eql(66)
		g.Assert(tree.PrevItem(65)).Eql(64)

		g.Assert(tree.NextItem(67)).Eql(68)
		g.Assert(tree.PrevItem(67)).Eql(66)

	})
}

func TestSetBST(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("BST - Binary Search Tree - Intersection", func() {
		var tree1 = NewBST(cmp.Int)
		var tree2 = NewBST(cmp.Int)
		var tree3 = NewBST(cmp.Int)
		var tree4 = NewBST(cmp.Int)
		var tree5 = NewBST(cmp.Int)

		tree1.Insert(5)
		tree1.Insert(1)
		tree1.Insert(10)
		tree1.Insert(0)
		tree1.Insert(4)
		tree1.Insert(7)
		tree1.Insert(9)

		//===
		tree2.Insert(10)
		tree2.Insert(7)
		tree2.Insert(20)
		tree2.Insert(4)
		tree2.Insert(91)

		//===
		tree3.Insert(11)
		tree3.Insert(12)
		tree3.Insert(21)
		tree3.Insert(41)
		tree3.Insert(92)

		var empty = make([]interface{}, 0)

		g.Assert(tree1.Intersection(tree2)).Eql(
			[]interface{}{4, 7, 10},
		)
		g.Assert(tree2.Intersection(tree1)).Eql(
			[]interface{}{4, 7, 10},
		)
		g.Assert(tree3.Intersection(tree1)).Eql(empty)
		g.Assert(tree2.Intersection(tree3)).Eql(empty)

		diff1 := []interface{}{0, 1, 5, 9}
		diff2 := []interface{}{20, 91}
		sdiff := []interface{}{0, 1, 5, 9, 20, 91}
		g.Assert(tree1.Difference(tree2)).Eql(diff1)
		g.Assert(tree2.Difference(tree1)).Eql(diff2)
		g.Assert(tree1.XOR(tree2)).Eql(sdiff)

		g.Assert(tree1.Difference(tree1)).Eql(empty)
		g.Assert(tree2.Difference(tree2)).Eql(empty)

		//merge two trees
		mged := []interface{}{
			0, 1, 4, 5, 7,
			9, 10, 20, 91}
		mgd := tree1.Union(tree2)
		fmt.Println("merged->", mgd)
		g.Assert(mgd).Eql(mged)

		mgd2 := tree1.UnionTree(tree2)
		g.Assert(mgd).Eql(mgd2.ToArray())

		//union with an empty tree
		mgd4 := tree4.UnionTree(tree1)

		//clone an empty tree
		mgd5 := tree5.Clone()
		tree1.EachItem(func(o interface{}) bool {
			mgd5.Insert(o)
			return true
		})
		g.Assert(mgd5.ToArray()).Eql(mgd4.ToArray())

		fmt.Println("\ntree-----1:\n", tree1)
		fmt.Println("\ntree-----2:\n", tree2)
		fmt.Println("\ntree union:\n", mgd2)
	})
}

func TestSetBST2(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("BST - Binary Search Tree 2 - difference ", func() {
		var tree1 = NewBST(cmp.Int)
		var tree2 = NewBST(cmp.Int)

		tree1.Insert(53)
		tree1.Insert(33)
		tree1.Insert(63)
		tree1.Insert(03)
		tree1.Insert(43)
		tree1.Insert(73)
		tree1.Insert(93)
		//===
		tree2.Insert(13)
		tree2.Insert(71)
		tree2.Insert(20)
		tree2.Insert(42)
		tree2.Insert(91)

		var empty = make([]interface{}, 0)

		g.Assert(tree1.Intersection(tree2)).Eql(empty)
		g.Assert(tree2.Intersection(tree1)).Eql(empty)

		g.Assert(tree1.Difference(tree2)).Eql(tree1.ToArray())
		g.Assert(tree2.Difference(tree1)).Eql(tree2.ToArray())

		xor12 := []interface{}{
			03, 13, 20, 33, 42, 43,
			53, 63, 71, 73, 91, 93}
		g.Assert(tree1.XOR(tree2)).Eql(xor12)
	})
}
