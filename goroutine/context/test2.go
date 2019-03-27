package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出了...")
				return
			default:
				fmt.Println("监控中...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	time.Sleep(5 * time.Second)

	cancel()

	fmt.Println("监控结束...")
}
