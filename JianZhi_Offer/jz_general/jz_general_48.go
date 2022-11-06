package jz_general

func LengthOfLongestSubstring(s string) int {
	cache := map[byte]int{} //index= s[i] value =i
	start := 0
	res := 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		value, ok := cache[v]
		if ok && i > value {
			temp := i - start
			if temp > res {
				res = temp
			}
			start = cache[v] + 1
		}
		cache[v] = i
	}
	return res

}
