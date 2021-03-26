package quicksort

func QuickSort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}
	pivot, unsorted := unsorted[0], unsorted[1:]
	var left, right []int
	for _, num := range unsorted {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	var result []int
	result = append(result, QuickSort(left)...)
	result = append(result, pivot)
	result = append(result, QuickSort(right)...)

	return result
}
