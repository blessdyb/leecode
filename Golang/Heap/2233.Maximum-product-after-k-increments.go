package main

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maximumProduct(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
	}
	for i := 0; i < k; i++ {
		heap.Push(h, heap.Pop(h).(int)+1)
	}

	result := heap.Pop(h).(int)
	for h.Len() > 0 {
		result = (result * heap.Pop(h).(int)) % (1e9 + 7)

	}

	return result
}

func maximumProductTLE(nums []int, k int) int {
	for k > 0 {
		sort.Ints(nums)
		nums[0]++
		k--
	}
	ret := 1
	for _, num := range nums {
		ret = ret * num % (1000000007)
	}
	return ret
}
