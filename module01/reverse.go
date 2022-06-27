package module01

import "strings"

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {

	var retVal string

	for _, r := range word {
		retVal = string(r) + retVal
	}

	return retVal
}

func Reverse2(word string) string {
	var sb strings.Builder
	for i := len(word) - 1; i >= 0; i-- {
		sb.WriteByte(word[i]) // collects bytes
	}
	return sb.String()
}

func Reverse3(word string) string {
	rev := ""
	for i := len(word) - 1; i >= 0; i-- {
		rev = rev + string(word[i])
	}
	return rev
}
