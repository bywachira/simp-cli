package helpers

import (
	"log"
	"os"
)

// Logger holds all methods to handle logging
type Logger struct{}

// CreateLoggerFileAndPopulate handles creating a logger file
func (l *Logger) CreateLoggerFileAndPopulate(logData string) {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("simping.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println(logData)
}
