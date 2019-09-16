package trees

// BTree contains a value, and two nodes (left and right)
type BTree struct {
	Left  *BTree
	Val   int
	Right *BTree
}

// New returns a newly created BTree
func New(v int) *BTree {
	return &BTree{
		Left:  nil,
		Val:   v,
		Right: nil,
	}
}

// Insert adds a new node to the tree, or creates a tree if one doesn't already exist
func (t *BTree) Insert(v int) {
	if t == nil {
		t = New(v)
	}

	if v < t.Val {
		t.Left = New(v)
	} else {
		t.Right = New(v)
	}
}
