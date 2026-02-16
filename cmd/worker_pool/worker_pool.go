package main

func Worker(numWorker int, inp <-chan int, result chan<- int) {

}


func main() {
	digits := [10]int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	var result [10]int
	chInput := make(chan int)
	chOutput := make(chan int)
	for i := range 10 {
		go Worker(i, chInput, chOutput)
	}
}