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
	cubes  []cubeSet
}

type cubeSet struct {
	redCubes   int
	blueCubes  int
	greenCubes int
}

func (game game) CalculateTotalCubes() cubeSet {
	var cubes cubeSet

	for _, currentSet := range game.cubes {
		cubes.redCubes += currentSet.redCubes
		cubes.greenCubes += currentSet.greenCubes
		cubes.blueCubes += currentSet.blueCubes
	}

	return cubes
}

func (game game) MinCubesRequired() cubeSet {
	minSet := cubeSet{redCubes: 0, blueCubes: 0, greenCubes: 0}

	for _, set := range game.cubes {
		if set.redCubes > minSet.redCubes {
			minSet.redCubes = set.redCubes
		}
		if set.blueCubes > minSet.blueCubes {
			minSet.blueCubes = set.blueCubes
		}
		if set.greenCubes > minSet.greenCubes {
			minSet.greenCubes = set.greenCubes
		}
	}

	return minSet
}

func (game game) IsValidGame(maxRedCubes int, maxBlueCubes int, maxGreenCubes int) bool {

	for _, set := range game.cubes {
		if set.redCubes > maxRedCubes || set.blueCubes > maxBlueCubes || set.greenCubes > maxGreenCubes {
			return false
		}
	}

	return true
}

func ExecuteSolution(input io.Reader) (q1Answer int, q2Answer int) {
	scanner := bufio.NewScanner(input)
	q1Answer = 0
	q2Answer = 0
	for scanner.Scan() {
		text := scanner.Text()
		log.Debug().Msg("Parsed text: " + text)

		game := ParseGame(text)
		LogGame(game)

		if game.IsValidGame(12, 14, 13) {
			q1Answer += game.gameId
		}

		q2Answer += game.MinCubesRequired().blueCubes * game.MinCubesRequired().redCubes * game.MinCubesRequired().greenCubes

	}
	return q1Answer, q2Answer
}

func ParseGame(text string) game {
	gameId := ParseGameNumber(text)

	gameNumberIndex := strings.Index(text, ":")
	pullsText := text[gameNumberIndex+1:]

	cubeSets := ExtractCubePulls(pullsText)

	return game{gameId: gameId, cubes: cubeSets}
}

func LogGame(game game) {
	log.Debug().Int("gameId", game.gameId).
		Int("redCubes", game.MinCubesRequired().blueCubes).
		Int("blueCubes", game.MinCubesRequired().blueCubes).
		Int("greenCubes", game.MinCubesRequired().greenCubes).
		Bool("valid", game.IsValidGame(12, 14, 13)).
		Send()
}

func ParseGameNumber(text string) int {
	gameNumberIndex := strings.Index(text, ":")
	gameNumberString := text[5:gameNumberIndex]
	gameNumber, _ := strconv.Atoi(gameNumberString)
	return gameNumber
}

func ExtractCubePulls(pullText string) []cubeSet {
	cubePulls := make([]cubeSet, 0)

	pulls := strings.Split(pullText, ";")

	for _, pull := range pulls {
		var set cubeSet

		pull := strings.Split(pull, ",")

		for _, cubeColour := range pull {
			trimmed := strings.Trim(cubeColour, " ")
			split := strings.Split(trimmed, " ")
			switch split[1] {
			case "red":
				set.redCubes, _ = strconv.Atoi(split[0])
			case "green":
				set.greenCubes, _ = strconv.Atoi(split[0])
			case "blue":
				set.blueCubes, _ = strconv.Atoi(split[0])
			}
		}

		cubePulls = append(cubePulls, set)
	}

	return cubePulls
}
