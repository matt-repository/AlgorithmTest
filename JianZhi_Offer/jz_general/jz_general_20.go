package jz_general

import "strings"

//表示数值的字符串

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if s == "" {
		return false
	}
	sStr := &sStruct{
		S: []rune(s),
	}

	result := isInt(sStr)

	if len(sStr.S) > 0 && sStr.S[0] == '.' {
		sStr.S = sStr.S[1:]
		isUnIntResult := isUnInt(sStr)
		result = result || isUnIntResult // 一位数字后面 跟一个点 就够了，所以这里 可以写成 或
	}

	if len(sStr.S) > 0 && (sStr.S[0] == 'e' || sStr.S[0] == 'E') {
		sStr.S = sStr.S[1:]
		result = result && isInt(sStr)
	}
	return result && len(sStr.S) == 0
}

// 有符号 整数
func isInt(s *sStruct) bool {
	if len(s.S) > 0 && (s.S[0] == '+' || s.S[0] == '-') {
		s.S = s.S[1:]
	}
	return isUnInt(s)
}

// 无符号整数
func isUnInt(s *sStruct) bool {
	result := false
	for len(s.S) > 0 && s.S[0] >= '0' && s.S[0] <= '9' {
		result = true
		s.S = s.S[1:]
	}
	return result
}

type sStruct struct {
	S []rune
}
