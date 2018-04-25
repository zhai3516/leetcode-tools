package tools

import (
	"testing"

	"github.com/zhai3516/leetcode-tools/model"
)

func TestSlice2TreeNode(t *testing.T) {
	if Slice2TreeNodes(nil) != nil {
		t.Error("Slice2TreeNode error with: nil")
	}

	var empty []interface{}
	if Slice2TreeNodes(empty) != nil {
		t.Error("Slice2TreeNode error with: []")
	}

	var tree1 = []interface{}{1, 2, 3}
	root := Slice2TreeNodes(tree1)
	if root.Val != 1 || root.Left.Val != 2 || root.Right.Val != 3 {
		t.Error("Slice2TreeNode error with: [1, 2, 3]")
	}

	var tree2 = []interface{}{1, nil, 2, 3}
	root = Slice2TreeNodes(tree2)
	if root.Val != 1 || root.Left != nil || root.Right.Val != 2 || root.Right.Left.Val != 3 {
		t.Error("Slice2TreeNode error with: [1, nil, 2, 3]")
	}

}

func TestTreeNode2Slice(t *testing.T) {
	var tree = &model.TreeNode{Val: 0}
	var treeList, results []interface{}

	if len(TreeNode2Slice(nil)) != 0 {
		t.Error("TreeNode2Slice error with: nil")
	}

	tree.Val = 1
	if len(TreeNode2Slice(tree)) != 1 {
		t.Error("TreeNode2Slice error with: [1]")
	}

	treeList = []interface{}{1, nil, 2, 3}
	tree = Slice2TreeNodes(treeList)
	results = TreeNode2Slice(tree)
	if len(results) != len(treeList) {
		t.Error("TreeNode2Slice error with: [1, nil, 2, 3]")
	}
	for i, v := range treeList {
		if v != results[i] {
			t.Error("TreeNode2Slice error with: [1, nil, 2, 3]")
			break
		}
	}
}
