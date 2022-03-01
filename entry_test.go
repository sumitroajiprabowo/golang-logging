package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// Create function TestLoggingEntry example with json formatter
func TestLoggingEntry(t *testing.T) {

	// Create a new instance of the Entry type
	log := logrus.NewEntry(logrus.New())

	// Output: {"level":"info","msg":"Hello World"}
	log.WithFields(logrus.Fields{
		"email":    "bowo@anakdesa.id",
		"username": "bowo_anakdesa",
	}).Info("Hello World")

	// Output: {"level":"info","msg":"Hello World","full_name":"Sumitro Aji Prabowo"}
	log.WithField("full_name", "Sumitro Aji Prabowo").Info("Hello World")

}
