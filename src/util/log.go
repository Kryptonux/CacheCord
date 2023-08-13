package util

import (
	"fmt"
	"os"
)

var (
	verboseEnabled = false
)

func EnableVerboseLogging() {
	verboseEnabled = true
}

func LogSuccess(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[SUCCESS] %s\n", message)
}

func LogError(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "[ERROR] %s\n", message)
}

func LogVerbose(format string, args ...interface{}) {
	if verboseEnabled {
		message := fmt.Sprintf(format, args...)
		fmt.Printf("[VERBOSE] %s\n", message)
	}
}
