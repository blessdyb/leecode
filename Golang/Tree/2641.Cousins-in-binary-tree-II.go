package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	hashmap := map[*TreeNode]int{}
	for len(stack) > 0 {
		child := []*TreeNode{}
		childSum := 0
		for _, s := range stack {
			if s.Left != nil && s.Right != nil {
				hashmap[s.Left] = s.Right.Val
				hashmap[s.Right] = s.Left.Val
			}
			if s.Left != nil {
				child = append(child, s.Left)
				childSum += s.Left.Val
			}
			if s.Right != nil {
				child = append(child, s.Right)
				childSum += s.Right.Val
			}
		}
		for _, c := range child {
			c.Val = childSum - c.Val - hashmap[c]
		}
		stack = child
	}
	root.Val = 0
	return root
}
