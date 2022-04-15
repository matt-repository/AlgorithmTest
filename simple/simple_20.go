package simple

//leedCode 20题. 有效的括号
//
//输入：s = "()"
//输出：true

//输入：s = "()[]{}"
//输出：true

//利用队列先进先出的思路
func isValid(s string) bool {

	mapCache := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}

	for _, v := range []rune(s) {
		_, ok := mapCache[v]
		if ok {
			stack = append(stack, v)
		} else {
			if len(stack) < 1 {
				return false
			}
			lastV, ok1 := mapCache[stack[len(stack)-1]]
			if !ok1 {
				return false
			}
			if v != lastV {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

//自己写的思路
func isValid1(s string) bool {
	data := []rune(s)

	for i := 0; i < len(data); {
		if i+1 < len(data) && jude(data[i], data[i+1]) {

			data = append(data[0:i], data[i+2:len(data)]...)
			i = 0
		} else if i-1 > 0 && jude(data[i-1], data[i]) {

			data = append(data[0:i-1], data[i+1:len(data)]...)
			i = 0
		} else {
			i++
		}

	}

	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

func jude(s1 rune, s2 rune) bool {
	if s1 == '(' && s2 == ')' {
		return true
	}
	if s1 == '[' && s2 == ']' {
		return true
	}
	if s1 == '{' && s2 == '}' {
		return true
	}
	return false
}
