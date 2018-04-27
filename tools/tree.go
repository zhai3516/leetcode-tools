package tools

import (
	"github.com/Workiva/go-datastructures/queue"
	"github.com/zhai3516/leetcode-tools/model"
)

func Slice2TreeNodes(s []interface{}) *model.TreeNode {
	if s == nil || len(s) == 0 {
		return nil
	}

	var treeLists = make([]*model.TreeNode, len(s))
	for i, v := range s {
		if v != nil {
			treeLists[i] = &model.TreeNode{
				Val: v.(int),
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
			for cur < len(treeLists)-1 {
				cur += 1
				if treeLists[cur] != nil {
					left = true
					break
				}
			}
		}
	}

	return treeLists[0]
}

func TreeNode2Slice(root *model.TreeNode) []interface{} {
	var results []interface{} = make([]interface{}, 0)

	if root == nil {
		return results
	}

	var rootFlag = true
	var que = queue.New(0)
	que.Put(root)

	// do while is better
	for !que.Empty() {
		tempTree, _ := que.Get(1)
		if tempTree[0] != nil {
			results = append(results, tempTree[0].(*model.TreeNode).Val)

			if tempTree[0].(*model.TreeNode).Left != nil {
				que.Put(tempTree[0].(*model.TreeNode).Left)
			} else if rootFlag || que.Len() > 0 {
				// root should be trade differe
				que.Put(nil)
				rootFlag = false
			}

			if tempTree[0].(*model.TreeNode).Right != nil {
				que.Put(tempTree[0].(*model.TreeNode).Right)
			} else if que.Len() > 1 {
				que.Put(nil)
			}
		} else {
			results = append(results, nil)
		}
		rootFlag = false
	}

	if results[len(results)-1] == nil {
		results = append(results[:len(results)-1])
	}

	return results
}
