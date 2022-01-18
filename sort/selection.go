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
	for i := 0; i < len(arr)-1; i++ {
		ki := i
		for j := i + 1; j < len(arr); j++ {
			if !desc && arr[ki] > arr[j] || desc && arr[ki] < arr[j] {
				ki = j
			}
		}
		if ki != i {
			arr[i] ^= arr[ki]
			arr[ki] ^= arr[i]
			arr[i] ^= arr[ki]
		}
	}
	return arr
}
