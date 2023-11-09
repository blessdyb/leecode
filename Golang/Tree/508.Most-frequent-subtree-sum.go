package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	hashmap := map[int]int{}
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		value := node.Val + helper(node.Left) + helper(node.Right)
		hashmap[value]++
		return value
	}
	helper(root)
	frequencyHashmap := map[int][]int{}
	max := 0
	for k, v := range hashmap {
		if _, ok := frequencyHashmap[v]; !ok {
			frequencyHashmap[v] = []int{}
		}
		if max < v {
			max = v
		}
		frequencyHashmap[v] = append(frequencyHashmap[v], k)
	}
	return frequencyHashmap[max]
}
