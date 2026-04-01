func isValid(s string) bool {
    stack := []byte{}
	closeToOpen := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i:=0; i<len(s); i++ {
		if _, exists := closeToOpen[s[i]]; !exists {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) > 0 && stack[len(stack)-1] == closeToOpen[s[i]] {
			stack = stack[0:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
