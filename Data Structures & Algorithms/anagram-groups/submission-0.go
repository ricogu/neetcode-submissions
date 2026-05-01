func groupAnagrams(strs []string) [][]string {

	m := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		m[sortedStr] = append(m[sortedStr],str)
	}

	var result [][]string

	for _, strlist := range m {
		result = append(result, strlist)
		
	}
  
	return result



}

func sortString(s string) string{
	b := []byte(s)

	sort.Slice(b, func(i,j int) bool {
			return b[i] < b[j]
	})
   
	return string(b)

}
