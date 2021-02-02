package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	passports := LoadPassports("input.txt")

	var count int
	var isValid int

	for _, passport := range *passports {
		if passport.HasEnoughEntries() {
			count++
			if passport.IsValidPassport() {
				isValid++
			}
		}
	}

	fmt.Printf("There are %d many passports with the required amount of fields.\nAmong those, there are %d many passports which are valid.\n", count, isValid)

}

type Passport struct {
	data map[string]string
}

var RequiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

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

	return &passports
}

func (passport Passport) HasEnoughEntries() bool {
	for _, required := range RequiredFields {
		_, ok := passport.data[required]
		if !ok {
			return false
		}
	}

	return true
}

func (passport Passport) IsValidPassport() bool {
	for _, required := range RequiredFields {
		key, _ := passport.data[required]

		switch required {
		case "byr":
			if !ValidateBYR(key) {
				return false
			}
		case "iyr":
			if !ValidateIYR(key) {
				return false
			}
		case "eyr":
			if !ValidateEYR(key) {
				return false
			}
		case "hgt":
			if !ValidateHGT(key) {
				return false
			}
		case "hcl":
			if !ValidateHCL(key) {
				return false
			}
		case "ecl":
			if !ValidateECL(key) {
				return false
			}
		case "pid":
			if !ValidatePID(key) {
				return false
			}
		}
	}

	return true
}

func ValidateBYR(byr string) bool {
	re := regexp.MustCompile(`^\d{4}$`)

	if re.MatchString(byr) {
		year, _ := strconv.Atoi(byr)
		if 1920 <= year && year <= 2002 {
			return true
		}
	}

	return false
}

func ValidateIYR(iyr string) bool {
	re := regexp.MustCompile(`^\d{4}$`)

	if re.MatchString(iyr) {
		year, _ := strconv.Atoi(iyr)
		if 2010 <= year && year <= 2020 {
			return true
		}
	}

	return false
}

func ValidateEYR(eyr string) bool {
	re := regexp.MustCompile(`^\d{4}$`)

	if re.MatchString(eyr) {
		year, _ := strconv.Atoi(eyr)
		if 2020 <= year && year <= 2030 {
			return true
		}
	}

	return false
}

func ValidateHGT(hgt string) bool {
	re := regexp.MustCompile(`^(\d+)(cm|in)$`)

	if re.MatchString(hgt) {
		sub := re.FindStringSubmatch(hgt)
		measure := sub[2]
		height, _ := strconv.Atoi(sub[1])

		switch measure {
		case "cm":
			if 150 <= height && height <= 193 {
				return true
			}
		case "in":
			if 59 <= height && height <= 76 {
				return true
			}
		}
	}

	return false
}

func ValidateHCL(hcl string) bool {
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)

	if re.MatchString(hcl) {
		return true
	}

	return false
}

func ValidateECL(ecl string) bool {
	re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)

	if re.MatchString(ecl) {
		return true
	}

	return false
}

func ValidatePID(pid string) bool {
	re := regexp.MustCompile(`^\d{9}$`)

	if re.MatchString(pid) {
		return true
	}

	return false
}
