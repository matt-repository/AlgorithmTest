package jz_general

func NthUglyNumber(n int) int {

	index := 1
	for i := 1; ; i++ {
		if isTrue(i) {
			if index == n {
				return i
			}
			index++
		}
	}
	return -1

}

func isTrue(n int) bool {
	if n == 1 {
		return true
	}

	if n%2 == 0 || n%3 == 0 || n%5 == 0 {
		return true
	} else {
		return false
	}

}
