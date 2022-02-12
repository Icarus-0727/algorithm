// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/2/12 14:55:00
package sort

import "math"

func BucketInt(arr []int, bucketSize int, desc bool) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	min, max := arr[0], arr[0]
	for i := range arr {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}

	// 初始化桶
	if bucketSize < 1 {
		bucketSize = 5
	}
	bucketCount := int(math.Floor(float64(max-min)/float64(bucketSize))) + 1
	buckets := make([][]int, bucketCount)
	for i := range buckets {
		buckets[i] = make([]int, 0)
	}

	// 分配元素
	for i := range arr {
		buckets[int(math.Floor(float64(arr[i]-min)/float64(bucketSize)))] = append(buckets[int(math.Floor(float64(arr[i]-min)/float64(bucketSize)))], arr[i])
	}

	arr = make([]int, 0)
	if desc {
		for i := range buckets {
			BubbleInt(buckets[len(buckets)-i-1], desc)
			arr = append(arr, buckets[len(buckets)-i-1]...)
		}
	} else {
		for i := range buckets {
			BubbleInt(buckets[i], desc)
			arr = append(arr, buckets[i]...)
		}
	}

	return arr
}
