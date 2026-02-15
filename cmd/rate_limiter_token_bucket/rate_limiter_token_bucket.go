package main

import "fmt"

func Allow() bool {
	return false
}

func main() {

	for i := 0; i < 100; i++ {
		if Allow() {
			fmt.Printf("Запрос %v разрешается делать", i + 1)
		} else {
			fmt.Printf("Запрос %v не разрешается делать", i + 1)
		}
	}
}