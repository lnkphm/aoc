package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func GetSum(arr []string, sorted bool) []int {
	var sum []int
	curr := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			tmp, err := strconv.Atoi(arr[i])
			if err != nil {
				log.Fatal(err)
			}
			curr += tmp
		} else {
			sum = append(sum, curr)
			curr = 0
		}
	}
	// In case no last line
	if curr != 0 {
		sum = append(sum, curr)
	}

	if sorted {
		sort.Ints(sum)
	}

	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, line)
	}

	sums := GetSum(inputs, true)
	top1 := sums[len(sums)-1]
	fmt.Println("Part 1:", top1)
	top3 := sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3]
	fmt.Println("Part 2:", top3)
}
