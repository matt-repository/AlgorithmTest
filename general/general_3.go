package general

//leed code 第3题 无重复字符的最长子串
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func LengthOfLongestSubstring(s string) int {
	mapCache := map[rune]int{}
	max := 0
	start := 0
	for i, v := range []rune(s) {
		index, ok := mapCache[v]
		if ok && index > start {
			start = index + 1
		}
		if i-start+1 > max {
			max = i - start + 1
		}
		mapCache[v] = i
	}
	return max
}

//类似冒泡了。这个效率比较低
func LengthOfLongestSubstring1(s string) int {
	length := len(s)
	max := 0
	temp := 0
	mapCache := map[uint8]byte{}
	for i, j := 0, 0; i < length && j < length; {
		_, ok := mapCache[s[j]]
		if !ok { //报错。代表无重复
			temp++
		} else { //不报错 代表有重复
			i++
			j = i
			mapCache = map[uint8]byte{}
			temp = 1
		}
		if temp > max {
			max = temp
		}
		mapCache[s[j]] = 1
		j++
	}
	return max
}

func LengthOfLongestSubstring2(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxlength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start { //必须大于
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxlength
}
