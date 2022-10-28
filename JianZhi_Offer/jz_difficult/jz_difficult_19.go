package jz_difficult

//正则表达式匹配

//请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，
//而'*'表示它前面的字符可以出现任意次（含0次）。
//在本题中，匹配是指字符串的所有字符匹配整个模式。例如，
//字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

func IsMatch(s, p string) bool {

	sLength := len(s)
	pLength := len(p)
	f := make([][]bool, sLength+1)
	for i := 0; i < sLength+1; i++ {
		temp := make([]bool, pLength+1)
		f[i] = temp
	}

	for i := 0; i <= sLength; i++ {
		for j := 0; j <= pLength; j++ {
			//分成空正则和非空正则两种
			if j == 0 {
				f[i][j] = i == 0
			} else {
				//非空正则分为两种情况 * 和 非*
				if p[j-1] != '*' {
					if i > 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
						f[i][j] = f[i-1][j-1]
					}
				} else {
					//碰到 * 了，分为看和不看两种情况
					//不看
					if j >= 2 {
						f[i][j] = f[i][j-2]
					}
					//看
					if i >= 1 && j >= 2 && (s[i-1] == p[j-2] || p[j-2] == '.') {
						f[i][j] = f[i][j] || f[i-1][j]
					}
				}
			}
		}
	}
	return f[sLength][pLength]
}
