package main

import (
	"context"
	"fmt"
	"time"
)

func infiniteLoop(ctx context.Context) {
	// 終わらないやつ
	for {
		fmt.Println("Help!")
	}
}

func main() {
	ctx := context.Background()

	// 2秒待ってキャンセルする
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go infiniteLoop(ctx)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
