package main

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	arr := []string{"100", "", "100", "200", "", "300", "200"}
	sums := GetSum(arr, true)
	if len(sums) != 3 {
		t.Fatal("Result length must be 3")
	}
	if sums[0] != 100 || sums[len(sums)-1] != 500 {
		t.Fatal("Result must be sorted")
	}
}
