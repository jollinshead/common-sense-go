package iferror

import "log"

// Panic invokes panic given a non-nil error
//
//   iferror.Panic(err)
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

// LogPrint invokes log.Print given a non-nil error
//
//   iferror.LogPrint(err, "My message")
func LogPrint(err error, v ...interface{}) {
	logFunc(err, log.Print, v)
}

// LogFatal invokes log.Fatal given a non-nil error
//
//   iferror.LogFatal(err, "My message")
func LogFatal(err error, v ...interface{}) {
	logFunc(err, log.Fatal, v)
}

var logFunc = func(err error, f func(...interface{}), v ...interface{}) {
	if err != nil {
		f(v)
	}
}

