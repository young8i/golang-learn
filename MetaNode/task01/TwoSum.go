package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := make([]int, 2)
	kv := make(map[int]int)
	for i, v := range nums {
		kv[v] = i
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if target == nums[j]+nums[i] {
				result[0] = i
				result[1] = j
				return result
			}
		}
	}
	return result
}

/**
 * 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
 */
func main() {
	fmt.Println("result======>", twoSum([]int{2, 7, 7, 9}, 9))
}
