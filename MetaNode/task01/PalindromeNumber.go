package main

import (
	"fmt"
	"strconv"
)

func PalindromeNumber1(num int) bool {

	numString := []rune(strconv.Itoa(num))

	var asc string
	var desc string

	for i := 0; i < len(numString); i++ {
		asc += string(numString[i])
	}

	for i := len(numString) - 1; i >= 0; i-- {
		desc += string(numString[i])
	}

	a, _ := strconv.Atoi(asc)

	d, _ := strconv.Atoi(desc)

	return a == d
}

func PalindromeNumber2(num int) bool {

	numString := []rune(strconv.Itoa(num))

	n := len(numString)

	for i := 0; i < n/2; i++ {
		if numString[i] != numString[n-i-1] {
			return false
		}
	}

	return true
}

/**
 * 回文数
 * 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
 * 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 */
func main() {
	fmt.Println(PalindromeNumber1(1221))
	fmt.Println(PalindromeNumber2(1221))
}
