package jz_general

func LengthOfLongestSubstring(s string) int {
	cache := map[byte]int{} //index= s[i] value =i
	start := 0
	res := 0
	for i := 0; i < len(s); i++ {
		v := s[i]
		value, ok := cache[v]
		if ok && value >= start { // 这里是value >=start 不是  i
			// 为啥是大于等于，因为当等于的时候 也要重新 置为 start
			start = value + 1
		}
		lenValue := i - start + 1 // 这里的逻辑要写在上面的if 外面 因为每次都要判断
		if lenValue > res {
			res = lenValue
		}
		cache[v] = i
	}
	return res

}
