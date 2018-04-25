package tools

import (
	"github.com/zhai3516/leetcode-tools/model"
)

func IntSlice2TreeNodes(s []int) *model.TreeNode {
	if s == nil || len(s) == 0 {
		return nil
	}

	var treeLists [len(s)]*model.TreeNode
	for i, v := range s {
		if v != nil {
			treeLists[i] = &TreeNode{
				Val: v,
			}
		} else {
			treeLists[i] = nil
		}
	}

	left, cur := true, 0
	for i, v := range treeLists {
		if i == 0 {
			continue
		}

		if left {
			treeLists[cur].Left = v
			left = false
		} else {
			treeLists[cur].Right = v
			for cur < len(treeLists) {
				cur += 1
				if treeLists[cur] != nil {
					left = false
					break
				}
			}
		}
	}

	return treeLists[0]
}
