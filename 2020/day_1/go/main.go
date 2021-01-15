package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	array := ReadInts("input.txt")
	num1, num2 := FindTwoNumbers(array)

	if num1 == -1 {
		log.Fatal("Could not find two numbers that add to 2020")
	}

	fmt.Println("---Part One---")
	fmt.Printf("The two numbers that add to 2020 are %d and %d.\n", num1, num2)
	fmt.Printf("Their product is: %d\n", num1*num2)

	num1, num2, num3 := FindThreeNumbers(array)

	if num1 == -1 {
		log.Fatal("Could not find three numbers that add to 2020")
	}

	fmt.Println("---Part Two---")
	fmt.Printf("The three numbers that add to 2020 are %d, %d, and %d.\n", num1, num2, num3)
	fmt.Printf("Their product is: %d\n", num1*num2*num3)
}

func ReadInts(filePath string) *[]int {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var result []int

	for scanner.Scan() {

		number, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &result
}

func FindTwoNumbers(result *[]int) (int, int) {
	array := ReadInts("input.txt")

	for i := 0; i < len(*array); i++ {
		for j := i + 1; j < len(*array); j++ {
			if (*array)[i]+(*array)[j] == 2020 {
				return (*array)[i], (*array)[j]
			}
		}
	}
	return -1, -1
}

func FindThreeNumbers(result *[]int) (int, int, int) {
	array := ReadInts("input.txt")

	for i := 0; i < len(*array); i++ {
		for j := i + 1; j < len(*array); j++ {
			for k := j + 1; k < len(*array); k++ {
				if (*array)[i]+(*array)[j]+(*array)[k] == 2020 {
					return (*array)[i], (*array)[j], (*array)[k]
				}
			}
		}
	}
	return -1, -1, -1
}
