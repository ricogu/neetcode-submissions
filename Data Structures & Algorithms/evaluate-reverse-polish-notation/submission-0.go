
func evalRPN(tokens []string) int {
   var stack []int

   for _, token := range tokens {
	if token != "+" && token != "-" && token != "*" && token != "/" {
	   operand, _ := strconv.Atoi(token)
	   stack = append(stack, operand)
	   continue
	}

	operand_one := stack[len(stack)-2]
	operand_two := stack[len(stack)-1]

	var res int

	switch operator := token; operator {
	case "+": res = operand_one + operand_two
	case "-": res = operand_one - operand_two
	case "*": res = operand_one * operand_two
	case "/": res = operand_one / operand_two


	}

	stack = append(stack[:len(stack)-2],res)
   }

   return stack[len(stack)-1]
}
