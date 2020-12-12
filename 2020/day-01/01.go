package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    expenses, err := readFile()
    checkError(err)

    var part1Found bool
    var part2Found bool

    for k := range expenses {
        // With the first number in hand, work out what the matching value would be
        firstMatch := 2020 - k

        // If our "set" contains the matching value
        // Print and exit
        if !part1Found && expenses[firstMatch] {
            part1Found = true
            fmt.Println(fmt.Sprintf("Found that %d and %d add to 2020", k, firstMatch))
            fmt.Println(fmt.Sprintf("Answer for part one: %d", k*firstMatch))
        }

        // Wanted to avoid inner loops but hey ho
        for k2 := range expenses {
            secondMatch := firstMatch - k2
            if !part2Found && expenses[secondMatch] {
                part2Found = true
                fmt.Println(fmt.Sprintf("Found that %d, %d and %d add to 2020", k, secondMatch, k2))
                fmt.Println(fmt.Sprintf("Answer for part two: %d", k*secondMatch*k2))
            }
        }

        // Otherwise, continue...
        if part1Found && part2Found {
            os.Exit(0)
        }
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
