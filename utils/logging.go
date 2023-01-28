package utils

import (
	"log"
	"os"
)

var (
	// Log is used for logging throughout the application
	Log *log.Logger
)

func init() {
	Log = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
}

// Info logs a message with severity info
func LogInfo(v ...interface{}) {
	Log.SetPrefix("[INFO] ")
	Log.Println(v...)
}

// Warning logs a message with severity warning
func LogWarning(v ...interface{}) {
	Log.SetPrefix("[WARNING] ")
	Log.Println(v...)
}

// Error logs a message with severity error
func LogError(v ...interface{}) {
	Log.SetPrefix("[ERROR] ")
	Log.Println(v...)
}
