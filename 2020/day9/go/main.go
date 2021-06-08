package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sequence := LoadSequence("input.txt")
	fmt.Printf("Part One: %d\n", PartOne(sequence))

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

func isValidNumber(sequence []int, number int) bool {
	for _, i := range sequence {
		current := i - number
		if current < 0 {
			current = number - i
		}
		for _, j := range sequence {
			if current == j {
				return true
			}
		}
	}
	return false

}

func PartOne(sequence *[]int) int {
	for i := 25; i < len(*sequence); i++ {
		if !isValidNumber((*sequence)[i-25:i], (*sequence)[i]) {
			return (*sequence)[i]
		}
	}
	return -1
}
