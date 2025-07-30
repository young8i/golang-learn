package main

import "fmt"

func getCommonPrefix(strs []string) string {
	res := make([][]rune, maxStringLength(strs))
	for _, str := range strs {
		for i, s := range str {
			res[i] = append(res[i], s)
		}
	}
	fmt.Println(res)
	returnStr := []rune{}
	for _, r := range res {
		if allCharsSame(r) && len(r) == len(strs) {
			returnStr = append(returnStr, r[0])
		}
	}
	return string(returnStr)
}

func maxStringLength(arr []string) int {
	maxLen := 0
	for _, str := range arr {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}
	return maxLen
}

func allCharsSame(arr []rune) bool {
	if len(arr) == 0 {
		return false
	}
	first := arr[0]
	for _, v := range arr[1:] {
		if v != first {
			return false
		}
	}
	return true
}

/**
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 * 如果不存在公共前缀，返回空字符串 ""。
 */
func main() {
	fmt.Println("result===>", getCommonPrefix([]string{"flower", "flow", "floight"}))
}
