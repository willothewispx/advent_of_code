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
	num1, num2 := FindNumbers(array)

	if num1 == -1 {
		log.Fatal("Could not find two numbers that add to 2020")
	}

	fmt.Printf("The two numbers that add to 2020 are %d and %d.\n", num1, num2)
	fmt.Printf("Their product is: %d", num1*num2)
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

func FindNumbers(result *[]int) (int, int) {
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
