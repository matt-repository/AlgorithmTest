package jz_general

//股票的最大利润

func maxProfit(prices []int) int {
	max := 0
	count := 0
	for i := 1; i < len(prices); i++ {
		count += prices[i] - prices[i-1]
		if count < 0 {
			count = 0
		}
		if count > max {
			max = count
		}
	}

	return max

}
