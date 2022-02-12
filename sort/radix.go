// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/2/12 15:22:00
package sort

func RadixInt(arr []int, desc bool) []int {
	maxDigit := 2
	mod, dev := 10, 1

	for maxDigit > 0 {
		buckets := make([][]int, 10)
		for i := range arr {
			bucket := (arr[i] % mod) / dev
			if buckets[bucket] == nil {
				buckets[bucket] = make([]int, 0)
			}
			buckets[bucket] = append(buckets[bucket], arr[i])
		}
		arri := 0
		if desc {
			arri = len(arr) - 1
		}
		for i := range buckets {
			if buckets[i] != nil {
				if desc {
					for j := range buckets[i] {
						arr[arri] = buckets[i][len(buckets[i])-j-1]
						arri--
					}
					continue
				}
				for j := range buckets[i] {
					arr[arri] = buckets[i][j]
					arri++
				}
			}
		}

		maxDigit--
		mod *= 10
		dev *= 10
	}

	return arr
}
