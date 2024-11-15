package validbraces

func ValidBraces(str string) bool {
	brackets := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	stack := []string{}

	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" || string(str[i]) == "{" || string(str[i]) == "[" {
			stack = append(stack, string(str[i]))
			continue
		}

		if len(stack) >= 1 {
			if brackets[string(str[i])] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				continue
			}
		}

		return false
	}
	return len(stack) == 0
}
