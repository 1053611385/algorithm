package sort

func Quick(arr []int) {
	_quick(arr, 0, len(arr)-1)
}

func _quick(arr []int, left, right int) {
	if left < right {
		index := quickSort(arr, left, right)
		_quick(arr, left, index-1)
		_quick(arr, index+1, right)
	}
}

// quick sort
func quickSort(arr []int, left, right int) int {
	baseIndex, i, j := left, left+1, left+1
	for i <= right {
		if arr[i] < arr[baseIndex] {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
		i++
	}
	arr[baseIndex], arr[j-1] = arr[j-1], arr[baseIndex]
	return j - 1
}
