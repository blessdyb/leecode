package main

func canVisitAllRooms(rooms [][]int) bool {
	hashmap := map[int]bool{}
	var helper func(room int)
	helper = func(room int) {
		if !hashmap[room] {
			hashmap[room] = true
			for _, item := range rooms[room] {
				helper(item)
			}
		}
	}
	helper(0)
	for i := 0; i < len(rooms); i++ {
		if !hashmap[i] {
			return false
		}
	}
	return true
}
