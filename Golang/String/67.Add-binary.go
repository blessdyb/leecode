package main

func addBinary(a1 string, b1 string) string {
	a, b := []byte(a1), []byte(b1)
	if len(a) < len(b) {
		a, b = b, a
	}
	var carrier byte = '0'
	for i, j := len(a)-1, len(b)-1; i >= 0; {
		if j >= 0 {
			if b[j] == '1' && a[i] == '1' {
				if carrier == '0' {
					a[i] = '0'
					carrier = '1'
				}
			} else if b[j] == '1' || a[i] == '1' {
				if carrier == '1' {
					a[i] = '0'
					carrier = '1'
				} else {
					a[i] = '1'
				}
			} else {
				a[i] = carrier
				carrier = '0'
			}
		} else {
			if a[i] == '1' {
				if carrier == '1' {
					a[i] = '0'
					carrier = '1'
				}
			} else {
				a[i] = carrier
				carrier = '0'
			}
		}
		i--
		j--
	}
	if carrier == '1' {
		return "1" + string(a)
	}
	return string(a)
}
