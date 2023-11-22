package main

import "fmt"

// Go passes by values for int and string
// but arrays and maps are mutated unless you copy
func CopyArray[T any](arr []T) []T {
	arrCopy := make([]T, len(arr))
	copy(arrCopy, arr)
	return arrCopy
}

func copyArrayExample() {
	original := []int{1, 2, 3, 4, 5}
	copy := CopyArray(original)
	original[0] = 45
	fmt.Println("original: ", original)
	fmt.Println("copy: ", copy)
}
