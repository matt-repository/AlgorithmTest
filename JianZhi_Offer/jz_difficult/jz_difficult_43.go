package jz_difficult

// 1~n 中n个整数的十进制表示中1出现的次数

// 分析：
// a=n/base*10 b=n%base cur=(n/base)%10
// cur>1 (a+1) *base   [3101]5[92]
//                      0~3102  *  0~99
// cur==1 a*base+b+1   [310]1[592]
//                      0~310  *  0~999   +1  *     592
// cur<1  a*base       [31]0[1592]
//                      0~30  * 0~9999

func CountDigitOne(n int) int {
	res := 0
	base := 1
	for n >= base {
		a := n / (base * 10)
		b := n % base
		cur := (n / base) % 10

		if cur > 1 {
			res += (a + 1) * base
		} else if cur == 1 {
			res += a*base + b + 1
		} else {
			res += a * base
		}
		base *= 10
	}
	return res
}
