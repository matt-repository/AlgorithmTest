package jz_simple

// ||左旋字符串

func reverseLeftWords(s string, n int) string {
	result := ""
	length := len(s)
	if length == 0 {
		return s
	}
	n_ := n % length

	for i := n_; i < length+n_; i++ {
		result += string(s[i%length])
	}

	return result

}
