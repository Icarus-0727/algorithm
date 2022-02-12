// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/2/8 10:41:00
package sort

// CountingInt
// @Description: 计数排序
// @param arr
// @param desc
func CountingInt(arr []int, desc bool) {
	if arr == nil || len(arr) < 2 {
		return
	}

	max := 0
	for i := 1; i < len(arr); i++ {
		if arr[max] < arr[i] {
			max = i
		}
	}

	bucket := make([]int, arr[max]+1)
	for i := range arr {
		bucket[arr[i]]++
	}

	flag := 0
	if desc {
		flag = len(arr) - 1
	}
	for i := range bucket {
		for bucket[i] != 0 {
			arr[flag] = i
			bucket[i]--
			if desc {
				flag--
				continue
			}
			flag++
		}
	}
}
