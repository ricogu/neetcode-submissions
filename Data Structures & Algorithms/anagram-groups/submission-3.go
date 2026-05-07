func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)

	for _, str := range strs {
		var key [26]int

		for _, ch := range str {
			key[ch-'a']++
		}

		m[key] = append(m[key], str)
	}

	var res [][]string

	for _, strs := range m {
		res = append(res, strs)
	}

	return res
}
