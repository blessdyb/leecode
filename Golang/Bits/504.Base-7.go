package main

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	flag := false
	if num < 0 {
		num = -num
		flag = true
	}
	ret := ""
	for num > 0 {
		ret = string(num%7+'0') + ret
		num /= 7
	}
	if flag {
		return "-" + ret
	}
	return ret
}
