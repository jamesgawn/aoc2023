package day1

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var textDigits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func ExecuteSolution(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		log.Println("Parsed text: " + text)

		numbers := FindNumbers(text)
		log.Println(numbers)

		number := CalculateNumber(numbers)
		log.Println("Calculated Number: " + strconv.Itoa(number))

		log.Println()
		total += number
	}
	return total
}

func FindNumbers(line string) (numbers []string) {
	numbers = make([]string, 0)

	for i := 0; i < len(line); i++ {
		digit := line[i : i+1]
		if unicode.IsDigit([]rune(digit)[0]) {
			numbers = append(numbers, digit)
		}
		for _, potentialDigit := range textDigits {
			if strings.HasPrefix(line[i:], potentialDigit) {
				numbers = append(numbers, potentialDigit)
				break
			}
		}
	}
	return numbers
}

func CalculateNumber(numbers []string) int {
	lastNumberPosition := len(numbers)
	firstNumber := NormaliseNumber(numbers[0])
	secondNumber := NormaliseNumber(numbers[lastNumberPosition-1])
	numberString := firstNumber + secondNumber
	number, err := strconv.Atoi(numberString)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func NormaliseNumber(number string) string {
	re := regexp.MustCompile(`[1-9]`)
	isNumeric := re.MatchString(number)
	if isNumeric {
		return string(number)
	}

	switch string(number) {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		log.Fatal("Number not found")
		return "0"
	}
}
