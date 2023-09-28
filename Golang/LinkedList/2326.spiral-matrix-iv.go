package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ret := make([][]int, m*n)
	for i := 0; i < m; i++ {
		ret[i] = make([]int, n)
		for j := 0; j < n; j++ {
			ret[i][j] = -1
		}
	}
	up, down, left, right := 0, m-1, 0, n-1
	for i := 0; i < m*n; {
		for j := left; j <= right; j++ {
			if head != nil {
				ret[up][j] = head.Val
				head = head.Next
			} else {
				return ret
			}
			i++
		}
		for j := up + 1; j <= down; j++ {
			if head != nil {
				ret[j][right] = head.Val
				head = head.Next
			} else {
				return ret
			}
			i++
		}
		if left != right {
			for j := right - 1; j >= left; j-- {
				if head != nil {
					ret[down][j] = head.Val
					head = head.Next
				} else {
					return ret
				}
				i++
			}
		}
		if up != down {
			for j := down - 1; j > up; j-- {
				if head != nil {
					ret[j][left] = head.Val
					head = head.Next
				} else {
					return ret
				}
				i++
			}
		}
		up++
		down--
		left++
		right--
	}
	return ret
}
