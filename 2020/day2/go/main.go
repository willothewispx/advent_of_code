package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passwords := LoadPasswords("input.txt")
	oldCount, newCount := 0, 0
	for _, password := range passwords {
		if IsValidPassword(password) {
			oldCount++
		}
		if IsValidPasswordNew(password) {
			newCount++
		}
	}
	fmt.Printf("There are %d valid passwords with respect to the old policy.\n", oldCount)
	fmt.Printf("There are %d valid passwords with respect to the new policy.\n", newCount)
}

type Entry struct {
	lower    int
	upper    int
	letter   string
	password string
}

func LoadPasswords(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var output []string

	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

func LoadEntry(s string) (*Entry, error) {
	r := regexp.MustCompile(`^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

	match := r.FindStringSubmatch(s)

	entry := &Entry{}

	if len(match) != 5 {
		return entry, errors.New("Wrong input.")
	}

	if lower, err := strconv.Atoi(match[1]); err == nil {
		entry.lower = lower
	}

	if upper, err := strconv.Atoi(match[2]); err == nil {
		entry.upper = upper
	}

	entry.letter = match[3]

	entry.password = match[4]

	return entry, nil
}

// old policy
func IsValidPassword(password string) bool {
	entry, err := LoadEntry(password)
	if err != nil {
		log.Fatal(err)
	}

	occurence := strings.Count(entry.password, entry.letter)

	if entry.lower <= occurence && occurence <= entry.upper {
		return true
	}
	return false
}

// new policy
func IsValidPasswordNew(password string) bool {
	entry, err := LoadEntry(password)
	if err != nil {
		log.Fatal(err)
	}

	first := string(entry.password[entry.lower-1])
	second := string(entry.password[entry.upper-1])

	if first == entry.letter && second != entry.letter || second == entry.letter && first != entry.letter {
		return true
	}
	return false
}
