package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		for {
			if time.Now().Second()%15 == 0 {
				cancel()
			}
		}
	}()

	go P1(ctx, 0)
	P2(ctx, 0)

	defer cancel()
}
func P1(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			println("Task Cancelled")
			return
		default:
			num++
			time.Sleep(time.Second * 1)
			println("Processing Add ", num)
		}
	}
}

func P2(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			println("Task Cancelled")
			return
		default:
			num--
			time.Sleep(time.Second * 1)
			println("Processing  Sub ", num)
		}
	}
}
