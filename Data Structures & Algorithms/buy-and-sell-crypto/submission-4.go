func maxProfit(prices []int) int {
    maxProfit := 0

    for i,j := 0,1; j<=len(prices)-1;  {
        profit := prices[j] - prices[i]

        if profit < 0 {
            i++
        } else {
            j++
        }

        if profit > maxProfit {
            maxProfit = profit
        }
    }

    return maxProfit

}
