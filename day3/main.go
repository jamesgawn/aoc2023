package day3

import (
	"bufio"
	"io"

	"github.com/rs/zerolog/log"
)

func ExecuteSolution(input io.Reader) (q1Answer int, q2Answer int) {
	scanner := bufio.NewScanner(input)
	q1Answer = 0
	q2Answer = 0
	for scanner.Scan() {
		text := scanner.Text()
		log.Debug().Msg("Parsed text: " + text)

	}
	return q1Answer, q2Answer
}
