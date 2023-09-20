package main

type BrowserHistory struct {
	data    []string
	pointer int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		data:    []string{homepage},
		pointer: 0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	this.data = this.data[:this.pointer+1]
	this.pointer++
	this.data = append(this.data, url)
}

func (this *BrowserHistory) Back(steps int) string {
	if this.pointer-steps >= 0 {
		this.pointer -= steps
	} else {
		this.pointer = 0
	}
	return this.data[this.pointer]
}

func (this *BrowserHistory) Forward(steps int) string {
	if this.pointer+steps < len(this.data) {
		this.pointer += steps
	} else {
		this.pointer = len(this.data) - 1
	}
	return this.data[this.pointer]
}
