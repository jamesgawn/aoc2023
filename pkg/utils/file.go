package utils

import (
	"os"

	"github.com/rs/zerolog/log"
)

func LoadFile(inputFileLocation string) *os.File {
	log.Debug().Msg("Using input file: " + inputFileLocation)
	file, err := os.Open(inputFileLocation)
	if err != nil {
		log.Info().Msg("I can't solve every problem... not yet anyway")
		log.Fatal().Err(err).Send()
	}

	return file
}

func CloseFile(file *os.File) {
	cErr := file.Close()
	if cErr != nil {
		log.Fatal().Err(cErr).Send()
	}
}
