package config

import (
	"log"
	"os"
)

type Logger struct {
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
}

var applogger *Logger

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	InfoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	applogger = &Logger{
		Warning: WarningLogger,
		Info:    InfoLogger,
		Error:   ErrorLogger,
	}

}

func GetLogger() *Logger {
	return applogger
}
