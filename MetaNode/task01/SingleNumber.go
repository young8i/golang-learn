package main

import "fmt"

func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}

/**
 * 只出现一次的数字
 */
func main() {
	var nums []int = []int{2, 2, 1}
	fmt.Println("result===>", singleNumber(nums))
}
