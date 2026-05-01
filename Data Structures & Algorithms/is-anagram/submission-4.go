func isAnagram(s string, t string) bool {

    if len(s) != len(t){
        return false
    }

    m1 := makeWordMap(s)
    m2 := makeWordMap(t)



    for k,v := range m1 {
       if m2[k] != v {
         return false
       }

    }

    return true



}

func makeWordMap(s string) map[rune]int {
    m := make(map[rune]int)

    for _,v := range s {
        m[v]++
    }

    return m
}
