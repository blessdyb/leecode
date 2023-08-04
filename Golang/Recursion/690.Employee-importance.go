package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	hashmap := map[int]*Employee{}
	for _, e := range employees {
		hashmap[e.Id] = e
	}
	var helper func(id int) int
	helper = func(id int) int {
		ret := hashmap[id].Importance
		for _, s := range hashmap[id].Subordinates {
			ret += helper(s)
		}
		return ret
	}
	return helper(id)
}
