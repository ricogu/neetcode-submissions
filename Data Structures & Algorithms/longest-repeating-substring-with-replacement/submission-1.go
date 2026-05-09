func characterReplacement(s string, k int) int {
    counts := make([]int,26)
    maxLength := 0
    maxFreq := 0

    for i,j := 0,0; j<len(s);j++ {
        counts[s[j]-'A']++

        maxFreq = max(maxFreq,counts[s[j]-'A'])
        
        for (j-i+1) - maxFreq > k {
            counts[s[i]-'A']--
            i++
        }

        maxLength = max(maxLength,j-i+1)


    }

    return maxLength

}

func max(a,b int) int{
    if a > b {
        return a
    }

    return b
}
