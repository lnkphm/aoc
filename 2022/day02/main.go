package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1(arr []string) int {
	sum := 0
	points := [][]int{
		{4, 8, 3},
		{1, 5, 9},
		{7, 2, 6},
	}
	oppMoves := map[string]int{"A": 0, "B": 1, "C": 2}
	myMoves := map[string]int{"X": 0, "Y": 1, "Z": 2}

	for _, round := range arr {
		opp := string(round[0])
		my := string(round[2])
		sum += points[oppMoves[opp]][myMoves[my]]
	}
	return sum
}

func Part2(arr []string) int {
	sum := 0

	points := [][]int{
		{3, 4, 8},
		{1, 5, 9},
		{2, 6, 7},
	}
	oppMoves := map[string]int{"A": 0, "B": 1, "C": 2}
	myMoves := map[string]int{"X": 0, "Y": 1, "Z": 2}

	for _, round := range arr {
		opp := string(round[0])
		my := string(round[2])
		sum += points[oppMoves[opp]][myMoves[my]]
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []string
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}
	fmt.Println("Part1:", Part1(arr))
	fmt.Println("Part2:", Part2(arr))

}
