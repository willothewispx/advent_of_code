package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	cdfs := LoadCDF("input.txt")
	var countOne, countTwo int

	for _, cdf := range *cdfs {
		countOne += CountAnswers(cdf)
		countTwo += CountAnsweredByEveryone(cdf)
	}
	fmt.Printf("Part One: %d.\n", countOne)
	fmt.Printf("Part Two: %d.\n", countTwo)
}

func LoadCDF(path string) *[][]string {
	var cdfs [][]string

	inputFile, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	input := string(inputFile)
	input = strings.TrimSuffix(input, "\n")

	grouped := strings.Split(input, "\n\n")

	for _, a := range grouped {
		parts := strings.Split(a, "\n")
		cdfs = append(cdfs, parts)
	}

	return &cdfs
}

func CountAnswers(answers []string) int {
	all := strings.Join(answers, "")
	var count int

	for {
		if len(all) == 0 {
			return count
		}

		answer := string(all[0])
		all = strings.ReplaceAll(all, answer, "")
		count++
	}

}

func CountAnsweredByEveryone(answers []string) int {
	var count int
	first := answers[0]

OUTER:
	for {
		if len(first) == 0 {
			return count
		}

		answer := string(first[0])
		first = strings.ReplaceAll(first, answer, "")

		for _, a := range answers {
			if !strings.Contains(a, answer) {
				continue OUTER
			}
		}

		count++
	}
}
