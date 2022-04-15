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
