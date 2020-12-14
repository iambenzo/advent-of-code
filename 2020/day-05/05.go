package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("I need an input file")
		os.Exit(1)
	}
	passes, err := readFile(os.Args[1])
	checkError(err)

	highestBoardingPass := 0
	lowestBoardingPass := 127*8 + 7
	bumsOnSeats := make(map[int]bool)

	for _, pass := range passes {
        rowMin := 0
        rowMax := 127
        seatMin := 0
        seatMax := 7
        move := 2.0
		for num, step := range []rune(pass) {
			if num == 7 {
				move = 2.0
			}
			switch string(step) {
			case "F":
				rowMax = int(math.Floor(float64(rowMax) / 2))
			case "B":
				rowMin = rowMin + int(math.Floor(float64(128)/move))
			case "L":
				seatMax = int(math.Floor(float64(seatMax) / 2))
			case "R":
				seatMin = seatMin + int(math.Floor(float64(8)/move))
			}
			move += move
		}
		seatId := rowMin*8 + seatMin

		if seatId > highestBoardingPass {
			highestBoardingPass = seatId
		}

		if seatId < lowestBoardingPass {
			lowestBoardingPass = seatId
		}

		bumsOnSeats[seatId] = true
	}

	fmt.Printf("[Part 1] The highest seat ID is: %d\n", highestBoardingPass)

	for i := lowestBoardingPass; i < highestBoardingPass; i++ {
		if !bumsOnSeats[i] && (bumsOnSeats[i-1] && bumsOnSeats[i+1]) {
			fmt.Printf("[Part 2] My seat is ID: %d\n", i)
			break
		}
	}
}

func readFile(file string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(file)
	return strings.Split(string(fileBytes), "\n"), err
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}
