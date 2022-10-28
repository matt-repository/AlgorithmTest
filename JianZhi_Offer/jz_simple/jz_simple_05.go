package jz_simple

//替换空格

func replaceSpace(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result += "%20"
		} else {
			result += string(s[i])
		}
	}
	return result
}

// 这样写 效率更高，因为减少了  空间的申请。每次一个新的字符串都是开辟了新的空间
func replaceSpace1(s string) string {

	length := 0

	for i := 0; i < len(s); i++ {
		length++
		if s[i] == ' ' {
			length = length + 2
		}
	}

	result := make([]byte, length)

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			result[length-1] = s[i]
			length--
		} else {
			result[length-1] = '0'
			result[length-2] = '2'
			result[length-3] = '%'
			length = length - 3
		}
	}
	return string(result)

}
