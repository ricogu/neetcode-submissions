func lengthOfLongestSubstring(s string) int {
    m := make(map[byte]int)
    maxLength := 0

    for i,j:=0,0; j<len(s); j++{
        m[s[j]]++

        for m[s[j]]>1 {
            m[s[i]]--
            i++
        }

        if (j-i+1) > maxLength {
            maxLength = j-i+1
        }

    }

    return maxLength
}
