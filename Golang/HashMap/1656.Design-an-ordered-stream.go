package main

type OrderedStream struct {
	data    map[int]string
	n       int
	pointer int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		data:    make(map[int]string),
		n:       n,
		pointer: 1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.data[idKey] = value
	ret := []string{}
	for this.pointer <= this.n && this.data[this.pointer] != "" {
		ret = append(ret, this.data[this.pointer])
		this.pointer++
	}
	return ret
}
