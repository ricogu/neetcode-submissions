func characterReplacement(s string, k int) int {
    counts := make([]int,26)
    maxFreq := 0
    maxLength:= 0

    l,r := 0,0

    for r<len(s) {
        idx := s[r]-'A'
        counts[idx]++

        maxFreq = max(maxFreq, counts[idx])
        

        for (r-l+1) - maxFreq>k {
            counts[s[l]-'A']--
            l++
            
        }

        maxLength = max(maxLength, r-l+1)

        r++
    }

    return maxLength

}

func max(a,b int) int{
    if a > b {
        return a
    }

    return b
}
