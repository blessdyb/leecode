package main

type CustomStack struct {
	data []int
	size int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		data: []int{},
		size: maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if len(this.data) < this.size {
		this.data = append(this.data, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(this.data) > 0 {
		ret := this.data[len(this.data)-1]
		this.data = this.data[:len(this.data)-1]
		return ret
	}
	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(this.data); i++ {
		this.data[i] += val
	}
}
