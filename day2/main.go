package day2

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type game struct {
	gameId int
	// redCubes   int
	// blueCubes  int
	// greenCubes int
}

func ExecuteSolution(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	total := 0
	for scanner.Scan() {
		text := scanner.Text()
		log.Debug().Msg("Parsed text: " + text)

		game := ParseGame(text)
		log.Debug().Interface("dict", game)
	}
	return total
}

func ParseGame(text string) game {
	gameId := ParseGameNumber(text)
	return game{gameId: gameId}
}

func ParseGameNumber(text string) int {
	gameNumberIndex := strings.Index(text, ":")
	gameNumberString := text[5:gameNumberIndex]
	gameNumber, _ := strconv.Atoi(gameNumberString)
	return gameNumber
}
