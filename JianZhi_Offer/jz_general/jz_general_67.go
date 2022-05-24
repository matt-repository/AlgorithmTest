package jz_general

import (
	"math"
	"strings"
)

//把字符串转换成整数

func strToInt(str string) int {
	str = strings.Trim(str, " ")
	var result float64 = 0
	isNegative := false
	length := len(str)
	tempStr := []byte{}
	for i := 0; i < length; i++ {
		if i == 0 {
			if str[i] == '-' {
				isNegative = true
			} else if str[i] == '+' {
				continue
			} else if str[i] >= '0' && str[i] <= '9' {
				tempStr = append(tempStr, str[i])
			} else {
				return 0
			}
		} else {
			if str[i] < '0' || str[i] > '9' {
				if len(tempStr) > 0 {
					break
				} else {
					return 0
				}
			}
			tempStr = append(tempStr, str[i])
		}
	}
	for i := 0; i < len(tempStr); i++ {
		result += float64(tempStr[i]-'0') * math.Pow(10.0, float64(len(tempStr)-1-i))
	}
	if isNegative {
		result = -result
	}
	if result > math.MaxInt32 {
		result = math.MaxInt32
	} else if result < math.MinInt32 {
		result = math.MinInt32
	}
	return int(result)
}
