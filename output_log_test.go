package main

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLoggingFileOutput(t *testing.T) {

	log := logrus.New()

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Debug("Hello World")
	log.Info("Hello World")
	log.Warn("Hello World")
	log.Error("Hello World")
	// log.Fatal("Hello World")

}
