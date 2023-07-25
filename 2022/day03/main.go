package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

func GetCommon(arr []string) int {
	m := make(map[string]bool)
	ref := arr[0]
	for _, runeValue := range ref {
		m[string(runeValue)] = true
	}
	for i := 1; i < len(arr); i++ {
		tmp := make(map[string]bool)
		for _, runeValue := range arr[i] {
			if _, ok := m[string(runeValue)]; ok {
				tmp[string(runeValue)] = true
			}
		}
		m = tmp
	}
	sum := 0
	for key := range m {
		sum += GetValue(key)
	}
	return sum
}

func GetValue(char string) int {
	var value int
	runeValue, _ := utf8.DecodeRuneInString(char)
	if runeValue >= 65 && runeValue <= 90 {
		value = int(runeValue) - 38
	} else if runeValue >= 97 && runeValue <= 122 {
		value = int(runeValue) - 96
	}
	return value
}

func Part1(arr []string) int {
	sum := 0
	for _, line := range arr {
		middle := len(line) / 2
		lines := []string{line[:middle], line[middle:]}
		share := GetCommon(lines)
		sum += share
	}
	return sum
}

func Part2(arr []string) int {
	var share int
	i := 0
	for i < len(arr) {
		groups := arr[i : i+3]
		share += GetCommon(groups)
		i += 3
	}
	return share
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
