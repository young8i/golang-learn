package main

import "fmt"

func isValid(s string) bool {

	left := []rune{'{', '[', '('}
	right := []rune{'}', ']', ')'}
	// 切片
	stackLeft := []int{}
	stackRight := []int{}
	for _, v := range s {
		for i := 0; i < len(left); i++ {
			if v == left[i] {
				stackLeft = append(stackLeft, i)
			}
		}
		for i := 0; i < len(right); i++ {
			if v == right[i] {
				stackRight = append(stackRight, i)
			}
		}
	}

	if !(len(stackLeft) == len(stackRight)) {
		return false
	}

	for i := 0; i < len(stackLeft); i++ {
		if !(stackLeft[i] == stackRight[i]) && !(stackLeft[i] == stackRight[len(stackLeft)-i-1]) {
			return false
		}
	}
	return true
}

/**
 * 字符串处理、栈的使用
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
 */
func main() {
	fmt.Println("result===>", isValid("[({})]"))
}
