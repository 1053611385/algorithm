package sort

func merge(l []int) []int {
	length := len(l)
	if len(l) < 2 {
		return l
	}
	middle := length / 2
	left := l[0:middle]
	right := l[middle:]
	return mergeSort(merge(left), merge(right))
}

func mergeSort(l, r []int) []int {
	var result = make([]int, len(l)+len(r))
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			result[i+j] = l[i]
			i++
		} else {
			result[i+j] = r[j]
			j++
		}
	}
	if i < len(l) && j == len(r) {
		result = append(result[0:i+j], l[i:]...)
	}
	if j < len(r) && i == len(l) {
		result = append(result[0:i+j], r[j:]...)
	}
	return result
}
