package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestorRecursive(root.Left, p, q)
	right := lowestCommonAncestorRecursive(root.Right, p, q)
	if left != nil && right != nil {
		// it means for given root, we can find p and q in it's left/right sub trees
		return root
	} else if left == nil {
		return right
	}
	return left
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	type ComplextNode struct {
		Node      *TreeNode
		Ancestors []*TreeNode
	}
	var getAncestors func(root *TreeNode, node *TreeNode) []*TreeNode
	getAncestors = func(root *TreeNode, node *TreeNode) []*TreeNode {
		var ret []*TreeNode
		stack := []ComplextNode{ComplextNode{
			Node:      root,
			Ancestors: []*TreeNode{},
		}}
		for len(stack) > 0 {
			children := []ComplextNode{}
			for i := 0; i < len(stack); i++ {
				temp := stack[i]
				if temp.Node == node {
					ret = append(temp.Ancestors, temp.Node)
					break
				}
				if temp.Node.Left != nil {
					children = append(children, ComplextNode{
						Node:      temp.Node.Left,
						Ancestors: append(temp.Ancestors, temp.Node),
					})
				}
				if temp.Node.Right != nil {
					children = append(children, ComplextNode{
						Node:      temp.Node.Right,
						Ancestors: append(temp.Ancestors, temp.Node),
					})
				}
			}
			stack = children
		}
		return ret
	}
	ancestorsP := getAncestors(root, p)
	ancestorsQ := getAncestors(root, q)
	var ret *TreeNode
	i := 0
	for i < len(ancestorsP) && i < len(ancestorsQ) {
		if ancestorsP[i] == ancestorsQ[i] {
			ret = ancestorsP[i]
		} else {
			break
		}
		i++
	}
	return ret
}
