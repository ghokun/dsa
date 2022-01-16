package chapter02

import (
	"testing"

	arrays "github.com/ghokun/dsa/arrays"
)

func TestMergeSort(t *testing.T) {

	runs := [][]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {2, 1, 5, 4, 3}}

	for _, run := range runs {
		sorted := mergeSort(run)
		if !arrays.Equal(sorted, []int{1, 2, 3, 4, 5}) {
			t.Errorf("mergeSort(%v) == %v, want %v", run, sorted, []int{1, 2, 3, 4, 5})
		}
	}
}
