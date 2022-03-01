package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// Create function TestLoggingFormatterJson example
func TestLoggingFormatterJson(t *testing.T) {

	// Create a new instance of the FormatterJson type
	log := logrus.New()

	// Set the formatter to be FormatterJson
	log.SetFormatter(&logrus.JSONFormatter{})

	// Output: {"level":"info","msg":"Hello World"}
	log.Info("Hello World")
}

// Create function TestLoggingFormatterText example
func TestLoggingFormatterText(t *testing.T) {

	// Create a new instance of the FormatterText type
	log := logrus.New()

	// Set the formatter to be FormatterText
	log.SetFormatter(&logrus.TextFormatter{})

	// Output: [INFO] Hello World
	log.Info("Hello World")
}
