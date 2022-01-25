// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/1/17 17:12:00
package sort

// SelectionInt
// @Description: 选择排序
// @param arr
// @param desc
// @return []int
func SelectionInt(arr []int, desc bool) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		ki := i
		for j := i + 1; j < len(arr); j++ {
			if xor(desc, arr[ki] > arr[j]) {
				ki = j
			}
		}
		if ki != i {
			swap(&arr[ki], &arr[i])
		}
	}
	return arr
}
