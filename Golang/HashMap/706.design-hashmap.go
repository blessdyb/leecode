package main

const size = 1e6 + 1

type MyHashMap struct {
	values [size]int
}

func NewMyHashMap() MyHashMap {
	values := [size]int{}
	for i := 0; i < size; i++ {
		values[i] = -1
	}
	return MyHashMap{values}
}

func (this *MyHashMap) Put(key, value int) {
	this.values[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.values[key]
}

func (this *MyHashMap) Remove(key int) {
	this.values[key] = -1
}
