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
	if desc {
		buildMinHeap(arr)
	} else {
		buildMaxHeap(arr)
	}
	for i := len(arr) - 1; i > 0; i-- {
		swap(&arr[0], &arr[i])
		heapify(arr[:i], 0, !desc)
	}
	return arr
}

// buildMaxHeap
// @Description: 构建大根堆
// @param arr
func buildMaxHeap(arr []int) {
	for i := int(math.Ceil(float64(len(arr))/2)) - 1; i >= 0; i-- {
		heapify(arr, i, true)
	}
}

// buildMinHeap
// @Description: 构建小根堆
// @param arr
func buildMinHeap(arr []int) {
	for i := int(math.Ceil(float64(len(arr))/2)) - 1; i >= 0; i-- {
		heapify(arr, i, false)
	}
}

// heapify
// @Description: 调整堆
// @param arr
// @param root 调整的起点
// @param max true：调整大根堆；false：调整小根堆
func heapify(arr []int, root int, max bool) {
	flag, left, right := root, 2*root+1, 2*root+2
	if left < len(arr) && xor(max, arr[flag] > arr[left]) {
		flag = left
	}
	if right < len(arr) && xor(max, arr[flag] > arr[right]) {
		flag = right
	}
	if flag != root {
		swap(&arr[root], &arr[flag])
		heapify(arr, flag, max)
	}
}
