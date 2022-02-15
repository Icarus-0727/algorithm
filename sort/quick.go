// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/1/18 16:55:00
package sort

import (
	"sync"
)

func QuickInt(arr []int, desc bool) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	pivot := partition(arr, desc)

	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		QuickInt(arr[:pivot], desc)
		wait.Done()
	}()
	go func() {
		QuickInt(arr[pivot+1:], desc)
		wait.Done()
	}()

	wait.Wait()

	return arr
}

func partition(arr []int, desc bool) int {
	pivot, key := 0, 1
	for i := 1; i < len(arr); i++ {
		if xor(desc, arr[i] < arr[pivot]) {
			swap(&arr[i], &arr[key])
			key++
		}
	}
	swap(&arr[pivot], &arr[key-1])

	return key - 1
}
