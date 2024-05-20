package main

// Go passes by values for int and string
// but arrays, slices and maps are mutated unless you copy
func CopySlice[T any](arr []T) []T {
	arrCopy := make([]T, len(arr))
	copy(arrCopy, arr)
	return arrCopy
}
