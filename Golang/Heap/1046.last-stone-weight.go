package main

import "sort"

type Heap struct {
	data   []int
	length int
}

func (this *Heap) Len() int {
	return this.length
}

func (this *Heap) Add(value int) {
	this.data = append(this.data, value)
	index := len(this.data) - 1
	parentIndex := (index - 1) / 2
	for index >= 0 && this.data[parentIndex] < this.data[index] {
		this.data[parentIndex], this.data[index] = this.data[index], this.data[parentIndex]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
	this.length++
}

func (this *Heap) Pop() int {
	ret := this.data[0]
	this.data[0] = this.data[this.Len()-1]
	parentIndex := 0
	left := 2*parentIndex + 1
	for left < this.Len()-1 {
		child := left
		if child+1 < len(this.data)-1 && this.data[child] < this.data[child+1] {
			child = child + 1
		}
		if this.data[child] > this.data[parentIndex] {
			this.data[parentIndex], this.data[child] = this.data[child], this.data[parentIndex]
			parentIndex = child
			left = 2*parentIndex + 1
		} else {
			break
		}
	}
	this.data = this.data[:this.Len()-1]
	this.length--
	return ret
}

func lastStoneWeightHeap(stones []int) int {
	heap := &Heap{
		data: []int{},
	}
	for _, stone := range stones {
		heap.Add(stone)
	}
	for heap.Len() > 1 {
		a := heap.Pop()
		b := heap.Pop()
		if a != b {
			heap.Add(a - b)
		}
	}
	return heap.Pop()
}

func lastStoneWeightHeapBuiltIn(stones []int) int {
	return 0
}

func lastStoneWeightHeapSort(stones []int) int {
	sort.Ints(stones)
	length := len(stones)
	for length > 1 {
		y := stones[length-1]
		x := stones[length-2]
		stones = stones[:length-2]
		length -= 2
		if y != x {
			stones = append(stones, y-x)
			sort.Ints(stones)
			length++
		}
	}
	if length == 1 {
		return stones[0]
	}
	return 0
}
