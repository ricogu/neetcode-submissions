func isValid(s string) bool {

	var stack []rune

	for _, ch := range s {
		if ch == '[' || ch == '(' || ch == '{' {
			stack = append(stack, ch)
			continue
		}

		if len(stack) == 0 {
			return false
		}

        if (ch == ']' && stack[len(stack)-1] != '[') || (ch == ')' && stack[len(stack)-1] != '(') || (ch == '}' && stack[len(stack)-1] != '{') {
			return false
		}

		stack = stack[:len(stack)-1]

	}

	return len(stack) == 0
    
}
