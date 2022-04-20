package general

//字母大小写全排列

//输入：s = "a1b2"
//输出：["a1b2", "a1B2", "A1b2", "A1B2"]

//回溯解决
func letterCasePermutation(s string) []string {
	temp := &Temp{
		Queue:  []byte(s),
		Result: []string{},
	}
	dfs_784(len(s), 0, temp)
	return temp.Result
}

func dfs_784(n, first int, temp *Temp) {
	if len(temp.Queue) == n {
		tempQueue := make([]byte, n)
		copy(tempQueue, temp.Queue)
		temp.Result = append(temp.Result, string(tempQueue))
	}

	for i := first; i < n; i++ {
		istrue, v := strTrans(temp.Queue[i])
		if istrue {
			temp.Queue[i] = v
			dfs_784(n, i+1, temp)
			_, v = strTrans(temp.Queue[i])
			temp.Queue[i] = v
		}
	}
}

type Temp struct {
	Queue  []byte
	Result []string
}

func strTrans(value byte) (bool, byte) {
	tempValue1 := value - 'a'
	tempValue2 := value - 'A'

	if tempValue1 >= 0 && tempValue1 <= 25 {
		return true, 'A' + tempValue1
	}
	if tempValue2 >= 0 && tempValue2 <= 25 {
		return true, 'a' + tempValue2
	}
	return false, '0'
}
