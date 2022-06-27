package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

	for _, v := range list {

		if v == num {
			return true
		}
	}

	return false
}

func NumInList2(list []int, num int) bool {

	trackMap := make(map[int]int)

	for _, v := range list {
		trackMap[v] = v
	}

	_, ok := trackMap[num]
	if ok {
		return ok
	}

	return false
}
