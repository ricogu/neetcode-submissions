//3 -> 3
//2 -> 2
//1-> 1

// [nil,[1],[2],[3]]

func topKFrequent(nums []int, k int) []int {
   m := make(map[int]int)
   freq := make([][]int, len(nums)+1)

   for _, num := range nums {
      m[num]++
   }

   for num, occurance := range m {
    freq[occurance] = append(freq[occurance],num)
   } 

   var res []int

   for idx := len(freq)-1; idx >0 ; idx-- {
     for _, num :=  range freq[idx]{
        res = append(res, num)
        if len(res) == k {
            return res
        }
     }
   }

   return res


}
