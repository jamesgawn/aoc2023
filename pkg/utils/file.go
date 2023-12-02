package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LoadFile(inputFileLocation string) *os.File {
	log.Info("Using input file: " + inputFileLocation)
	file, err := os.Open(inputFileLocation)
	if err != nil {
		log.Info("I can't solve every problem... not yet anyway")
		log.Fatal(err)
	}

	return file
}

func CloseFile(file *os.File) {
	cErr := file.Close()
	if cErr != nil {
		log.Fatal(cErr)
	}
}
