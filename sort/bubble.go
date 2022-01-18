// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 1/15/2022 4:09 PM:00
package sort

// BubbleInt
// @Description: 冒泡排序
// @param arr
// @param desc
// @return []int
func BubbleInt(arr []int, desc bool) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	for i := len(arr); i > 0; i-- {
		for j := 1; j < i; j++ {
			if !desc && arr[j-1] > arr[j] || desc && arr[j-1] < arr[j] {
				arr[j-1] ^= arr[j]
				arr[j] ^= arr[j-1]
				arr[j-1] ^= arr[j]
			}
		}
	}

	return arr
}
