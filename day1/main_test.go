package day1

import (
	"aoc2023/pkg/test"
	"aoc2023/pkg/utils"
	"testing"
)

func TestValidateWithExample(t *testing.T) {
	file := utils.LoadFile("./input_test.txt")
	result := ExecuteSolution(file)
	test.EqualsInt(t, 142, result)
}
