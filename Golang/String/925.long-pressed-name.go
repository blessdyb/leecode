package main

func isLongPressedName(name, typed string) bool {
	i, j := 0, 0
	for i < len(name) || j < len(typed) {
		if i < len(name) && j < len(name) && name[i] == typed[j] {
			i++
			j++
		} else if j > 0 && j < len(name) && typed[j] == typed[j-1] {
			j++
		} else {
			return false
		}
	}
	return i == len(name) && j == len(typed)
}
