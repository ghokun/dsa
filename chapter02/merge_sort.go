package chapter02

func mergeSort(items []int) []int {
	sorted := append([]int{}, items...)
	if len(sorted) <= 1 {
		return sorted
	}
	sort(sorted, 0, len(sorted)-1)
	return sorted
}

func sort(arr []int, beg int, end int) {
	if beg < end {
		mid := int((beg + end) / 2)
		sort(arr, beg, mid)
		sort(arr, mid+1, end)
		merge(arr, beg, mid, end)
	}
}

func merge(arr []int, beg int, mid int, end int) {
	left := append([]int{}, arr[beg:mid+1]...)
	right := append([]int{}, arr[mid+1:end+1]...)

	i := 0
	j := 0
	k := beg
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}
