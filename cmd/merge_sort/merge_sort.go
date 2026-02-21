package main

type MyNumbers interface {
	~int | ~int64 | ~float64 | ~string
}

func MergeSort[T MyNumbers](a []T) []T {
	
	return merge()  
}

func merge[T MyNumbers](a []T, b []T) []T{
	result := make([]T, len(a) + len(b))
	return result
}

func main() {
	arr := []int{4, 6, 7, 8, 9, 6}

	MergeSort(arr)

}
