package main

import (
	"context"
	"os"
	"time"
)

func infiniteLoop(ctx context.Context) {
	for {
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	go infiniteLoop(ctx)
	select {
	case <-ctx.Done():
		os.Exit(0)
	}
}
