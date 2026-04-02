func evalRPN(tokens []string) int {
	stack := []int{}
	for i:=0; i<len(tokens); i++ {
		num, err := strconv.Atoi(tokens[i])
		
		// if number, push to stack
		if err == nil {
			stack = append(stack, num)
			fmt.Println(stack)
			continue
		}

		// if operator
		val1 := stack[len(stack)-2]
		val2 := stack[len(stack)-1]
		res := 0

		if tokens[i] == "+" {
			res = val1 + val2
		}
		if tokens[i] == "-" {
			res = val1 - val2
		}
		if tokens[i] == "*" {
			res = val1 * val2
		}
		if tokens[i] == "/" {
			if val2 == 0 {
				return math.MaxInt
			}
			res = val1 / val2
		}
		stack = stack[:len(stack)-2]
		stack = append(stack, res)
	}

	return stack[0]
}
