func isAnagram(s string, t string) bool {

    if len(s) != len(t){
        return false
    }

    m1 := makeWordMap(s)
    m2 := makeWordMap(t)

    if len(m1) != len(m2) {
        return false
    }

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
        if _,ok := m[v]; ok {
            m[v]++
            continue
        }

        m[v] = 1
    }

    return m
}
