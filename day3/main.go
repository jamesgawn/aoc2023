package day3

import (
	"bufio"
	"io"
	"strconv"
	"unicode"

	"github.com/rs/zerolog/log"
)

type part struct {
	number          int
	numberString    string
	matrixPosXStart int
	matrixPosXEnd   int
	matrixPosY      int
}

func ExecuteSolution(input io.Reader) (q1Answer int, q2Answer int) {
	scanner := bufio.NewScanner(input)
	q1Answer = 0
	q2Answer = 0

	matrix := make([]string, 0)
	lineNumber := 0

	for scanner.Scan() {
		text := scanner.Text()
		log.Debug().Msg("Parsed text: " + text)

		matrix = append(matrix, text)

		lineNumber++
	}

	parts := FindPartNumbers(matrix)

	log.Debug().Interface("parts", parts).Msg("Parsed Parts")

	return q1Answer, q2Answer
}

func FindPartNumbers(matrix []string) []part {
	parts := make([]part, 0)

	for i, line := range matrix {
		potentialPartNo := ""
		startOfPartNo := 0

		for j, character := range line {
			isDigit := unicode.IsDigit(character)

			if isDigit {
				if len(potentialPartNo) == 0 {
					startOfPartNo = j
				}
				potentialPartNo = potentialPartNo + string(character)
			}

			if !isDigit {
				if len(potentialPartNo) != 0 {
					partNo, _ := strconv.Atoi(potentialPartNo)
					currentPart := part{numberString: potentialPartNo, number: partNo, matrixPosY: i, matrixPosXStart: startOfPartNo, matrixPosXEnd: j}

					parts = append(parts, currentPart)

					startOfPartNo = 0
					potentialPartNo = ""
				}
			}
		}
	}

	return parts
}
