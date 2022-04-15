package simple

//Leed 第121题 买卖股票的最佳时机
//给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。
//
//你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
//返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

func MaxProfit(prices []int) int {
	max := 0
	current := 0
	for i := 1; i < len(prices); i++ {
		if current+prices[i]-prices[i-1] > 0 {
			current = current + prices[i] - prices[i-1]
		} else {
			current = 0
		}
		if max < current {
			max = current
		}
	}
	return max
}
