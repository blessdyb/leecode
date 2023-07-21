package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// For BST, in-order tranversal can give us a non-decreasing list
// Two cases, a) switch neighbor nodes, b) switch non-neighbour nodes
func recoverTreeBST(root *TreeNode) {
	var prev, first, second *TreeNode
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node != nil {
			helper(node.Left)
			if prev != nil && prev.Val > node.Val {
				if first == nil {
					// Case a), switch prev and node
					first = prev
					second = node
				} else {
					// case b), switch first and second
					second = node
				}
			}
			prev = node
			helper(node.Right)
		}
	}
	helper(root)
	first.Val, second.Val = second.Val, first.Val
}

func recoverTreeMorris(root *TreeNode) {
	var prevNode, first, second *TreeNode
	current := root
	for current != nil {
		if current.Left == nil {
			if prevNode != nil && current.Val < prevNode.Val {
				if first == nil {
					first = prevNode
					second = current
				} else {
					second = current
				}
			}
			prevNode = current
			current = current.Right
		} else {
			prev := current.Left
			for prev.Right != nil && prev.Right != current {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = current
				current = current.Left
			}
			if prev.Right == current {
				prev.Right = nil
				if prevNode != nil && current.Val < prevNode.Val {
					if first == nil {
						first = prevNode
						second = current
					} else {
						second = current
					}
				}
				prevNode = current
				current = current.Right
			}
		}
	}
	first.Val, second.Val = second.Val, first.Val
}
