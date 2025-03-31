package main

import "fmt"

var RomanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	var testData = []string{
		"IV",      // 4
		"CCXLVI",  // 246
		"MMCDXXI", // 2421
		"M",       // 1000
		"MCMLXIV", // 1964
		"CDXLIV",  // 444
		"LXXXVI",  // 86
	}

	for _, val := range testData {
		fmt.Println(val, convRomanNumerals(val))
	}
}

func convRomanNumerals(code string) int {
	var num int
	var prev int
	var current int

	if len(code) < 1 {
		return num
	} else if len(code) == 1 {
		return RomanNumerals[code]
	}

	for i, val := range code {
		current = RomanNumerals[string(val)]
		if i == 0 {
			num += current
			continue
		}
		prev = RomanNumerals[string(code[i-1])]
		if current <= prev {
			num += current
		} else {
			num += current - 2*prev
		}
	}

	return num
}
