package main

func countStudents(students, sandwiches []int) int {
	hashmap := map[int]int{}
	for _, student := range students {
		hashmap[student]++
	}
	for len(students) > 0 {
		if students[0] == sandwiches[0] {
			hashmap[students[0]]--
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else if hashmap[students[0]] == 0 {
			break
		} else {
			students = append(students[1:], students[0])
		}
	}
	return len(students)
}
