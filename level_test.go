package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevelLoggerTraceLevel(t *testing.T) {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)
	log.Trace("Hello World")
}

func TestLevelLoggerDebugLevel(t *testing.T) {

	log := logrus.New()

	log.SetLevel(logrus.DebugLevel) // Tracelevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel

	log.Debug("Hello World")

	log.Info("Hello World")

	log.Warn("Hello World")

	log.Error("Hello World")

	// log.Fatal("Hello World")

	// log.Panic("Hello World")
}

func TestLevelLoggerInfoLevel(t *testing.T) {

	log := logrus.New()

	log.SetLevel(logrus.InfoLevel) // Tracelevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel

	log.Info("Hello World")

	log.Warn("Hello World")

	log.Error("Hello World")

	// log.Fatal("Hello World")

	// log.Panic("Hello World")
}

func TestLevelLoggerWarnLevel(t *testing.T) {

	log := logrus.New()

	log.SetLevel(logrus.WarnLevel) // Tracelevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel

	log.Warn("Hello World")

	log.Error("Hello World")

	// log.Fatal("Hello World")

	// log.Panic("Hello World")

}

func TestLevelLoggerErrorLevel(t *testing.T) {

	log := logrus.New()

	log.SetLevel(logrus.ErrorLevel) // Tracelevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel

	log.Error("Hello World")

	// log.Fatal("Hello World")

	// log.Panic("Hello World")

}

// func TestLevelLoggerFatalLevel(t *testing.T) {

// 	log := logrus.New()

// 	log.SetLevel(logrus.FatalLevel) // Tracelevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel

// 	log.Fatal("Hello World")

// 	// log.Panic("Hello World")

// }

// func TestLevelLoggerPanicLevel(t *testing.T) {

// 	log := logrus.New()

// 	log.SetLevel(logrus.PanicLevel) // Tracelevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel, FatalLevel, PanicLevel

// 	log.Panic("Hello World")

// }
