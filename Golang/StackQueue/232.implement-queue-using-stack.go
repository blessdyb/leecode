package main

type Stack struct {
	data   []int
	length int
}

func NewStack() *Stack {
	return &Stack{
		data:   []int{},
		length: 0,
	}
}

func (this *Stack) Empty() bool {
	return this.length == 0
}

func (this *Stack) Pop() int {
	var ret int
	if !this.Empty() {
		ret = this.data[this.length-1]
		this.data = this.data[:this.length-1]
	}
	return ret
}

func (this *Stack) Peek() int {
	var ret int
	if !this.Empty() {
		ret = this.data[this.length-1]
	}
	return ret
}

func (this *Stack) Push(value int) {
	this.data = append(this.data, value)
	this.length++
}

type Queue struct {
	s1 *Stack
	s2 *Stack
}

func NewQueue() *Queue {
	return &Queue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (this *Queue) Empty() bool {
	return this.s1.Empty() && this.s2.Empty()
}

func (this *Queue) Peek() int {
	if this.s1.Empty() {
		for !this.s2.Empty() {
			this.s1.Push(this.s2.Pop())
		}
	}
	return this.s1.Peek()
}

func (this *Queue) Push(value int) {
	this.s2.Push(value)
}

func (this *Queue) Pop() int {
	this.Peek()
	return this.s1.Pop()
}
