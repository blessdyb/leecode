package main

type Foo struct {
	ch1, ch2 chan struct{}
}

func NewFoo() *Foo {
	return &Foo{
		ch1: make(chan struct{}),
		ch2: make(chan struct{}),
	}
}

func (this *Foo) first(printFirst func()) {
	printFirst()
	close(this.ch1)
}

func (this *Foo) second(printSecond func()) {
	<-this.ch1
	printSecond()
	close(this.ch2)
}

func (this *Foo) third(printThird func()) {
	<-this.ch2
	printThird()
}
