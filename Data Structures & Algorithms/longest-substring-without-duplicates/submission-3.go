func lengthOfLongestSubstring(s string) int {



   window := make(map[byte]bool)
   longest := 0

   for l,r:=0,0; r<len(s); r++ {
        for window[s[r]] {
            delete(window, s[l])
            l++
        }
        
        window[s[r]] = true
        longest = max(longest, r-l+1)
    
   }

   return longest

}

func max(a,b int) int{
    if a>b {
        return a
    }
    return b
}
