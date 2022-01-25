// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/1/18 16:55:00
package sort

import "sync"

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
		QuickInt(arr[pivot:], desc)
		wait.Done()
	}()

	wait.Wait()

	return arr
}

func partition(arr []int, desc bool) int {
	key, j := arr[0], len(arr)

	for i := len(arr) - 1; i > 0; i-- {
		if !desc && arr[i] >= key || desc && arr[i] <= key {
			j--
			if i != j {
				arr[i] ^= arr[j]
				arr[j] ^= arr[i]
				arr[i] ^= arr[j]
			}
		}
	}

	j--
	if j != 0 {
		arr[0] ^= arr[j]
		arr[j] ^= arr[0]
		arr[0] ^= arr[j]
	}

	return j + 1
}
