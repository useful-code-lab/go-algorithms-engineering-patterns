package main

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	tokens float64
	capacity float64
	lastUpdate time.Time

	mu sync.Mutex
}


func Allow(tokenBucket TokenBucket) bool {

	defer tokenBucket.mu.Unlock()
	return false
}

func main() {
	tokenBucket := TokenBucket{tokens: 10, capacity: 10}

	for i := range 100 {
		if Allow(tokenBucket) {
			fmt.Printf("Запрос %v разрешается делать", i + 1)
		} else {
			fmt.Printf("Запрос %v не разрешается делать", i + 1)
		}
	}
}