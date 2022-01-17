// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 1/15/2022 3:36 PM:00
package sort

// InsertionInt
// @Description: 插入排序
// @param arr
// @param desc
// @return []int
func InsertionInt(arr []int, desc bool) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && (!(arr[j] > key) && desc || arr[j] > key && !desc) {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
