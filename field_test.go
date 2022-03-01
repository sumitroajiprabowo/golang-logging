package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// Create function TestLoggingFields example with json formatter
func TestLoggingFields(t *testing.T) {

	// Create a new instance of the FormatterJson type
	log := logrus.New()

	// Set the formatter to be FormatterJson
	log.SetFormatter(&logrus.JSONFormatter{})

	// Output: {"level":"info","msg":"Hello World"}
	log.WithFields(logrus.Fields{
		"email":    "bowo@anakdesa.id",
		"username": "bowo_anakdesa",
	}).Info("Hello World")

	// Output: {"level":"info","msg":"Hello World","full_name":"Sumitro Aji Prabowo"}
	log.WithField("full_name", "Sumitro Aji Prabowo").Info("Hello World")
}
