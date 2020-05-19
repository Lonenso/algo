package ch07

import (
	"fmt"
	"testing"
)

var a = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

func assert(i, j string) {
	if i != j {
		panic("failed")
	}
}

func Test_quicksort(t *testing.T) {
	quicksort(a, 0, len(a) - 1)
	fmt.Printf("%v", a)
}
