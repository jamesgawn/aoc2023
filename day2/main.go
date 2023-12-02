package day2

import (
	"bufio"
	"io"

	log "github.com/sirupsen/logrus"
)

func ExecuteSolution(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		log.Info("Parsed text: " + text)
	}
	return total
}
