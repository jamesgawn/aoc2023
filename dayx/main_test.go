package dayx

import (
	"aoc2023/pkg/test"
	"aoc2023/pkg/utils"
	"testing"
)

func TestValidateWithExample1(t *testing.T) {
	file := utils.LoadFile("./input_test.txt")
	q1Answer, _ := ExecuteSolution(file)
	test.EqualsInt(t, 0, q1Answer)
}

func TestValidateWithExample2(t *testing.T) {
	file := utils.LoadFile("./input_test.txt")
	_, q2Answer := ExecuteSolution(file)
	test.EqualsInt(t, 0, q2Answer)
}

func TestWorksWithRealFile(t *testing.T) {
	file := utils.LoadFile("./input.txt")
	ExecuteSolution(file)
}
