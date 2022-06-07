package order

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{10, 5, 100, 50, 30, 12, 2}
	expected := []int{2, 5, 10, 12, 30, 50, 100}

	result := InsertionSort(arr)

	assert.Equal(t, result, expected)
}
