package main

type SmallestInfiniteSet struct {
	data    []int
	dataMap map[int]bool
}

func Constructor() SmallestInfiniteSet {
	sis := SmallestInfiniteSet{
		data:    []int{},
		dataMap: map[int]bool{},
	}
	for i := 1; i <= 1000; i++ {
		sis.AddBack(i)
	}
	return sis
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var ret int
	if len(this.data) > 0 {
		ret = this.data[0]
		this.data[0] = this.data[len(this.data)-1]
		parentIndex := 0
		left := parentIndex*2 + 1
		for left < len(this.data)-1 {
			child := left
			if child+1 < len(this.data)-1 && this.data[child+1] < this.data[child] {
				child = child + 1
			}
			if this.data[child] < this.data[parentIndex] {
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

func (this *SmallestInfiniteSet) AddBack(num int) {
	if !this.dataMap[num] {
		this.data = append(this.data, num)
		index := len(this.data) - 1
		parentIndex := (index - 1) / 2
		for index >= 0 && this.data[parentIndex] > this.data[index] {
			this.data[parentIndex], this.data[index] = this.data[index], this.data[parentIndex]
			index = parentIndex
			parentIndex = (index - 1) / 2
		}
		this.data[num] = true
	}
}
