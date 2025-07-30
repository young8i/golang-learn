package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按区间起始位置升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 用结果切片存合并后的区间
	res := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1] // 当前合并区间的最后一个
		current := intervals[i] // 当前要处理的区间

		// 如果有重叠（即当前的 start <= 上一个 end）
		if current[0] <= last[1] {
			// 合并：更新末尾的 end
			last[1] = m(last[1], current[1])
		} else {
			// 没有重叠，直接添加新的区间
			res = append(res, current)
		}
	}

	return res
}

func m(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * 合并区间
 * 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
 * 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
 */
func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := merge(intervals)
	fmt.Println("合并后的区间：", result)
}
