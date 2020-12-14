package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr, iyr, eyr, cid int
	hgt, hcl, ecl, pid string
}

func (p *passport) isValid() bool {
	validFields := p.byr != 0 && p.iyr != 0 && p.eyr != 0 && p.pid != "" && p.hgt != "" && p.hcl != "" && p.ecl != ""
	validByr := p.byr <= 2002 && p.byr >= 1920
	validIyr := p.iyr >= 2010 && p.iyr <= 2020
	validEyr := p.eyr >= 2020 && p.eyr <= 2030
	var validHgt bool
	validEcl := p.ecl == "amb" || p.ecl == "blu" || p.ecl == "brn" || p.ecl == "gry" || p.ecl == "grn" || p.ecl == "hzl" || p.ecl == "oth"

	// validHgt
	r := regexp.MustCompile(`^\d+cm$`)
	if r.MatchString(p.hgt) {
		num, err := strconv.Atoi(strings.ReplaceAll(p.hgt, "cm", ""))
		checkError(err)
		validHgt = num >= 150 && num <= 193
	}

	r = regexp.MustCompile(`^\d+in$`)
	if r.MatchString(p.hgt) {
		num, err := strconv.Atoi(strings.ReplaceAll(p.hgt, "in", ""))
		checkError(err)
		validHgt = num >= 59 && num <= 76
	}

	// validHcl
	r = regexp.MustCompile(`^#([0-9]|[a-f]){6}$`)
	validHcl := r.MatchString(p.hcl)

	// validPid
	r = regexp.MustCompile(`^\d{9}$`)
	validPid := r.MatchString(p.pid)

	return validFields && validByr && validIyr && validEyr && validHgt && validHcl && validEcl && validPid
}

func main() {
	// Add some arg parsing sauce
	if len(os.Args) <= 1 {
		fmt.Println("I need an input file")
		os.Exit(1)
	}

	passports, err := readFile(os.Args[1])
	checkError(err)

	// count
	validPassports := 0

	// fmt.Printf("%#v", passports)
	for _, p := range passports {
		if p.isValid() {
			validPassports++
		}
	}

	fmt.Printf("Number of valid passports: %d\n", validPassports)

}

// Dirty util function for checking errors
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// Read file into array of passport data
// Key:Value indicates we either need a map, or an object of some sort
func readFile(file string) ([]passport, error) {
	fileBytes, err := ioutil.ReadFile(file)

	output := make([]passport, len(strings.Split(string(fileBytes), "\n\n")))
	for number, line := range strings.Split(string(fileBytes), "\n\n") {

		// Skip over empty lines because who needs them?
		if len(line) > 0 {
			// remove those pesky new lines
			cleanRecord := strings.ReplaceAll(line, "\n", " ")
			var p passport
			for _, pair := range strings.Split(cleanRecord, " ") {
				if len(pair) > 0 {
					key := strings.Split(pair, ":")[0]
					value := strings.Split(pair, ":")[1]
					switch key {
					case "byr":
						p.byr, err = strconv.Atoi(value)
						// Give an invalid value - just to show it's not missing
						if err != nil {
							p.byr = 1
						}
					case "iyr":
						p.iyr, err = strconv.Atoi(value)
						if err != nil {
							p.iyr = 1
						}
					case "eyr":
						p.eyr, err = strconv.Atoi(value)
						if err != nil {
							p.eyr = 1
						}
					case "cid":
						p.cid, err = strconv.Atoi(value)
						if err != nil {
							p.cid = 1
						}
					case "pid":
						p.pid = value
					case "hgt":
						p.hgt = value
					case "hcl":
						p.hcl = value
					case "ecl":
						p.ecl = value
					}
				}

			}

			output[number] = p
		}
	}
	return output, err
}
