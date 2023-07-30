package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	root *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{
		root: root,
	}
}

func (this *CBTInserter) Insert(val int) int {
	stack := []*TreeNode{this.root}
	ret := -1
	for len(stack) > 0 {
		child := []*TreeNode{}
		for _, s := range stack {
			if s.Left != nil {
				child = append(child, s.Left)
			}
			if s.Right != nil {
				child = append(child, s.Right)
			}
		}
		// next layer is not full
		if len(stack)*2 > len(child) {
			if len(child)%2 == 0 {
				parentNode := stack[(len(child)+1)/2]
				ret = parentNode.Val
				parentNode.Left = &TreeNode{
					Val: val,
				}
			} else {
				parentNode := stack[len(child)/2]
				ret = parentNode.Val
				parentNode.Right = &TreeNode{
					Val: val,
				}
			}
			break
			// next layer is a full layer
		} else if len(child) == 0 {
			parentNode := stack[0]
			ret = parentNode.Val
			parentNode.Left = &TreeNode{
				Val: val,
			}
			break
		}
		stack = child
	}
	return ret
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}
