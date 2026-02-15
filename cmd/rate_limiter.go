package main

import (
	"context"
	"fmt"

	"golang.org/x/time/rate"
)

func main() {

	limit := rate.Limit(1000)
	rateLimiter := rate.NewLimiter(limit, 2)
	ctx := context.Background()

	for i := 0; i <= 10; i++ {
		err := rateLimiter.Wait(ctx)

		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			return
		}

		fmt.Printf("Печать %v\n", i)
	}
}
