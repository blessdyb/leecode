package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	ret := []int{}
	current1, current2 := root1, root2
	stack1, stack2 := []*TreeNode{}, []*TreeNode{}
	for {
		if current1 != nil {
			stack1 = append(stack1, current1)
			current1 = current1.Left
		} else if current2 != nil {
			stack2 = append(stack2, current2)
			current2 = current2.Right
		} else if len(stack1) > 0 || len(stack2) > 0 {
			if len(stack2) == 0 {
				node := stack1[len(stack1)-1]
				stack1 = stack1[:len(stack1)-1]
				ret = append(ret, node.Val)
				current1 = node.Right
			} else if len(stack1) == 0 {
				node := stack2[len(stack2)-1]
				stack2 = stack2[:len(stack2)-1]
				ret = append(ret, node.Val)
				current2 = node.Right
			} else {
				node1 := stack1[len(stack1)-1]
				node2 := stack2[len(stack2)-1]
				if node1.Val > node2.Val {
					stack2 = stack2[:len(stack2)-1]
					ret = append(ret, node2.Val)
					current2 = node2.Right
				} else {
					stack1 = stack1[:len(stack1)-1]
					ret = append(ret, node1.Val)
					current1 = node1.Right
				}
			}
		} else {
			break
		}
	}
	return ret
}
