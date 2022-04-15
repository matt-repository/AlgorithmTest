package simple

//leed code 387 题 字符串中的第一个唯一字符
//  给定一个字符串 s ，找到 它的第一个不重复的字符，并返回它的索引 。如果不存在，则返回 -1 。

func FirstUniqChar(s string) int {

	mapTimeCache := map[byte]int{} //值 次数
	for i := 0; i < len(s); i++ {
		mapTimeCache[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if mapTimeCache[s[i]] == 1 {
			return i
		}
	}
	return -1
}
