package jz_simple

// 第一个出现的小写字母

// 自己思考的

func FirstUniqChar(s string) byte {
	data := make([]*tempStruct, 26)
	for i := 0; i < len(s); i++ {
		num := s[i] - 'a'
		v := data[num]
		if v == nil {
			data[num] = &tempStruct{
				Index: i,
				Count: 1,
			}
		} else {
			v.Count++
		}
	}
	res := ' '
	index := len(s)
	for i := 0; i < len(data); i++ {
		if data[i] != nil && data[i].Count == 1 {
			if data[i].Index < index {
				index = data[i].Index //我个傻逼上次写 忘了这一步，死都检查不出来。记录一下
				res = rune(i + int('a'))
			}
		}
	}
	return byte(res)
}

type tempStruct struct {
	Index int
	Count int
}

// 标准方法

func FirstUniqChar2(s string) byte {
	data := [26]int{}

	for i := 0; i < len(s); i++ {
		data[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if data[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
