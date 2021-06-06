package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	commands := LoadCommands("input.txt")
	fmt.Printf("Part One: %d\n", PartOne(commands))
	fmt.Printf("Part Two: %d\n", PartTwo(commands))
}

func PartTwo(commands *[]Command) int {
	changed := make(map[int]struct{})

	for i := 0; i < len(*commands); i++ {
		if _, ok := changed[i]; !ok {
			if (*commands)[i].command == "nop" {
				changed[i] = struct{}{}
				tmp := make([]Command, len(*commands))
				copy(tmp, *commands)
				(tmp)[i].command = "jmp"
				if isLoop, acc := checkLoop(&tmp); !isLoop {
					return acc
				}
			} else if (*commands)[i].command == "jmp" {
				changed[i] = struct{}{}
				tmp := make([]Command, len(*commands))
				copy(tmp, *commands)
				(tmp)[i].command = "nop"
				if isLoop, acc := checkLoop(&tmp); !isLoop {
					return acc
				}
			}
		}
	}
	return -1
}

func PartOne(commands *[]Command) int {
	if isLoop, acc := checkLoop(commands); isLoop {
		return acc
	}

	return -1
}

func checkLoop(commands *[]Command) (bool, int) {
	indices := make(map[int]struct{})
	var i int
	var acc int

	for {
		if _, ok := indices[i]; ok {
			return true, acc
		}

		if i == len(*commands) {
			return false, acc
		}

		indices[i] = struct{}{}

		command := (*commands)[i]
		acc, i = command.RunCommand(acc, i)
	}

}

func LoadCommands(path string) *[]Command {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var commands []Command

	for scanner.Scan() {
		commands = append(commands, ReadCommand(scanner.Text()))
	}

	return &commands
}

type Command struct {
	command string
	number  int
}

func ReadCommand(c string) Command {
	var command Command

	command.command = c[0:3]

	number, _ := strconv.Atoi(c[4:])
	command.number = number

	return command

}

func (command Command) RunCommand(acc int, i int) (int, int) {
	switch command.command {
	case "acc":
		acc += command.number
		i++
	case "jmp":
		i += command.number
	case "nop":
		i++
	}

	return acc, i
}
