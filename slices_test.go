package main

import (
	"reflect"
	"testing"
)

func mutateSlice[P Primitives](arr []P) []P {
	switch arr := any(arr).(type) {
	case []string:
		arr[2] = "hello"
	case []int:
		arr[2] = 42
	case []bool:
		arr[2] = true
	}
	return arr
}

func TestCopySlice(t *testing.T) {
	t.Run("mutated slice should not equal original value", func(t *testing.T) {
		inp := []string{"a", "b", "c", "d"}
		mutateSlice(inp)
		got := inp
		want := []string{"a", "b", "c", "d"}
		if reflect.DeepEqual(got, want) {
			t.Errorf(`subject %v should not equal test slice %v after mutation`, got, want)
		}
	})
	t.Run("uncopied slice should be directly mutated by function", func(t *testing.T) {
		inp := []string{"a", "b", "c", "d"}
		mutateSlice(inp)
		got := inp
		want := []string{"a", "b", "hello", "d"}
		CheckGotWantPrimitive(t, got[2], want[2])
	})
	t.Run("copied slice of strings should not be mutated by function", func(t *testing.T) {
		inp := []string{"a", "b", "c", "d"}
		inpCopy := CopySlice(inp)
		mutateSlice(inpCopy)
		got := inp
		want := []string{"a", "b", "c", "d"}
		CheckGotWantSlice(t, got, want)
	})
	t.Run("copied slice of ints should not be mutated by function", func(t *testing.T) {
		inp := []int{0, 0, 0, 0, 0, 0}
		inpCopy := CopySlice(inp)
		mutateSlice(inpCopy)
		got := inp
		want := []int{0, 0, 0, 0, 0, 0}
		CheckGotWantSlice(t, got, want)
	})
}
