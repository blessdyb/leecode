package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	type TreeNodeWithParent struct {
		Val    int
		Left   *TreeNodeWithParent
		Right  *TreeNodeWithParent
		Parent *TreeNodeWithParent
	}
	ret := []int{}
	visited := map[*TreeNodeWithParent]bool{}
	var targetWithParent *TreeNodeWithParent
	var transformTree func(node *TreeNode, parent *TreeNodeWithParent) *TreeNodeWithParent
	var dfs func(node *TreeNodeWithParent, distance int)
	transformTree = func(node *TreeNode, parent *TreeNodeWithParent) *TreeNodeWithParent {
		if node != nil {
			temp := &TreeNodeWithParent{
				Val:    node.Val,
				Parent: parent,
			}
			temp.Left = transformTree(node.Left, temp)
			temp.Right = transformTree(node.Right, temp)
			if node == target {
				targetWithParent = temp
			}
			return temp
		}
		return nil
	}
	dfs = func(node *TreeNodeWithParent, distance int) {
		if node != nil && !visited[node] {
			visited[node] = true
			if distance == 0 {
				ret = append(ret, node.Val)
			} else {
				dfs(node.Parent, distance-1)
				dfs(node.Left, distance-1)
				dfs(node.Right, distance-1)
			}
		}
	}
	transformTree(root, nil)
	dfs(targetWithParent, k)
	return ret
}
