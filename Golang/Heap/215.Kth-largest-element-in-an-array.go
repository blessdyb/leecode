package main

type Heap struct {
	data []int
	k    int
}

func NewHeap(k int) *Heap {
	return &Heap{
		data: []int{},
		k:    k,
	}
}

func (this *Heap) Len() int {
	return len(this.data)
}

func (this *Heap) Top() int {
	return this.data[0]
}

func (this *Heap) Add(value int) {
	if this.Len() < this.k {
		this.data = append(this.data, value)
		index := this.Len() - 1
		parentIndex := (index - 1) / 2
		for index >= 0 && this.data[parentIndex] > this.data[index] {
			this.data[parentIndex], this.data[index] = this.data[index], this.data[parentIndex]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
	} else if this.data[0] < value {
		this.data[0] = value
		parentIndex := 0
		left := parentIndex*2 + 1
		for left < this.k {
			child := left
			if child+1 < this.k && this.data[child] > this.data[child+1] {
				child = child + 1
			}
			if this.data[child] < this.data[parentIndex] {
				this.data[parentIndex], this.data[child] = this.data[child], this.data[parentIndex]
				parentIndex = child
				left = parentIndex*2 + 1
			} else {
				break
			}
		}
	}
}

func findKthLargest(nums []int, k int) int {
	heap := NewHeap(k)
	for _, num := range nums {
		heap.Add(num)
	}
	return heap.Top()
}
