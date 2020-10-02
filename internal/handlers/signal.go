package handlers

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func WatchSignals(g *errgroup.Group, ctx context.Context) {
	g.Go(func() error {
		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
		select {
		case sig := <-signalChannel:
			fmt.Printf("Received signal: %s\n", sig)
			ctx.Done()
		}
		return nil
	})

	if err := g.Wait(); err == nil || err == context.Canceled {
		fmt.Println("finished clean")
	} else {
		fmt.Printf("received error: %v", err)
	}
}
