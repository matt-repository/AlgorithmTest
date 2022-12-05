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
// 牛逼

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

// 二刷 写的，出现的问题在下面的注释
func firstUniqChar(s string) byte {

	cache := [26]S{}

	for i := 0; i < len(s); i++ {
		v := s[i] - 'a'
		if cache[v].Index == 0 {
			cache[v].Index = i
		}
		cache[v].Count++
	}
	index := -1 // 这里置为-1
	val := -1   // 这里置为-1
	for i := 0; i < 26; i++ {
		if cache[i].Count == 1 {
			if index == -1 {
				index = cache[i].Index
				val = i
			}

			if cache[i].Index < index {
				index = cache[i].Index
				val = i
			}
		}
	}
	if val == -1 { // 这里最开始也没写，出错。要考虑到不存在的情况
		return ' '
	}

	return byte(val + 'a') // 这里是val 不是index ，最开始没写value，用的index，直接错误

}

type S struct {
	Index int
	Count int
}
