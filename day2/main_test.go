package day2

import (
	"aoc2023/pkg/test"
	"aoc2023/pkg/utils"
	"testing"
)

func TestValidateWithExample1(t *testing.T) {
	file := utils.LoadFile("./input_test.txt")
	result := ExecuteSolution(file)
	test.EqualsInt(t, 8, result)
}

// func TestValidateWithExample2(t *testing.T) {
// 	file := utils.LoadFile("./input_test2.txt")
// 	result := ExecuteSolution(file)
// 	test.EqualsInt(t, 281, result)
// }

func TestWorksWithRealFile(t *testing.T) {
	file := utils.LoadFile("./input.txt")
	ExecuteSolution(file)
}
