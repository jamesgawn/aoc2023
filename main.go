package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/pkg/utils"
	"flag"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	debug := flag.Bool("debug", false, "sets the log level to debug")
	day := flag.String("day", "1", "specify the day to run for")
	inputFile := flag.String("inputFile", "", "specify an input file for testing")

	flag.Parse()

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	defer utils.TimeTrack(time.Now(), "Total execution")

	inputFileLocation := "./day" + *day + "/input.txt"
	if *inputFile != "" {
		inputFileLocation = *inputFile
	}

	file := utils.LoadFile(inputFileLocation)

	FindSolution(*day, file)

	utils.CloseFile(file)
}

func FindSolution(day string, input io.Reader) {
	switch day {
	case "1":
		day1.ExecuteSolution(input)
	case "2":
		day2.ExecuteSolution(input)
	default:
		log.Fatal().Msg("I can't solve every problem... not yet anyway")
	}
}
