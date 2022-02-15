// Package sort
// @Author Icarus-0727
// @QQ 1067177968
// @CreateAt 2022/1/25 10:15:00
package sort

// swap
// @Description: 交换两个指针指向的元素的值
// @param a
// @param b
func swap(a, b *int) {
	if *a == *b {
		return
	}
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

// xor
// @Description: 自定义 bool 类型的异或运算
// @param a
// @param b
// @return bool
func xor(a, b bool) bool {
	return !a && b || a && !b
}
