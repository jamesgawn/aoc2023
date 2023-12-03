package utils

import "github.com/rs/zerolog/log"

func PrintAnswer(day int, question int, answer int) {
	log.Info().Int("day", day).Int("question", question).Int("answer", answer).Send()
}
