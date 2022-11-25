package Test

func Test() {

}
func constructArr(a []int) []int {

	if len(a) == 0 {
		return []int{}
	}
	res := make([]int, len(a))

	res[0] = 1
	for i := 1; i < len(a); i++ {
		res[i] = res[i-1] * a[i-1]
	}

	temp := 1
	for i := len(a) - 1; i >= 0; i-- {
		res[i] *= temp
		temp *= a[i]
	}
	return res
}
