package main

func addStrings(num1 string, num2 string) string {
	n1, n2 := []byte(num1), []byte(num2)
	if len(n1) < len(n2) {
		n1, n2 = n2, n1
	}
	carrier := 0
	for i, j := len(n1)-1, len(n2)-1; i >= 0; {
		a := int(n1[i] - '0')
		if j >= 0 {
			b := int(n2[j] - '0')
			a += b
		}
		a += carrier
		if a > 9 {
			a -= 10
			carrier = 1
		} else {
			carrier = 0
		}
		n1[i] = byte(a) + '0'
		i--
		j--
	}
	if carrier == 1 {
		return "1" + string(n1)
	}
	return string(n1)
}
