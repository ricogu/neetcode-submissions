func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := buildFrequencyMap(s)
	m2 := buildFrequencyMap(t)

	for k,v := range m1{
		if m2[k] != v {
			return false
		}
	}

	return true

}

func buildFrequencyMap(s string) map[rune]int {
	m := make(map[rune]int)

	for _,ch := range s{
		m[ch]++
	}

	return m
}
