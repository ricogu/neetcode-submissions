func maxProfit(prices []int) int {
    maxProfit := 0
    l,r := 0,1

    for r<len(prices) {
        if prices[l] > prices[r] {
            l = r
        } else {
            profit := prices[r] - prices[l]

            if profit > maxProfit {
                maxProfit = profit
            }
        }

        r++
        
    }

    return maxProfit

}
