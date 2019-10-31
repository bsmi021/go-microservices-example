package tests

import (
	"golang-microservices/mvc/utils"
	"sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getElements(n int) []int{
	result := make([]int, n)
	i := 0
	for j := n -1; j>= 0; j-- {
		result[i] = j
		i++
	}

	return result
}

func TestBubbleSort (t *testing.T){
	// Initialization
	elements := []int { 9, 7, 6, 5, 2}

	// Execution
	utils.BubbleSort(elements)

	// Validation
	assert.NotNil(t, elements)
	assert.EqualValues(t, 5, len(elements))
	assert.EqualValues(t, elements[0], 2)
	assert.EqualValues(t, elements[len(elements)-1], 9)
}

func BenchmarkBubbleSort(b *testing.B){
	els := getElements(10000)
	for i:=0; i<b.N; i++ {
		utils.BubbleSort(els)
	}
}

func BenchmarkNativesort(b *testing.B){
	els := getElements(10000)
	for i:=0; i<b.N; i++ {
		sort.Ints(els)
	}
}