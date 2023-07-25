package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CheckContain(first []int, second []int) bool {
	contained := false
	if (first[0] >= second[0] && first[1] <= second[1]) ||
		(second[0] >= first[0] && second[1] <= first[1]) {
		contained = true
	}
	return contained
}

func CheckOverlap(first []int, second []int) bool {
	contained := false
	if first[0] <= second[1] && second[0] <= first[1] {
		contained = true
	}
	return contained
}

func ConvertToInt(arr []string) []int {
	var res []int
	for _, value := range arr {
		num, _ := strconv.Atoi(value)
		res = append(res, num)
	}
	return res
}

func GetLineData(line string) [][]int {
	var data [][]int
	pair := strings.Split(line, ",")
	first := strings.Split(pair[0], "-")
	second := strings.Split(pair[1], "-")
	data = append(data, ConvertToInt(first))
	data = append(data, ConvertToInt(second))
	return data
}

func Part1(arr []string) int {
	count := 0
	for _, line := range arr {
		data := GetLineData(line)
		first := data[0]
		second := data[1]
		if CheckContain(first, second) {
			count += 1
		}
	}
	return count
}

func Part2(arr []string) int {
	count := 0
	for _, line := range arr {
		data := GetLineData(line)
		first := data[0]
		second := data[1]
		if CheckOverlap(first, second) {
			count += 1
		}
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var arr []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}

	part1 := Part1(arr)
	fmt.Println("Part 1:", part1)

	part2 := Part2(arr)
	fmt.Println("Part 2:", part2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
