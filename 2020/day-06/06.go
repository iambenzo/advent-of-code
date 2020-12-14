package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
    if len(os.Args) <= 1 {
        fmt.Println("I need an input file")
        os.Exit(1)
    }

    // This time I'm going to process the file as I read it
    fileBytes, err := ioutil.ReadFile(os.Args[1])
    checkError(err)

    anyoneSum := 0
    everyoneSum := 0

    for _, group := range strings.Split(string(fileBytes), "\n\n") {
        // For each group, get an array of answers and run through them
        answerMap := make(map[string]int)
        headCount := len(strings.Split(group, "\n"))
        for _, answer := range []rune(strings.ReplaceAll(group, "\n", "")) {
            //For each answer, record the letter in our "set"
            answerMap[string(answer)]++
            if answerMap[string(answer)] == headCount {
                everyoneSum++
            }
        }
        anyoneSum += len(answerMap)
    }

    fmt.Printf("[Part 1] The sum of all groups answers: %d\n", anyoneSum)
    fmt.Printf("[Part 2] The sum of all groups answers: %d\n", everyoneSum)
}

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}
