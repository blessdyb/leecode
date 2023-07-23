package main

func distance(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

type Heap struct {
	Data     [][]int
	Capacity int
}

func NewHeap(value int) *Heap {
	return &Heap{
		Data:     [][]int{},
		Capacity: value,
	}
}

func (this *Heap) Push(value []int) {
	if len(this.Data) < this.Capacity {
		this.Data = append(this.Data, value)
		index := len(this.Data) - 1
		parentIndex := (index - 1) / 2
		for index >= 0 && distance(this.Data[parentIndex]) < distance(this.Data[index]) {
			this.Data[index], this.Data[parentIndex] = this.Data[parentIndex], this.Data[index]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
	} else if distance(this.Data[0]) > distance(value) {
		this.Data[0] = value
		index := 0
		left := index*2 + 1
		for left < this.Capacity {
			child := left
			if child+1 < this.Capacity && distance(this.Data[child+1]) > distance(this.Data[child]) {
				child++
			}
			if distance(this.Data[index]) < distance(this.Data[child]) {
				this.Data[index], this.Data[child] = this.Data[child], this.Data[index]
				index = child
				left = index*2 + 1
			} else {
				break
			}
		}
	}
}

func kClosest(points [][]int, k int) [][]int {
	heap := NewHeap(k)
	for _, point := range points {
		heap.Push(point)
	}
	return heap.Data
}
