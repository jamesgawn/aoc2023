package day2

import (
	"aoc2023/pkg/test"
	"aoc2023/pkg/utils"
	"testing"
)

func TestValidateWithExample1(t *testing.T) {
	file := utils.LoadFile("./input_test.txt")
	q1Answer, _ := ExecuteSolution(file)
	test.EqualsInt(t, 8, q1Answer)
}

func TestValidateWithExample2(t *testing.T) {
	file := utils.LoadFile("./input_test.txt")
	_, q2Answer := ExecuteSolution(file)
	test.EqualsInt(t, 2286, q2Answer)
}

func TestWorksWithRealFile(t *testing.T) {
	file := utils.LoadFile("./input.txt")
	ExecuteSolution(file)
}

func TestIsValidCheckWhenValid(t *testing.T) {
	set := cubeSet{redCubes: 12, blueCubes: 14, greenCubes: 13}
	game := game{gameId: 1, cubes: []cubeSet{set}}
	test.EqualsBool(t, true, game.IsValidGame(12, 14, 13))
}

func TestIsValidCheckWhenInvalid(t *testing.T) {
	set := cubeSet{redCubes: 13, blueCubes: 15, greenCubes: 14}
	game := game{gameId: 1, cubes: []cubeSet{set}}
	test.EqualsBool(t, false, game.IsValidGame(12, 14, 13))
}
