package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	hashmap := map[int]*TreeNode{}
	rootmap := map[int]bool{}
	for _, desc := range descriptions {
		var parent, child *TreeNode
		if hashmap[desc[0]] != nil {
			parent = hashmap[desc[0]]
		} else {
			parent = &TreeNode{
				Val: desc[0],
			}
			rootmap[desc[0]] = true
		}
		if hashmap[desc[1]] != nil {
			child = hashmap[desc[1]]
		} else {
			child = &TreeNode{
				Val: desc[1],
			}
		}
		rootmap[desc[1]] = false
		if desc[2] == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}
	}
	rootIndex := -1
	for key, val := range rootmap {
		if val {
			rootIndex = key
			break
		}
	}
	return hashmap[rootIndex]
}
