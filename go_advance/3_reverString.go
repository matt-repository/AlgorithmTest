package go_advance

//3. 翻转字符串
//问题描述
//请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字符串(可以使⽤单个过程变量)。
//给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于5000。
//解题思路
//翻转字符串其实是将⼀个字符串以中间字符为轴，前后翻转，即将str[len]赋值给str[0],将str[0] 赋值 str[len]。

func reverString(str string) string {
	strByte := []byte(str)
	len := len(strByte)
	for i := 0; i < len/2; i++ {
		strByte[i], strByte[len-i-1] = strByte[len-i-1], strByte[i]
	}
	return string(strByte)
}
