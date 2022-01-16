package chapter02

import (
	"testing"

	arrays "github.com/ghokun/dsa/arrays"
)

func TestInsertionSort(t *testing.T) {

	runs := [][]int{{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, {2, 1, 5, 4, 3}}

	for _, run := range runs {
		sorted := insertionSort(run)
		if !arrays.Equal(sorted, []int{1, 2, 3, 4, 5}) {
			t.Errorf("insertionSort(%v) == %v, want %v", run, sorted, []int{1, 2, 3, 4, 5})
		}
	}
}
