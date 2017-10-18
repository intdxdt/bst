package bst

//Branch Type
type Branch struct {
	B int
}

const (
	Left    = -1 + iota
	Neutral
	Right
)

//NewBranch - creates a new brach.
//Branch path defaults to neutral
func NewBranch() *Branch {
	return &Branch{Neutral}
}

//Clone clones a bst branch.
func (branch *Branch) Clone() *Branch {
	var clone = NewBranch()
	clone.B = branch.B
	return clone
}

//AsLeft sets branch as left.
func (branch *Branch) AsLeft() *Branch {
	branch.B = Left
	return branch
}

//AsNeutral sets branch as neutral.
func (branch *Branch) AsNeutral() *Branch {
	branch.B = Neutral
	return branch
}

//AsRight sets branch as right.
func (branch *Branch) AsRight() *Branch {
	branch.B = Right
	return branch
}

//IsNeutral check if branch is  neutral.
func (branch *Branch) IsNeutral() bool {
	return branch.B == Neutral
}

//IsLeft checks if branch is left.
func (branch *Branch) IsLeft() bool {
	return branch.B == Left
}

//IsRight checks if branch is right.
func (branch *Branch) IsRight() bool {
	return branch.B == Right
}

//ConjBranch computes the conjugate of a branch
func (branch *Branch) ConjBranch() *Branch {
	var conj = branch.Clone().AsNeutral()
	if branch.IsRight() {
		conj.AsLeft()
	} else if branch.IsLeft() {
		conj.AsRight()
	}
	return conj
}

//BranchMost gets Node at leftmost or rightmost brach/sub-branch of tree.
func BranchMost(node *Node, branch *Branch) *Node {
	if IsNil(node) || branch == nil {
		return nil
	}
	for node.IsBranchable(branch) {
		node = node.GetNode(branch)
	}
	return node
}
