package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    // Add some arg parsing sauce
    if len(os.Args) <= 1 {
        fmt.Println("I need an input file")
        os.Exit(1)
    }

    // Counters
    validPasswords1 := 0
    validPasswords2 := 0

    // Regex FTW!
    r := regexp.MustCompile(`(\d+)-(\d+) ([a-z]): ([a-z]+)`)

    entries, err := readFile(os.Args[1])
    checkError(err)

    // For each line, parse the rule
    // apply the rule, adjust the vaild count
    for _, line := range entries {
        parts := strings.Split(line, " ")
        if len(parts) == 3 {
            rgx := r.FindStringSubmatch(line)
            // fmt.Printf("%#v\n", rgx)

            // Convert strings to ints
            min, err := strconv.Atoi(rgx[1])
            checkError(err)
            max, err := strconv.Atoi(rgx[2])
            checkError(err)

            letterCount := strings.Count(rgx[4], rgx[3])

            // Check against part 1 rules
            if letterCount >= min && letterCount <= max {
                validPasswords1++
            }

            // Check against part 2 rules
            parsedPassword := []rune(rgx[4])
            if string(parsedPassword[min-1]) == rgx[3] && string(parsedPassword[max-1]) != rgx[3] {
                validPasswords2++
            } else if string(parsedPassword[min-1]) != rgx[3] && string(parsedPassword[max-1]) == rgx[3] {
                validPasswords2++
            }

        }
    }

    fmt.Println(fmt.Sprintf("[Part 1} Number of valid password: %d", validPasswords1))
    fmt.Println(fmt.Sprintf("[Part 2} Number of valid password: %d", validPasswords2))

}

// Dirty util function for checking errors
func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

func readFile(file string) ([]string, error) {
    fileBytes, err := ioutil.ReadFile(file)
    return strings.Split(string(fileBytes), "\n"), err
}
