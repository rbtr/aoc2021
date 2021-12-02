package common

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var rootCtx context.Context

// init() is executed before main() whenever this package is imported
// to do pre-run setup of things like exit signal handling and building
// the root context.
func init() {
	var cancel context.CancelFunc
	rootCtx, cancel = context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		// Wait until receiving a signal.
		sig := <-sigCh
		log.Printf("caught exit signal %v, exiting\n", sig)
		cancel()
		<-time.Tick(5 * time.Second) //nolint // wait up to 5s for the cancellation to propogate before exiting hard.
		log.Printf("exiting")
		os.Exit(1)
	}()
}
