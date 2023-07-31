package main

type MyHashSet struct {
	data []int
}

func Constructor() MyHashSet {
	return MyHashSet{
		data: []int{},
	}
}

func (this *MyHashSet) indexOf(key int) int {
	index := -1
	for idx, item := range this.data {
		if item == key {
			index = idx
			break
		}
	}
	return index
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.data = append(this.data, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	index := this.indexOf(key)
	if index != -1 {
		this.data = append(this.data[:index], this.data[index+1:]...)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	return this.indexOf(key) != -1
}
