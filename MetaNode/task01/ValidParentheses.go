package main

import "fmt"

func isValid(s string) bool {

	left := []rune{'{', '[', '('}
	right := []rune{'}', ']', ')'}
	// 切片
	stack := []rune{}
	for _, v := range s {
		for i := 0; i < len(left); i++ {
			if v == left[i] {
				stack = append(stack, v)
			}
		}
		fmt.Println("stack======>", stack)
		for i := 0; i < len(right); i++ {
			if v == right[i] {
				if len(stack) == 0 {
					return false
				}
				if left[i] != stack[len(stack)-1] {
					return false
				}
			}
		}
	}
	return true
}

/**
 * 字符串处理、栈的使用
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
 */
func main() {
	fmt.Println("result===>", isValid("[{}]"))
}
