package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	LoadPassports("test.txt")
}

type Passport struct {
	data map[string]string
}

func LoadPassports(path string) *[]Passport {
	var passports []Passport

	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	inputStr = strings.TrimSuffix(inputStr, "\n")

	parts := strings.Split(inputStr, "\n\n")

	for _, p := range parts {
		part := strings.Replace(p, "\n", " ", -1)
		fields := strings.Split(part, " ")

		data := make(map[string]string, 8)
		for _, f := range fields {
			v := strings.Split(f, ":")
			data[v[0]] = v[1]
		}
		passports = append(passports, Passport{data: data})
	}

	fmt.Println(passports)
	return &passports
}

func IsValidPassport(passport string) bool {
	isValid := strings.Contains(passport, "byr:") &&
		strings.Contains(passport, "iyr:") &&
		strings.Contains(passport, "eyr:") &&
		strings.Contains(passport, "hgt:") &&
		strings.Contains(passport, "hcl:") &&
		strings.Contains(passport, "ecl:") &&
		strings.Contains(passport, "pid:")

	return isValid
}
