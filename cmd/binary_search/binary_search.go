package main

import "fmt"

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func main() {
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range nums {
		fmt.Printf("%d-й элемент массива: %d\n", i+1, v)
	}

	if pos := binarySearch(nums[:], 4); pos != -1 {
		fmt.Printf("Позиция в массива элемента %d: %d", 4, pos)
	} else {
		fmt.Printf("Элемент в массиве не найден")
	}

}
