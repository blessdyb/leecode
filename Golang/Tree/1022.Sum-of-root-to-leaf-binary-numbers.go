package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	type TreeNodePath struct {
		Node *TreeNode
		Path string
	}
	paths := []string{}
	stack := []*TreeNodePath{&TreeNodePath{
		Node: root,
		Path: "",
	}}
	for len(stack) > 0 {
		child := []*TreeNodePath{}
		for _, s := range stack {
			newPath := fmt.Sprintf("%s%d", s.Path, s.Node.Val)
			if s.Node.Left == nil && s.Node.Right == nil {
				paths = append(paths, newPath)
			}
			if s.Node.Left != nil {
				child = append(child, &TreeNodePath{
					Node: s.Node.Left,
					Path: newPath,
				})
			}
			if s.Node.Right != nil {
				child = append(child, &TreeNodePath{
					Node: s.Node.Right,
					Path: newPath,
				})
			}
		}
		stack = child
	}
	var ret int64 = 0
	for _, path := range paths {
		val, _ := strconv.ParseInt(path, 2, 32)
		ret += val
	}
	return int(ret)
}
