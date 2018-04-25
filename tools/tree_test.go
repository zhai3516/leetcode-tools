package tools

import "testing"

func TestIntSlice2TreeNode(t *testing.T) {
	if IntSlice2TreeNodes(nil) != nil {
		t.Error("Error with: nil")
	}

	var empty []int
	if IntSlice2TreeNodes(a) != nil {
		t.Error("Error with: []")
	}

	var tree1 = [3]int{1, 2, 3}
	root := IntSlice2TreeNodes(a)
	if root.Val != 1 || root.Left.Val != 2 || root.Right.Val != 3 {
		t.Error("Error with: [1, 2, 3]")
	}

	var tree2 = [4]int{1, nil, 2, 3}
	root := IntSlice2TreeNodes(a)
	if root.Val != 1 || root.Left != nil || root.Right.Val != 2 || root.Right.Left.Val != 3 {
		t.Error("Error with: [1, nil, 2, 3]")
	}

}
