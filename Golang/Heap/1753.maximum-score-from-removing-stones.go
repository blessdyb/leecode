package main

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	return &Heap{
		data: []int{},
	}
}

func (this *Heap) Len() int {
	return len(this.data)
}

func (this *Heap) Push(value int) {
	this.data = append(this.data, value)
	index := len(this.data) - 1
	parentIndex := (index - 1) / 2
	for index >= 0 && this.data[parentIndex] < this.data[index] {
		this.data[parentIndex], this.data[index] = this.data[index], this.data[parentIndex]
		index = parentIndex
		parentIndex = (index - 1) / 2
	}
}

func (this *Heap) Pop() int {
	var ret int
	if len(this.data) > 0 {
		ret = this.data[0]
		length := len(this.data) - 1
		this.data[0] = this.data[length]
		parentIndex := 0
		left := 2*parentIndex + 1
		for left < length {
			child := left
			if child+1 < length && this.data[child+1] > this.data[child] {
				child++
			}
			if this.data[child] > this.data[parentIndex] {
				this.data[child], this.data[parentIndex] = this.data[parentIndex], this.data[child]
				parentIndex = child
				left = 2*parentIndex + 1
			} else {
				break
			}
		}
		this.data = this.data[:length]
	}
	return ret
}

func maximumScore(a int, b int, c int) int {
	heap := NewHeap()
	heap.Push(a)
	heap.Push(b)
	heap.Push(c)
	ret := 0
	for heap.Len() > 1 {
		t1 := heap.Pop()
		t2 := heap.Pop()
		ret++
		t1--
		t2--
		if t1 > 0 {
			heap.Push(t1)
		}
		if t2 > 0 {
			heap.Push(t2)
		}
	}
	return ret
}
