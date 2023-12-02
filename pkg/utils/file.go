package utils

import (
	"log"
	"os"
)

func LoadFile(inputFileLocation string) *os.File {
	log.Println("Using input file: " + inputFileLocation)
	file, err := os.Open(inputFileLocation)
	if err != nil {
		log.Println("I can't solve every problem... not yet anyway")
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
