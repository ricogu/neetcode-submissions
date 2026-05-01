//3 -> 3
//2 -> 2
//1-> 1


// 7 -> 2


func topKFrequent(nums []int, k int) []int {
    counts := make(map[int]int)
    

    for _, num := range nums {
        counts[num]++
    }

    freq := make([][]int, len(nums)+1)

    for num , count := range counts {
        freq[count] = append(freq[count],num)
    }

    var result []int

    for i:=len(freq)-1; i > 0; i-- {
        for _, num := range freq[i] {
            result = append(result, num)

            if len(result) == k {
                return result
            }
        }
    }

    return result

}
