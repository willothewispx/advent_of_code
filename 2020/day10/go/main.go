package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	one, adapters := ProblemOne("input.txt")
	fmt.Printf("Problem One: %d\n", one)
	two := ProblemTwo(adapters)
	fmt.Printf("Problem Two: %d\n", two)
}

// Possible combinations (difference of two does not occur):
// 3 - 3 (e.g. 4, 7, 10) => only 1 possible configuration
// 3 - 1 - 3 (e.g. 4, 7, 8, 11) => only 1 possible configuration
// 3 - 1 - 1 - 3 (e.g. 4, 7, 8, 9, 12) => 2 possible configurations (drop the middle number or don't)
// 3 - 1 - 1 - 1 - 3 (e.g. 4, 7, 8, 9, 10, 13) => 4 possible configurations (2 each for 8 and 9)
// 3 - 1 - 1 - 1 - 1 - 3 (e.g. 4, 7, 8, 9, 10, 11, 14) => 7 possible configurations (at least one of 8, 9, 10 must be present, so 2^3 - 1 since dropping all 3 not an option)
// There are no other configurations
func ProblemTwo(adapters *[]int) int {
	mult := 1
	var counter int
	var curr int

	for _, adapter := range *adapters {
		if adapter-curr == 1 {
			counter++
			curr += 1
		} else if adapter-curr == 3 {
			switch counter {
			case 2:
				mult *= 2
			case 3:
				mult *= 4
			case 4:
				mult *= 7
			}
			counter = 0
			curr += 3
		}
	}
	return mult
}

func ProblemOne(path string) (int, *[]int) {
	adapters := LoadAdapters(path)
	adapterSet := makeSet(adapters)
	var sorted []int
	var curr int
	var oneJolt int
	var threeJolt int

	for range *adapters {
		for j := 1; j < 4; j++ {
			if _, ok := (*adapterSet)[curr+j]; ok {
				curr += j
				sorted = append(sorted, curr)
				if j == 1 {
					oneJolt++
				} else {
					threeJolt++
				}
				break
			}
		}
	}
	return oneJolt * threeJolt, &sorted
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
