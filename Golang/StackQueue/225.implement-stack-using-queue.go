package main

type Queue struct {
	Data   []int
	length int
}

func NewQueue() *Queue {
	return &Queue{
		Data:   []int{},
		length: 0,
	}
}

func (this *Queue) Empty() bool {
	return this.length == 0
}

func (this *Queue) Size() int {
	return this.length
}

func (this *Queue) Peek() int {
	var ret int
	if !this.Empty() {
		ret = this.Data[0]
	}
	return ret
}

func (this *Queue) Push(value int) {
	this.Data = append(this.Data, value)
	this.length++
}

func (this *Queue) Pop() int {
	var ret int
	if !this.Empty() {
		ret = this.Data[0]
		this.length--
		this.Data = this.Data[1:]
	}
	return ret
}

type Stack struct {
	q *Queue
}

func NewStack() *Stack {
	return &Stack{
		q: NewQueue(),
	}
}

func (this *Stack) Push(x int) {
	n := this.q.Size()
	this.q.Push(x)
	for n > 0 {
		this.q.Push(this.q.Pop())
		n--
	}
}

func (this *Stack) Pop() int {
	var ret int
	if !this.Empty() {
		ret = this.q.Pop()
	}
	return ret
}

func (this *Stack) Top() int {
	var ret int
	if !this.Empty() {
		ret = this.q.Peek()
	}
	return ret
}

func (this *Stack) Empty() bool {
	return this.q.Empty()
}
