package module01

import (
	"fmt"
	"strconv"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase2(dec, base int) string {
	var retV string

	// 10%2  = 0 (data)
	// 10/2 = 5 (rem)

	// 5 % 2 = 1 (data)
	// 5/2 = 2

	// 2 % 2 = 0 (data)
	// 2 / 2 = 1

	// 1 % 2 = 1 (data)
	// 1 / 2 = 0

	// 1010

	var data int

	for dec > 0 {
		data = dec % base
		dec = dec / base

		retV = fmt.Sprintf("%s%s", hexaValue(data), retV)

	}

	return retV
}

func hexaValue(n int) string {
	var retV string
	switch n {
	case 10:
		retV = "A"
	case 11:
		retV = "B"
	case 12:
		retV = "C"
	case 13:
		retV = "D"
	case 14:
		retV = "E"
	case 15:
		retV = "F"
	default:
		retV = strconv.Itoa(n)
	}
	return retV
}

func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"

	var retV string

	for dec > 0 {

		rem := dec % base

		retV = fmt.Sprintf("%s%s", string(charset[rem]), retV)

		dec = dec / base //(with floating point data truncated!)

	}

	return retV

}
