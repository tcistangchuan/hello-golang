package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// 注意： withValue 的ctx 是 parent context
	newCtx := context.WithValue(ctx, "name", "tangchuan")
	newCtx2 := context.WithValue(newCtx, "old", "18")
	defer cancel()
	log.Println("start")
	time.Sleep(2 * time.Second)
	go func() {
		select {
		case <-newCtx2.Done():
			log.Println("end", newCtx2.Value("name"))
			return
		}
	}()
	select {
	case <-newCtx2.Done():
		log.Println("end2", newCtx2.Value("old"))
		return
	}
}
