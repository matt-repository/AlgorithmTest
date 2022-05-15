package jz_general

//表示数值的字符串

func isNumber(s string) bool {
	if s == "" {
		return false
	}
	sStr := &sStruct{
		S: []rune(s),
	}
	trim(sStr)

	result := isInt(sStr)

	if len(sStr.S) > 0 && sStr.S[0] == '.' {
		sStr.S = sStr.S[1:]
		isUnIntResult := isUnInt(sStr)
		result = result || isUnIntResult
	}

	if len(sStr.S) > 0 && (sStr.S[0] == 'e' || sStr.S[0] == 'E') {
		sStr.S = sStr.S[1:]
		result = result && isInt(sStr)
	}
	return result && len(sStr.S) == 0
}

func isUnInt(s *sStruct) bool {
	result := false
	for len(s.S) > 0 && s.S[0] >= '0' && s.S[0] <= '9' {
		result = true
		s.S = s.S[1:]
	}
	return result
}

func isInt(s *sStruct) bool {
	if len(s.S) > 0 && (s.S[0] == '+' || s.S[0] == '-') {
		s.S = s.S[1:]
	}
	return isUnInt(s)
}

func trim(s *sStruct) {
	for len(s.S) > 0 && s.S[0] == ' ' {
		s.S = s.S[1:]
	}
	for len(s.S) > 0 && s.S[len(s.S)-1] == ' ' {
		s.S = s.S[:len(s.S)-1]
	}

}

type sStruct struct {
	S []rune
}
