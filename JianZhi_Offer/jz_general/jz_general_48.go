package jz_general

func LengthOfLongestSubstring(s string) int {
	cache := map[byte]int{} //index= s[i] value =i
	start := 0
	res := 0
	for i := 0; i < len(s); i++ {
		v := s[i]

		value, ok := cache[v]
		if ok && value >= start { //当第二个b 时，start=2 当第二个a 时，起始应还是2
			// 为啥是大于等于，因为以下 start=value+1 而不是start=value
			start = value + 1
		}

		temp := i - start + 1
		if temp > res {
			res = temp
		}

		cache[v] = i
	}
	return res

}
