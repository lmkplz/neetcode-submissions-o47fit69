func isValid(s string) bool {
    stack := []byte{}

	for i:=0; i<len(s); i++ {
		if s[i]=='(' {
			stack = append(stack, ')')
			continue
		}
		if s[i]=='[' {
			stack = append(stack, ']')
			continue
		}
		if s[i]=='{' {
			stack = append(stack, '}')
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[0:len(stack)-1]
			fmt.Println(stack)
		} else {
			return false
		}
	}

	return len(stack) == 0
}
