package utils

import (
	"sort"
)

// BubbleSort sorts ascending []int {8,7,6,5,4}
// []int {4,5,6,7,8}
func BubbleSort(elements []int){
	keepRunning := true
	for keepRunning {
		keepRunning = false
		for i := 0; i < len(elements) -1; i++ {
			if elements[i] > elements[i + 1]{
				elements[i], elements[i + 1] = elements[i + 1], elements[i]
				keepRunning = true
			}
		}
	}

}

// Sort sorts elements based on the length of the given slice
func Sort(els []int){
	if len(els) <1000 {
		BubbleSort(els)
		return
	}

	sort.Ints(els)
	
}