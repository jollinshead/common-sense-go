package iferror

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
	"bytes"
	"log"
	"fmt"
)

func TestPanic(t *testing.T) {
	err := errors.New("test error")

	assert.Panics(t, func() { Panic(err)} )
}

func TestPanic_noError(t *testing.T) {
	var err error = nil

	assert.NotPanics(t, func() { Panic(err)} )
}

func TestLogPrint(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	err := errors.New("test error")
	message := fmt.Sprintf("We have an error! %s", err)

	LogPrint(err, message)

	assert.Contains(t, buf.String(), message)
}

func TestLogPrint_noError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)

	var err error = nil
	message := "Should not be used"

	LogPrint(err, log.Print, message)

	assert.Equal(t, 0, buf.Len())
}

func TestLogFatal(t *testing.T) {
	invocationCount := 0

	message := "We have a fatal error!"
	cleanUp := mockLogFunc(func(err error, f func(...interface{}), v ...interface{}) {
		invocationCount++
		expectedAddress := fmt.Sprintf("%v", log.Fatal)
		actualAddress := fmt.Sprintf("%v", f)
		assert.Equal(t, expectedAddress, actualAddress)

		assert.Len(t, v, 1)
		assert.Contains(t, v[0], message)
	})
	defer cleanUp()

	err := errors.New("test error")

	LogFatal(err, message)

	assert.Equal(t, 1, invocationCount)
}

func TestLogFatal_noError(t *testing.T) {
	invocationCount := 0

	message := "We have a fatal error!"
	cleanUp := mockLogFunc(func(err error, f func(...interface{}), v ...interface{}) {
		invocationCount++
		assert.NoError(t, err)
	})
	defer cleanUp()

	var err error = nil

	LogFatal(err, message)

	assert.Equal(t, 1, invocationCount)
}

// -- Mocks and helpers

func mockLogFunc(mockFunc func(err error, f func(...interface{}), v ...interface{})) (cleanUp func()) {
	originalLogFunc := logFunc
	logFunc = mockFunc
	return func() {logFunc = originalLogFunc}
}

