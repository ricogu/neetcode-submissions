func lengthOfLongestSubstring(s string) int {
   window := make(map[byte]int)
   longest := 0

   for i,j:=0,0; j<len(s); j++ {
    if found,ok := window[s[j]];ok && found >=i {
        i = found + 1
    }

    window[s[j]] = j

    longest = max(longest,j-i+1)
   }

   return longest
}

func max(a,b int) int{
    if a>b {
        return a
    }
    return b
}
