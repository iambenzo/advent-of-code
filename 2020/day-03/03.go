package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Add some arg parsing sauce
	if len(os.Args) <= 1 {
		fmt.Println("I need an input file")
		os.Exit(1)
	}

	slope, err := readFile(os.Args[1])
	checkError(err)

	part1 := slide(3, 1, slope)
	fmt.Printf("[Part 1] Tree count: %d\n", part1)

	fmt.Printf("[Part 2] Multiplied tree count: %d", slide(1, 1, slope)*slide(3, 1, slope)*slide(5, 1, slope)*slide(7, 1, slope)*slide(1, 2, slope))

}

func slide(moveX int, moveY int, slope [][]int32) int {
	// Define counts and key numbers
	mod := len(slope[0])
	slopeEnd := len(slope)
	treeCount := 0
	x := 0

	// Let's hit the slopes, dude
	for y := moveY; y < slopeEnd; y = y + moveY {
		// Make our move
		x = (x + moveX) % mod

		// Count the trees
		if string(slope[y][x]) == "#" {
			treeCount++
		}
	}

	return treeCount
}

// Dirty util function for checking errors
func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

// What we want to do here is read our file into
// a multi-dimension array so that we can use
// modular arithmetic to calculate the next position
func readFile(file string) ([][]rune, error) {
	fileBytes, err := ioutil.ReadFile(file)

	// Read each line and split the lines into character arrays
	// Well...Int32 arrays that can be parsed as strings...
	// Maybe theres a better way to do this??
	output := make([][]int32, len(strings.Split(string(fileBytes), "\n"))-1)
	for number, line := range strings.Split(string(fileBytes), "\n") {
		if len(line) > 0 {
			output[number] = []rune(line)
		}
	}
	return output, err
}
