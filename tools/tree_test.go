package tools

import "testing"

func TestIntSlice2TreeNode(t *testing.T) {
	if IntSlice2TreeNodes(nil) != nil {
		t.Error("Error with: nil")
	}

	var empty []interface{}
	if IntSlice2TreeNodes(empty) != nil {
		t.Error("Error with: []")
	}

	var tree1 = []interface{}{1, 2, 3}
	root := IntSlice2TreeNodes(tree1)
	if root.Val != 1 || root.Left.Val != 2 || root.Right.Val != 3 {
		t.Error("Error with: [1, 2, 3]")
	}

	var tree2 = []interface{}{1, nil, 2, 3}
	root = IntSlice2TreeNodes(tree2)
	if root.Val != 1 || root.Left != nil || root.Right.Val != 2 || root.Right.Left.Val != 3 {
		t.Error("Error with: [1, nil, 2, 3]")
	}

}
