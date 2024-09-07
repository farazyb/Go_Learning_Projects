package main

func romanToInt(s string) int {
	var result, temp, c int
	romans := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i := len(s) - 1; i >= 0; i-- {

		c = romans[s[i]]
		if c < temp {
			result = result - c
		} else if c >= temp {
			result = result + c
		}
		temp = c
	}
	return result
}
