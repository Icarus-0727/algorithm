// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 1/15/2022 3:36 PM:00
package sort

import "sync"

// MergeInt
// @Description: 归并排序，递归实现
// @param arr
// @param desc
// @return []int
func MergeInt(arr []int, desc bool) []int {
	// 数组为 nil 或元素少于 2，直接返回
	if arr == nil || len(arr) < 2 {
		return arr
	}

	var a []int
	var b []int
	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		a = MergeInt(arr[:len(arr)/2], desc)
		wait.Done()
	}()
	go func() {
		b = MergeInt(arr[len(arr)/2:], desc)
		wait.Done()
	}()

	wait.Wait()

	return mergeArr(a, b, desc)
}

// mergeArr
// @Description: 将两个有序数组合并成一个有序数组
// @param a
// @param b
// @return []int
func mergeArr(a, b []int, desc bool) []int {
	if a == nil || len(a) == 0 {
		return b
	}
	if b == nil || len(b) == 0 {
		return a
	}

	arr, i, ai, bi := make([]int, len(a)+len(b)), 0, 0, 0
	for ; ai < len(a) && bi < len(b); i++ {
		if xor(desc, a[ai] < b[bi]) {
			arr[i], ai = a[ai], ai+1
		} else {
			arr[i], bi = b[bi], bi+1
		}
	}

	if ai == len(a) {
		a, ai = b, bi
	}

	for ; ai < len(a); i, ai = i+1, ai+1 {
		arr[i] = a[ai]
	}

	return arr
}
