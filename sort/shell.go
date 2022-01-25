// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/1/18 15:51:00
package sort

// ShellInt
// @Description: 希尔排序
// @param arr
// @return []int
func ShellInt(arr []int, desc bool) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			j, key := i-gap, arr[i]
			for j >= 0 && xor(desc, arr[j] > key) {
				arr[j+gap], j = arr[j], j-gap
			}
			arr[j+gap] = key
		}
	}

	return arr
}
