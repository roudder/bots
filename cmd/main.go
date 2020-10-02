package main

import (
	"bots/internal/cache"
	. "bots/internal/handlers"
	. "bots/internal/routes"
	"context"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	cache.New()
	StartServer()
	WatchSignals(g, ctx)
}
