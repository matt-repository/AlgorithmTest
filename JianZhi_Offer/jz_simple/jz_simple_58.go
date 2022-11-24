package jz_simple

import "strings"

// ||左旋字符串

func ReverseLeftWords(s string, n int) string {
	_n := n % len(s)
	data := []byte(s)
	reverse(data, 0, _n-1)
	reverse(data, _n, len(s)-1)
	reverse(data, 0, len(data)-1)
	return string(data)
}

func reverse(data []byte, left, right int) {
	if left < right {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}
}

// 自己想的，时间复杂度低，但是空间复杂度高
func reverseLeftWords1(s string, n int) string {
	_n := n % len(s)
	res := make([]byte, 0)

	for i := _n; i < len(s); i++ {
		res = append(res, s[i])
	}
	for i := 0; i < _n; i++ {
		res = append(res, s[i])
	}
	return string(res)

}

// 翻转字符串

//输入: "the sky is blue"
//输出: "blue is sky the"
//
//输入: "  hello world!  "
//输出: "world! hello"
//解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//
//
//输入: "a good   example"
//输出: "example good a"
//解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	res := make([]byte, 0)
	right := len(s) - 1
	for j := len(s) - 1; j >= 0; j-- {
		if j == 0 || s[j] == ' ' {
			start := j + 1
			if j == 0 {
				start = j
			}
			for i := start; i <= right; i++ {
				res = append(res, s[i])
			}
			if j != 0 && res[len(res)-1] != ' ' {
				res = append(res, ' ')
			}
			right = j - 1
		}
	}
	return string(res)
}
