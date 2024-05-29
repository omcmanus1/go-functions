package main

func MutateSlice[T Primitives](slc []T) []T {
	switch slc := any(slc).(type) {
	case []string:
		slc[2] = "hello"
	case []int:
		slc[2] = 42
	case []bool:
		slc[2] = true
	}
	return slc
}

func CopySlice[T Primitives](slc []T) []T {
	slcCopy := make([]T, len(slc))
	return slcCopy
}
