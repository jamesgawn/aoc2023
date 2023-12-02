package main

import (
	"aoc2023/day1"
	"aoc2023/dayx"
	"aoc2023/pkg/utils"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Total execution")
	day := "4"
	if len(os.Args) > 1 {
		day = string(os.Args[1])
	}
	inputFileLocation := "./day" + day + "/input.txt"
	if len(os.Args) > 2 {
		inputFileLocation = string(os.Args[1])
	}

	file := utils.LoadFile(inputFileLocation)

	FindSolution(day, file)

	utils.CloseFile(file)
}

func FindSolution(day string, input io.Reader) {
	log.Println("Day " + day)
	var result int

	switch day {
	case "x":
		dayx.ExecuteSolution(input)
	case "1":
		result = day1.ExecuteSolution(input)
	default:
		log.Fatal("I can't solve every problem... not yet anyway")
	}

	log.Println("The answer for day " + day + " is: " + strconv.Itoa(result))
}
