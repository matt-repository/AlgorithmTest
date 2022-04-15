package jz_general

//机器人的运动范围

func MovingCount(m int, n int, k int) int {

	if k == 0 {
		return 1
	}

	array := make([][]bool, m)
	for i := range array {
		array[i] = make([]bool, n)
	}
	ans := 1
	array[0][0] = true
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 && j == 0) || getValue(i)+getValue(j) > k {
				continue
			}
			if i-1 > 0 {
				array[i][j] = array[i-1][j]
			}
			if j-1 > 0 {
				array[i][j] = array[i][j-1] || array[i][j]
			}
			if array[i][j] {
				ans++
			}

		}
	}

	return ans
}

func getValue(data int) int {

	return data/100 + (data%100)/10 + data%10

}
