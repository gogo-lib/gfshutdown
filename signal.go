package gfshutdown

import (
	"os"
	"os/signal"
	"sync"
)

var nTo1Notify sync.WaitGroup

// ExecuteBeforeQuit execute input function
// WARN: this function is blocking operation, should use with goroutine
func ExecuteBeforeQuit(fn func()) {
	nTo1Notify.Add(1)

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)

	select {
	// Execute before quit
	case <-sig:
		defer nTo1Notify.Done()
		fn()
	}
}

// Wait execute input function and wait all sub task to be executed
func Wait(fn func()) {
	ExecuteBeforeQuit(fn)

	nTo1Notify.Wait()
}
