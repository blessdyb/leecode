package main

func convertToTitle(columnNumber int) string {
	values := []byte{}
	for columnNumber > 0 {
		columnNumber--
		values = append(values, byte('A'+columnNumber%26))
		columnNumber /= 26
	}
	for i, j := 0, len(values)-1; i < j; {
		values[i], values[j] = values[j], values[i]
		i++
		j--
	}
	return string(values)
}
