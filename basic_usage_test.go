package main

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {

	log := logrus.New()

	log.Println("Hello World")

	fmt.Println("Hello World")
}
