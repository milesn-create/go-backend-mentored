package main

func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, simbol := range s {

		if len(stack) == 0 {
			if simbol == ']' || simbol == ')' || simbol == '}' {
				return false
			} else {
				stack = append(stack, simbol)
			}
		} else {
			if simbol == '[' || simbol == '(' || simbol == '{' {
				stack = append(stack, simbol)
			} else if pairs[(simbol)] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]

			} else {
				return false
			}

		}

	}
	if len(stack) != 0 {
		return false
	} else {
		return true
	}
}
