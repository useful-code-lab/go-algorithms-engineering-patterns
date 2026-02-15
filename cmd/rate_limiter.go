package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	limit := rate.Every(100 * time.Millisecond)
	rateLimiter := rate.NewLimiter(limit, 2)

	ctx = context.Background()

	for i:=0; i<=10; i++ {

	}
}