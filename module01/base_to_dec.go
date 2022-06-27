package module01

import (
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
//

/**
1    1    1    0

2^3  2^2  2^1  2^0

this is 1*2^3 + 1*2^2 + 1*2^1 + 1*2^0

**/

func BaseToDec(value string, base int) int {

	retv := 0

	p := len(value) - 1

	// const charset = "0123456789ABCDEF"

	for _, x := range value {

		acc := 1 // declared here in order to always reset acc back to 1 for a fresh initialization of the outer loop
		// to get the base raised to a power
		for i := p; i > 0; i-- {
			acc = acc * base
		}
		p = p - 1 // to enable one count from the next lower power

		d := string(x)

		// multiply value with digit in position
		data := HexaConv(d)
		retv += data * acc

	}

	return retv
}

func HexaConv(num string) int {
	v := 0
	switch num {
	case "A":
		v = 10
	case "B":
		v = 11
	case "C":
		v = 12
	case "D":
		v = 13
	case "E":
		v = 14
	case "F":
		v = 15
	default:
		v, _ = strconv.Atoi(num)
	}

	return v
}

func BaseToDec2(value string, base int) int {
	/**

	1	1	1	0
	2^3	2^2	2^1	2^0

	result = 1*2^3 + 1*2^2 + 1*2^1 + 0*2^0 = 8 + 4 + 2 + 0 = 14

	**/
	res := 0
	power := 0
	acc := 1
	for p := len(value) - 1; p >= 0; p-- {
		data := HexaConv(string(value[p]))

		if power != 0 {
			acc = base * acc
		}
		res = data*acc + res
		power++
	}

	return res

}
func BaseToDec3(value string, base int) int {

	const charset = "0123456789ABCDEF"
	res := 0
	muliplier := 1

	for i := len(value) - 1; i >= 0; i-- { // moving from the rightmost to the leftmost

		index := -1

		for j, char := range charset {

			if char == rune(value[i]) {
				index = j // index count gives the value to rep base value
				break
			}

		}

		res = res + index*muliplier
		muliplier = muliplier * base
	}

	return res
}
