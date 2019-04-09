package context

import (
	"context"
	"fmt"
	"log"
	"time"
)

// StartUp the code here
func StartUp() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	cancel()
	mySleepAndTalk(ctx, 5*time.Second, "Hello")
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
