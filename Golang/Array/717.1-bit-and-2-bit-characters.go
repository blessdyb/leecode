package main

func isOneBitCharacter(bits []int) bool {
	temp := -1
	for i := 0; i < len(bits)-1; i++ {
		if temp == -1 {
			temp = bits[i]
		} else if bits[i] == 1 && temp == 0 {
			temp = 1
		} else {
			temp = -1
		}
	}
	return temp == 0 || temp == -1
}
