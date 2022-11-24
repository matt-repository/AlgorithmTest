package Test

func Test() {

}
func dicesProbability(k int) []float64 {
	res := make([]float64, 6)
	for i := 0; i < 6; i++ {
		res[i] = 1 / 6
	}
	for n := 2; n <= k; n++ {
		temp := make([]float64, 5*n+1)
		for i := 0; i < len(res); i++ {
			for j := 0; j < 6; j++ {
				temp[i+j] += res[i] / 6
			}
		}
		res := make([]float64, 5*n+1)
		copy(res, temp)
	}
	return res

}
