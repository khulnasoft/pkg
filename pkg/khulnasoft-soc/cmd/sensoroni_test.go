package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLogging(tester *testing.T) {
	testFile := "/tmp/sensoroni_test.log"
	defer os.Remove(testFile)
	file, err := InitLogging(testFile, "debug")
	if assert.Nil(tester, err) {
		assert.NotNil(tester, file)
	}
}
