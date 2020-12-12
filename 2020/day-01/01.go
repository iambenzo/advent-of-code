package main

import (
	"bufio"
    "strconv"
	"fmt"
	"os"
)

func main() {
    expenses, err := readFile()
    checkError(err)

    for k, _ := range expenses {
        // With the first number in hand, work out what the matching value would be
        match := 2020 - k

        // If our "set" contains the matching value
        // Print and exit
        if expenses[match] {
            fmt.Println(fmt.Sprintf("Found that %d and %d add to 2020", k, match))
            fmt.Println(fmt.Sprintf("Answer: %d", k*match))
            os.Exit(0)
        }
        // Otherwise, continue...
    }

}

// Dirty util function for checking errors
func checkError(e error) {
    if e != nil {
        panic(e)
    }
}

// Read file contents into a map that's been bastardised
// into some form of set
func readFile() (map[int]bool, error) {
    out := make(map[int]bool)

    // Read the file into memory
    file, err := os.Open("./input.txt")
    checkError(err)
    // defer file closure until function completion
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        // String to number conversion
        num, err := strconv.Atoi(scanner.Text())
        checkError(err)
        out[num] = true
    }

    return out, scanner.Err()
}
