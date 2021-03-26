package quicksort_test

import (
	"algorithms/quicksort"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var unsorted []int
	for i := 0; i <= rand.Intn(50); i++ {
		unsorted = append(unsorted, rand.Intn(500))
	}
	sorted := quicksort.QuickSort(unsorted)
	for j := 1; j < len(sorted); j++ {
		assert.GreaterOrEqual(t, sorted[j], sorted[j-1])
	}
}