package main

import "fmt"

func Worker(numWorker int, inp <-chan int, result chan<- int) {
	for i := range inp {
		result <- i * i
	}
}

func main() {
	digits := [10]int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	chInput := make(chan int)
	chOutput := make(chan int)

	for i := range 3 {
		go Worker(i, chInput, chOutput)
	}

	go func() {
		for i := range 10 {
			chInput <- digits[i]
		}
		close(chInput)
	}()

	for range 10 {
		r := <-chOutput
		fmt.Printf("Вывод результата: %d\n", r)
	}
}
