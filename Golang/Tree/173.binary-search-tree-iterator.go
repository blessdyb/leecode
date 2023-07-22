package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	data    []*TreeNode
	current *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		data:    []*TreeNode{},
		current: root,
	}
}

func (this *BSTIterator) Next() int {
	for this.current != nil {
		this.data = append(this.data, this.current)
		this.current = this.current.Left
	}
	node := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	this.current = node.Right
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.current != nil || len(this.data) > 0
}
