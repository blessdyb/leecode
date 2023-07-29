package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	type NodeWithParent struct {
		*TreeNode
		Parent int
	}
	stack := []*NodeWithParent{&NodeWithParent{
		TreeNode: root,
		Parent:   0,
	}}
	for len(stack) > 0 {
		child := []*NodeWithParent{}
		var xNode, yNode *NodeWithParent
		for _, s := range stack {
			if s.Left != nil {
				temp := &NodeWithParent{
					TreeNode: s.Left,
					Parent:   s.Val,
				}
				child = append(child, temp)
				if temp.Val == x {
					xNode = temp
				} else if temp.Val == y {
					yNode = temp
				}
			}
			if s.Right != nil {
				temp := &NodeWithParent{
					TreeNode: s.Right,
					Parent:   s.Val,
				}
				child = append(child, temp)
				if temp.Val == x {
					xNode = temp
				} else if temp.Val == y {
					yNode = temp
				}
			}
		}
		if xNode != nil && yNode != nil && xNode.Parent != yNode.Parent {
			return true
		}
		stack = child
	}
	return false
}
