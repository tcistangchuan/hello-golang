package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	log.Println("start")
	go handle(ctx, 4*time.Second)
	select {
	case <-ctx.Done():
		log.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("handle", ctx.Err())
	case <-time.After(duration):
		log.Println("process request with", duration)
	}

}
