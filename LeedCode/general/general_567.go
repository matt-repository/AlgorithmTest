package general

//leed code567题，字符串的排列

//给你两个字符串s1和s2 ，写一个函数来判断 s2 是否包含 s1的排列。如果是，返回 true ；否则，返回 false 。
//
//换句话说，s1 的排列之一是 s2 的 子串 。

//输入：s1 = "ab" s2 = "eidbaooo"
//输出：true
//解释：s2 包含 s1 的排列之一 ("ba").

//输入：s1= "ab" s2 = "eidboaoo"
//输出：false

func CheckInclusion(s1 string, s2 string) bool {
	length1 := len(s1)
	length2 := len(s2)
	if length2 < length1 {
		return false
	}
	ct := [26]int{}

	for i := 0; i < length1; i++ {
		ct[s1[i]-'a']++
		ct[s2[i]-'a']--
	}
	diff := 0
	for _, v := range ct {
		if v != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := length1; i < length2; i++ {
		if s2[i] == s2[i-length1] {
			continue
		}
		if ct[s2[i]-'a'] == 0 {
			diff++
		}
		if ct[s2[i-length1]-'a'] == 0 {
			diff++
		}
		ct[s2[i]-'a']--
		ct[s2[i-length1]-'a']++

		//滑动字符串后一个
		if ct[s2[i]-'a'] == 0 {
			diff--
		}
		//滑动字符串前一个
		if ct[s2[i-length1]-'a'] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

//
func CheckInclusion1(s1 string, s2 string) bool {
	length1 := len(s1)
	length2 := len(s2)
	if length2 < length1 {
		return false
	}
	ct1 := [26]int{}
	ct2 := [26]int{}
	for i := 0; i < length1; i++ {
		ct1[s1[i]-'a']++
		ct2[s2[i]-'a']++
	}
	if ct2 == ct1 {
		return true
	}
	for i := length1; i < length2; i++ {
		ct2[s2[i]-'a']++
		ct2[s2[i-length1+1]-'a']--
		if ct2 == ct1 {
			return true
		}
	}
	return false
}
