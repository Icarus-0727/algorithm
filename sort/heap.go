// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/1/25 9:23:00
package sort

import (
	"math"
)

// HeapInt
// @Description: 堆排序
// @param arr 排序数组
// @param desc 降序
// @return []int
func HeapInt(arr []int, desc bool) []int {
	for i := len(arr); i > 1; i-- {
		setLatestElem(arr[:i], desc)
	}

	return arr
}

// setLatestElem
// @Description: 将最大/小的元素交换到最后
// @param arr 操作的数组
// @param desc 降序
// @return []int
func setLatestElem(arr []int, desc bool) {
	for i := int(math.Ceil(float64(len(arr))/2)) - 1; i >= 0; i-- {
		max, left, right := i, 2*i+1, 2*i+2
		if left < len(arr) && xor(desc, arr[max] < arr[left]) {
			max = left
		}
		if right < len(arr) && xor(desc, arr[max] < arr[right]) {
			max = right
		}
		if max != i {
			swap(&arr[i], &arr[max])
		}
	}
	swap(&arr[0], &arr[len(arr)-1])
}
