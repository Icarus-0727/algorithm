// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/1/25 9:23:00
package sort

// HeapInt
// @Description: 堆排序
// @param arr 排序数组
// @param desc 降序
// @return []int
func HeapInt(arr []int, desc bool) []int {
	buildMinOrMaxHeap(arr, !desc)
	for i := len(arr) - 1; i > 0; i-- {
		swap(&arr[0], &arr[i])
		heapify(arr[:i], 0, !desc)
	}
	return arr
}

// buildMinOrMaxHeap
// @Description: 构建大根堆或小根堆
// @param arr
// @param isMax true: 构建大根堆；false: 构建小根堆
func buildMinOrMaxHeap(arr []int, isMax bool) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, i, isMax)
	}
}

// heapify
// @Description: 调整堆
// @param arr
// @param root 调整的起点
// @param isMax true：调整大根堆；false：调整小根堆
func heapify(arr []int, root int, isMax bool) {
	flag := root
	left := 2*root + 1
	right := left + 1

	// 判断是否需要调整
	if left < len(arr) && xor(isMax, arr[flag] > arr[left]) {
		flag = left
	}
	if right < len(arr) && xor(isMax, arr[flag] > arr[right]) {
		flag = right
	}

	if flag != root {
		swap(&arr[root], &arr[flag]) // 调整
		heapify(arr, flag, isMax)    // 对子树重新调整
	}
}
