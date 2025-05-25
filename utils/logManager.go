package utils

import (
	"fmt"
	"os"
	"sync"
)

var printMu sync.Mutex

func SafePrint(format string, args ...interface{}) {
	printMu.Lock()
	defer printMu.Unlock()
	fmt.Printf(format, args...)
}

func SafeError(format string, args ...interface{}) {
	printMu.Lock()
	defer printMu.Unlock()
	fmt.Fprintf(os.Stderr, format, args...)
}
