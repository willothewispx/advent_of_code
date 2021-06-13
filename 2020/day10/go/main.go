package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Problem One: %d\n", ProblemOne("input.txt"))

}

func ProblemTwo(path string) int {
	// Count the number of unnecessary adapters
	return -1
}

func ProblemOne(path string) int {
	adapters := LoadAdapters(path)
	adapterSet := makeSet(adapters)
	var curr int
	var oneJolt int
	var threeJolt int

	for range *adapters {
		for j := 1; j < 4; j++ {
			if _, ok := (*adapterSet)[curr+j]; ok {
				curr += j
				if j == 1 {
					oneJolt++
				} else {
					threeJolt++
				}
				break
			}
		}
	}
	return oneJolt * threeJolt
}

func makeSet(adapters *[]int) *map[int]struct{} {
	set := make(map[int]struct{})
	for _, adapter := range *adapters {
		set[adapter] = struct{}{}
	}
	return &set
}

func LoadAdapters(path string) *[]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var adapters []int
	var max int

	for scanner.Scan() {
		if adapter, ok := strconv.Atoi(scanner.Text()); ok == nil {
			if adapter > max {
				max = adapter
			}
			adapters = append(adapters, adapter)
		}
	}
	adapters = append(adapters, max+3)

	return &adapters

}
