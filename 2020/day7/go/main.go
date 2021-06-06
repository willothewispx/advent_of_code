package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	Bags := LoadBags("input.txt")
	var count int

	for color := range *Bags {
		if color != "shiny gold" && CanHoldGold(color, Bags) > 0 {
			count++
		}
	}
	fmt.Printf("Part One: %d.\n", count)
	fmt.Printf("Part Two: %d.\n", CountIndividualBags("shiny gold", Bags))
}

func LoadBags(path string) *map[string]map[string]int {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	Bags := make(map[string]map[string]int)

	for scanner.Scan() {
		color, contains := ReadBag(scanner.Text())
		Bags[color] = contains
	}

	return &Bags
}

func ReadBag(input string) (string, map[string]int) {
	var color string
	contains := make(map[string]int)

	index := strings.Index(input, " bags")
	color = input[0:index]

	re := regexp.MustCompile(`\d\s[a-z]+\s[a-z]+`)
	matches := re.FindAllString(input, -1)

	if len(matches) != 0 {
		for _, match := range matches {
			key := match[strings.Index(match, " ")+1:]
			contains[key], _ = strconv.Atoi(match[0:strings.Index(match, " ")])
		}
	}

	return color, contains
}

func CanHoldGold(color string, bags *map[string]map[string]int) int {
	contains := (*bags)[color]

	if contains == nil {
		return 0
	}

	if _, ok := contains["shiny gold"]; ok {
		return 1
	} else {
		var b int
		for key := range contains {
			b += CanHoldGold(key, bags)
		}
		return b
	}

}

func CountIndividualBags(color string, bags *map[string]map[string]int) int {
	contains := (*bags)[color]

	if contains == nil {
		return 0
	}

	var count int
	for key, value := range contains {
		count += value + value*CountIndividualBags(key, bags)
	}
	return count

}
