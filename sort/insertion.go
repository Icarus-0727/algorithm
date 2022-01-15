package sort

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

func InsertionInt8(arr []int8, desc bool) []int8 {
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

func InsertionInt16(arr []int16, desc bool) []int16 {
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

func InsertionInt32(arr []int32, desc bool) []int32 {
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

func InsertionInt64(arr []int64, desc bool) []int64 {
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

func InsertionUint(arr []uint, desc bool) []uint {
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

func InsertionUint8(arr []uint8, desc bool) []uint8 {
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

func InsertionUint16(arr []uint16, desc bool) []uint16 {
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

func InsertionUint32(arr []uint32, desc bool) []uint32 {
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

func InsertionUint64(arr []uint64, desc bool) []uint64 {
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

func InsertionFloat32(arr []float32, desc bool) []float32 {
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

func InsertionFloat64(arr []float64, desc bool) []float64 {
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
