package chapter02

func insertionSort(items []int) []int{
	sorted := append([]int{}, items...)
	if len(sorted) <= 1 {
		return sorted
	}
	for j := 1; j < len(sorted); j++ {
		key := sorted[j]
		i := j - 1
		for i >= 0 && sorted[i] > key {
			sorted[i+1] = sorted[i]
			i--
		}
		sorted[i+1] = key
	}
	return sorted
}
