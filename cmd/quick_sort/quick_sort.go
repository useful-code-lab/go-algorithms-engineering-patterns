package main

import "fmt"

func QuickSort(arr []int) {

}

func QuickSortRecursive(arr []int, low, high int){
	if low >= high {
		return
	}

	pivot := Partition(arr, low, high)
	QuickSortRecursive(arr, low, pivot - 1)
	QuickSortRecursive(arr, pivot + 1, high)
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high - 1; j++ {
		if 	arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] 
		}
	}

	arr[high], arr[i + 1] = arr[i + 1], arr[high]

	return i + 1
}


func main(){
	arr := []int{3, 4, 7, 3, 7, 10, 9, 8}
	fmt.Println(arr)

	//QuickSort(arr)
}
