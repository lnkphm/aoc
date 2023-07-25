package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	arr := []string{"A Y", "B C", "C Z"}
	sum := Part1(arr)
	if sum != 15 {
		t.Fatal("Result must be 15")
	}
}

func TestPart2(t *testing.T) {
	arr := []string{"A Y", "B C", "C Z"}
	sum := Part2(arr)
	if sum != 12 {
		t.Fatal("Result must be 12")
	}
}
