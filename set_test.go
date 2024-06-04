package main

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("input slice should not equal output set if input contains duplicates", func(t *testing.T) {
		inp := []string{"hello", "hello", "hi", "hi", "first", "first"}
		got := Set(inp)
		if reflect.DeepEqual(got, inp) {
			t.Errorf("output (%v) is equal to input (%v)", got, inp)
		}
	})
}
