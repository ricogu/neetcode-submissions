func maxProfit(prices []int) int {
    maxProfit := 0

    for l,r:=0,0;r<len(prices); r++ {
        profit := prices[r] - prices[l]
        if profit > 0 {
            maxProfit = max(profit,maxProfit)
        } else {
            l=r
        }
    }
    return maxProfit
}

func max(a,b int) int{
    if a>b {
        return a
    }
    return b
}
