package gfshutdown

import (
	"os"
	"os/signal"
	"sync"
)

var nTo1Notify sync.WaitGroup

// ExecBeforeShutDown execute input function
func ExecBeforeShutDown(fn func()) {
	nTo1Notify.Add(1)

	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, os.Interrupt)

		select {
		// Execute before quit
		case <-sig:
			defer nTo1Notify.Done()
			fn()
		}
	}()
}

// Wait execute input function and wait all sub task to be executed
func Wait(fn func()) {
	ExecBeforeShutDown(fn)

	nTo1Notify.Wait()
}
