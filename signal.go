package spark

import (
	"os"
	"os/signal"
	"syscall"
)

func signalHandler() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1)
	for sig := range c {
		if sig == syscall.SIGUSR1 {
			relaunch = true
		}

		close(quit)
		break
	}
}
