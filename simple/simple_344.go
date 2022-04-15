package simple

//leedCode 344题，反转字符串
//输入：s = ["h","e","l","l","o"]
//输出：["o","l","l","e","h"]

func ReverseString(s []byte) {
	length := len(s)
	mid := length / 2
	for i := 0; i < mid; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}
