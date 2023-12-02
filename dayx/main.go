package dayx

import (
	"bufio"
	"io"

	"github.com/rs/zerolog/log"
)

func ExecuteSolution(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		log.Debug().Msg("Parsed text: " + text)
	}
	return total
}
