package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	arr := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	sum := Part1(arr)
	if sum != 2 {
		t.Fatal("Result must be 2")
	}
}

func TestPart2(t *testing.T) {
	arr := []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
	sum := Part2(arr)
	if sum != 4 {
		t.Fatal("Result must be 4")
	}
}
