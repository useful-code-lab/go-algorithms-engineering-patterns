package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens     float64
	capacity   float64 //На случай если пришло одновременно больше токенов
	lastUpdate time.Time
	weight     float64

	mu sync.Mutex
}

func Allow(tokenBucket *TokenBucket) bool {
	tokenBucket.mu.Lock()
	defer tokenBucket.mu.Unlock()

	now := time.Now()
	diff := time.Since(tokenBucket.lastUpdate).Seconds()
	tokenBucket.tokens += diff * tokenBucket.weight

	if tokenBucket.tokens > tokenBucket.capacity {
		tokenBucket.tokens = tokenBucket.capacity
	}

	tokenBucket.lastUpdate = now

	if tokenBucket.tokens >= 1 {
		tokenBucket.tokens--
		return true
	} 

	return false
}

func main() {
	tokenBucket := TokenBucket{tokens: 10, capacity: 10, weight: 1, lastUpdate: time.Now()}

	for i := range 100 {
		if Allow(&tokenBucket) {
			fmt.Printf("Запрос %v разрешается делать", i+1)
		} else {
			fmt.Printf("Запрос %v не разрешается делать", i+1)
		}
	}
}
