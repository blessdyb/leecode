package main

import "fmt"

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	return &Heap{
		data: []int{},
	}
}

func (this *Heap) Data() []int {
	return this.data
}

func (this *Heap) Push(value int) {
	this.data = append(this.data, value)
	index := len(this.data) - 1
	parentIndex := (index - 1) / 2
	for index >= 0 && this.data[parentIndex] < this.data[index] {
		this.data[index], this.data[parentIndex] = this.data[parentIndex], this.data[index]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

func (this *Heap) Pop() int {
	var ret int
	if len(this.data) > 0 {
		ret = this.data[0]
		this.data[0] = this.data[len(this.data)-1]
		parentIndex := 0
		left := parentIndex*2 + 1
		for left < len(this.data)-1 {
			child := left
			if child+1 < len(this.data)-1 && this.data[child+1] > this.data[child] {
				child++
			}
			if this.data[child] > this.data[parentIndex] {
				this.data[child], this.data[parentIndex] = this.data[parentIndex], this.data[child]
				parentIndex = child
				left = parentIndex*2 + 1
			} else {
				break
			}
		}
		this.data = this.data[:len(this.data)-1]
	}
	return ret
}

func minStoneSum(piles []int, k int) int {
	heap := NewHeap()
	for _, p := range piles {
		heap.Push(p)
	}
	for k > 0 {
		temp := heap.Pop()
		temp -= temp / 2
		heap.Push(temp)
		fmt.Println(temp)
		k--
	}
	ret := 0
	for _, v := range heap.Data() {
		ret += v
	}
	return ret
}
