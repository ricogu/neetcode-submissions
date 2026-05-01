//3 -> 3
//2 -> 2
//1-> 1

type Pair struct {
    num int
    count int
}

func topKFrequent(nums []int, k int) []int {
    m := make(map[int]int)

    for _, num := range nums {
        m[num]++
    }


    var pairs []Pair

    for k , v := range m {
        pairs = append(pairs, Pair{k,v})
    }

    sort.Slice(pairs, func(i,j int)bool{
        return pairs[i].count > pairs[j].count
    })

    var result []int

    for i:=0; i<= k-1; i++ {
        result = append(result,pairs[i].num)
    }

    return result

}
