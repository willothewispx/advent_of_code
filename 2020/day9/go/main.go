package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sequence := LoadSequence("input.txt")
	partOne := PartOne(sequence)
	fmt.Printf("Part One: %d\n", partOne)
	partTwo := PartTwo(sequence, partOne)
	fmt.Printf("Part Two: %d\n", partTwo)
}

func LoadSequence(input string) *[]int {
	file, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sequence []int

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		sequence = append(sequence, num)
	}

	return &sequence
}

func isValid(sequence []int, number int) bool {
	var set map[int]struct{}
	set = make(map[int]struct{})

	for _, num := range sequence {
		set[number-num] = struct{}{}
	}
	for _, num := range sequence {
		if _, ok := set[num]; ok {
			return true
		}
	}
	return false
}

func PartOne(sequence *[]int) int {
	for i := 25; i < len(*sequence); i++ {
		if !isValid((*sequence)[i-25:i], (*sequence)[i]) {
			return (*sequence)[i]
		}
	}
	return -1
}

func findContiguous(sequence *[]int, target int) (int, int) {
Outer:
	for start := 0; start < len(*sequence); start++ {
		var sum, big int
		var small = target

		for i := start; i < len(*sequence); i++ {
			val := (*sequence)[i]
			sum += val

			if val > big {
				big = val
			}
			if val < small {
				small = val
			}

			if sum == target {
				return small, big
			} else if sum > target {
				continue Outer
			}
		}
	}
	return -1, -1
}

func PartTwo(sequence *[]int, target int) int {
	small, big := findContiguous(sequence, target)
	return small + big
}
