package simple

import "strings"

//LeedCode 557题，反转字符串中的单词Ⅱ
//给定一个字符串 s ，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。
//输入：s = "Let's take LeetCode contest"
//输出："s'teL ekat edoCteeL tsetnoc"

func reverseWords(s string) string {
	length := len(s)
	indexStart := 0
	ret := []byte{} //能用byte 尽量用byte，string 比byte 查询效率差多了
	for i := 0; i <= length; i++ {
		if i == length || s[i] == ' ' {
			for j := i - 1; j >= indexStart; j-- {
				ret = append(ret, s[j])
			}
			indexStart = i + 1
			if i < length-1 {
				ret = append(ret, ' ')
			}
		}
	}
	return string(ret)
}

func ReverseWords1(s string) string {
	sSlice := strings.Split(s, " ")
	result := ""
	for index, str := range sSlice {
		length := len(str)
		for i := length - 1; i > -1; i-- {
			result = result + string(str[i])
		}
		if index < len(sSlice)-1 {
			result = result + " "
		}
	}
	return result
}
