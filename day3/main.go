package day3

import (
	"bufio"
	"io"
	"regexp"
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

	for _, currentPart := range parts {
		LogPart(currentPart, matrix)
		isValid := currentPart.IsValidPart(matrix)
		if isValid {
			q1Answer += currentPart.number
		}
	}

	return q1Answer, q2Answer
}

func LogPart(currentPart part, matrix []string) {
	log.Debug().
		Int("partNo", currentPart.number).
		Int("matrixPosXStart", currentPart.matrixPosXStart).
		Int("matrixPosXEnd", currentPart.matrixPosXEnd).
		Int("matrixPosY", currentPart.matrixPosY).
		Int("searchRangeStart", currentPart.getRangeStart()).
		Int("searchRangeEnd", currentPart.getRangeEnd()).
		Bool("valid", currentPart.IsValidPart(matrix)).
		Send()
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

			if !isDigit || j == len(line)-1 {
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

func (currentPart part) IsValidPart(matrix []string) bool {
	return currentPart.IsValidPartDueToSameLine(matrix) || currentPart.IsValidPartDueToPreviousLine(matrix) || currentPart.IsValidPartDueToNextLine(matrix)
}

func (currentPart part) IsValidPartDueToPreviousLine(matrix []string) bool {
	if currentPart.matrixPosY == 0 {
		return false
	}
	previousLinePos := currentPart.matrixPosY - 1
	searchRange := currentPart.getSearchRange(matrix, previousLinePos)
	isValid, _ := regexp.Match(`.*[@/\=\-\+*&%$#]+.*`, []byte(searchRange))
	//log.Info().Int("partNo", currentPart.number).Str("searchRange", searchRange).Bool("isValid", isValid).Msg("checking previous line")
	return isValid
}

func (currentPart part) IsValidPartDueToNextLine(matrix []string) bool {
	if currentPart.matrixPosY == len(matrix)-1 {
		return false
	}
	nextLinePos := currentPart.matrixPosY + 1
	searchRange := currentPart.getSearchRange(matrix, nextLinePos)
	isValid, _ := regexp.Match(`.*[@/\=\-\+*&%$#]+.*`, []byte(searchRange))
	//log.Info().Int("partNo", currentPart.number).Str("searchRange", searchRange).Bool("isValid", isValid).Msg("checking next line")
	return isValid
}

func (currentPart part) IsValidPartDueToSameLine(matrix []string) bool {
	searchRange := currentPart.getSearchRange(matrix, currentPart.matrixPosY)
	isValid, _ := regexp.Match(`([@/\=\-\+*&%$#][0-9]+.*)|(.*[0-9]+[@\=\-\+*&%$#])`, []byte(searchRange))
	//log.Info().Int("partNo", currentPart.number).Str("searchRange", searchRange).Bool("isValid", isValid).Msg("checking middle line")
	return isValid
}

func (currentPart part) getRangeStart() int {
	start := currentPart.matrixPosXStart - 1
	if currentPart.matrixPosXStart == 0 {
		return 0
	}
	return start
}

func (currentPart part) getRangeEnd() int {
	return currentPart.matrixPosXEnd + 1
}

func (currentPart part) getSearchRange(matrix []string, line int) string {
	startRange := currentPart.getRangeStart()
	endRange := currentPart.getRangeEnd()
	return matrix[line][startRange:endRange]
}
