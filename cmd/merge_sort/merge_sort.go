package main

import "fmt"

type MyNumbers interface {
	~int | ~int64 | ~float64 | ~string
}

func MergeSort[T MyNumbers](a []T) []T {
	if len(a) < 1 {
		return a
	}

	d := len(a) / 2
	m1 := MergeSort(a[:d])
	m2 := MergeSort(a[d:])

	return merge(m1, m2)
}

func merge[T MyNumbers](a []T, b []T) []T {
	result := make([]T, len(a)+len(b))

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			result = append(result, b[j])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}

func main() {
	arr := []int{4, 6, 7, 8, 9, 6}
	fmt.Print("Начальный массив:\n", arr)
	MergeSort(arr)

	fmt.Print("Отсортированный массив:\n", arr)


}
