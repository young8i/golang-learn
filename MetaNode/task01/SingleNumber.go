package main

import "fmt"

func singleNumber(nums []int) int {

	var result int

	numMap := make(map[int]int, 10)

	for _, num := range nums {
		numMap[num] += 1
	}

	for num, index := range numMap {
		if index == 1 {
			result = num
		}
	}

	return result
}

/**
 * 只出现一次的数字
 */
func main() {
	var nums []int = []int{2, 2, 1, 3, 3, 4, 4, 5, 5}
	fmt.Println("result===>", singleNumber(nums))
}
