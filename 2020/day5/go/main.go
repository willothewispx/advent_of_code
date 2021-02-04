package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bp := LoadBoardingPasses("input.txt")

	sort.Slice(*bp, func(i, j int) bool {
		return (*bp)[i].SeatID < (*bp)[j].SeatID
	})

	fmt.Printf("The highest seat id is %d.\n", (*bp)[len(*bp)-1].SeatID)

	start := (*bp)[0].SeatID

	for i, p := range *bp {
		if p.SeatID-start != int64(i) {
			fmt.Printf("The lost seat id is %d.\n", p.SeatID-1)
			break
		}
	}

}

type BoardingPass struct {
	bsp    string
	SeatID int64
}

func LoadBoardingPasses(path string) *[]BoardingPass {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bpass []BoardingPass

	for scanner.Scan() {
		var p BoardingPass

		p.bsp = scanner.Text()

		replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
		binary := replacer.Replace(p.bsp)

		row, _ := strconv.ParseInt(binary[0:7], 2, 32)
		column, _ := strconv.ParseInt(binary[7:10], 2, 32)

		p.SeatID = row*8 + column

		bpass = append(bpass, p)
	}

	return &bpass
}
