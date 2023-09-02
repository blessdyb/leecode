package main

import "sort"

func maxSubsequenceSortHash(nums []int, k int) []int {
	sorted := append([]int, nums...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] > sorted[j]
	})
	hashmap := map[int]int{}
	for _, num := range sorted[:k] {
		hashmap[num]++
	}
	ret := []int{}
	for _, num := range nums {
		if hashmap[num] > 0 {
			ret = append(ret, num)
			hashmap[num]--
		}
	}
	return ret
}

type Heap struct {
	Data [][2]int
	K    int
}

func NewHeap(nums []int, k int) *Heap {
	heap := &Heap{
		Data: [][2]int{},
		K:    k,
	}
	for idx, num := range nums {
		heap.Push(num, idx)
	}
	return heap
}

func (this *Heap) down() {
	index := 0
	left := 2*index + 1
	for left < this.K {
		child := left
		if child+1 < this.K && this.Data[child+1][0] < this.Data[child][0] {
			child++
		}
		if this.Data[index][0] > this.Data[child][0] {
			this.Data[index], this.Data[child] = this.Data[child], this.Data[index]
			index = child
			left = 2*index + 1
		} else {
			break
		}
	}
}

func (this *Heap) Push(num, idx int) {
	if len(this.Data) < this.K {
		this.Data = append(this.Data, [2]int{num, idx})
		index := len(this.Data) - 1
		parent := (index - 1) / 2
		for index >= 0 && this.Data[index][0] < this.Data[parent][0] {
			this.Data[index], this.Data[parent] = this.Data[parent], this.Data[index]
			index = parent
			parent = (index - 1) / 2
		}
	} else if this.Data[0][0] < num {
		this.Data[0] = [2]int{num, idx}
		this.down()
	}
}

func (this *Heap) Pop() [2]int {
	ret := this.Data[0]
	this.Data[0] = this.Data[len(this.Data)-1]
	this.Data = this.Data[:len(this.Data)-1]
	this.down()
	return ret
}
