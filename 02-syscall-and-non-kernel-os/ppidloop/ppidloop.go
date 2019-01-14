package main

import (
	"context"
	"os"
	"syscall"
	"time"
)

func infiniteLoop(ctx context.Context) {
	for {
		syscall.Getppid()
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
