package main

type Node struct {
	Val [2]int,
	Left, Right *Node,
}

type MyCalendar struct {
	data [][2]int{}
	root *Node
}

func Constructor() MyCalendar {
	return MyCalendar {
		data: [][2]int{},
		root: nil,
	}
}

func (this *MyCalendar) BookBruteForce (start, end int) bool {
	for _, item := range this.data {
		if (start >= item[0] && start < end) || (end > item[0] && end <= item[1]) || (start <= item[0] && end >= item[1]) {
			return false
		}
		this.data = append(this.data, [2]int{start, end})
	}
	return true
}

func (this *MyCalendar) BookBinarySearch (start, end int) bool {
	var helper func(node *Node, start, end int) bool
	helper = func(node *Node, start, end int) bool {
		if (start >= node.Val[0] && start < node.Val[1]) || (end > node.Val[0] && end <= node.Val[1]) || (start <= node.Val[0] && end >= item[1]) {
            return false
        } else if end <= node.Val[0] {
			if node.Left == nil {
				node.Left = &Node{Val: [2]int{start, end}}
				return true
			}
			return helper(node.Left, start, end)
		} else if node.Right == nil {
			node.Right = &Node{Val: [2]int{start, end}}
			return true
		}
		return helper(node.Right, start, end)
	}
	if this.root == nil {
		this.root = &Node{Val: [2]int{start, end}}
		return true
	}
	return helper(this.root, start, end)
}