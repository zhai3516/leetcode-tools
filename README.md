# Tools for leetcode
This project provide some tools to test leetcode answers more easy.

## Tree: Slice <-> TreeNode
This tool help to transform slice to TreeNode and vice versa.
```
Slice2TreeNode: []interface{} -> *model.TreeNode
TreeNode2Slice: *model.TreeNode -> []interface{}
```

# lckit
lckit is a tool that help to create leetcode code template

## Install
`go get github.com/zhai3516/leetcode-tools`

`go install github.com/zhai3516/leetcode-tools/lckit`

## Usage
`lckit --kit=tmp --name="Flatten Binary Tree to Linked List" --type=tree`
