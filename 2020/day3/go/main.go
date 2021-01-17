package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	m := LoadMap("input.txt")

	mult := 1

	count := m.GetTrees(1, 1)
	fmt.Printf("For 1 down and 1 right we see exactly %d many trees.\n", count)
	mult *= count

	count = m.GetTrees(3, 1)
	fmt.Printf("For 1 down and 3 right we see exactly %d many trees.\n", count)
	mult *= count

	count = m.GetTrees(5, 1)
	fmt.Printf("For 1 down and 5 right we see exactly %d many trees.\n", count)
	mult *= count

	count = m.GetTrees(7, 1)
	fmt.Printf("For 1 down and 7 right we see exactly %d many trees.\n", count)
	mult *= count

	count = m.GetTrees(1, 2)
	fmt.Printf("For 2 down and 7 right we see exactly %d many trees.\n", count)
	mult *= count

	fmt.Printf("Multiplying each number of trees encountered yields %d.\n", mult)
}

type TreeMap struct {
	topology []string
}

func LoadMap(path string) *TreeMap {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var m TreeMap

	for scanner.Scan() {
		m.topology = append(m.topology, scanner.Text())
	}

	return &m
}

func (m *TreeMap) GetChar(x, y int) string {
	width := len(m.topology[0])

	return string(m.topology[y][x%width])
}

func (m *TreeMap) GetTrees(xshift, yshift int) int {
	max := len(m.topology)
	x, y := 0, 0

	var count int

	for y < max-yshift {
		y += yshift
		x += xshift

		char := m.GetChar(x, y)

		if char == "#" {
			count++
		}
	}

	return count
}
