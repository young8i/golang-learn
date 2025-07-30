package main

import "fmt"

func plusOne(param []int) []int {
	param[len(param)-1] += 1
	for i := len(param) - 1; i >= 0; i-- {
		if param[i] == 10 {
			param[i] = 0
			if i-1 < 0 {
				param = append([]int{1}, param...)
			} else {
				param[i-1] += 1
			}

		}
	}
	return param
}

/**
 * 给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
 */
func main() {
	fmt.Println("result======>", plusOne([]int{1, 2, 3, 4, 2, 9}))
}
