package main

func nextGreatestLetter(letters []byte, target byte) byte {
	ret := letters[0]
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			ret = letters[i]
			break
		}
	}
	return ret
}
