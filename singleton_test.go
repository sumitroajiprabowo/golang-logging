package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// Create function TestSingleton
func TestSingleton(t *testing.T) {

	logrus.Info("TestSingleton")
	logrus.Warn("TestSingleton")
	logrus.Error("TestSingleton")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("TestSingleton")
	logrus.Warn("TestSingleton")
	logrus.Error("TestSingleton")

}
