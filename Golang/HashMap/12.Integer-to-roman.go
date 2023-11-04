package main

import "strings"

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	hashmap := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	ret := []string{}
	for num > 0 {
		for i := 0; i < len(values); {
			if num >= values[i] {
				ret = append(ret, hashmap[values[i]])
				num -= values[i]
			} else {
				i++
			}
		}
	}
	return strings.Join(ret, "")
}
