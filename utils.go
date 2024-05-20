package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Primitives interface {
	bool | string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

// compare output with expected value
// return error if not equal
func CheckGotWantPrimitive[P Primitives](t testing.TB, got, want P) {
	t.Helper()
	if got != want {
		t.Errorf(`got %v want %v`, got, want)
	}
}

func CheckGotWantSlice[P Primitives](t testing.TB, got, want []P) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %v want %v`, got, want)
	}
}

// check for specific error message
// return error if no error found
func CheckError(t testing.TB, err error, expectedErrMsg string) {
	print(err)
	t.Helper()
	if err == nil {
		t.Errorf("got %v want nil", err)
	}
	assert.ErrorContains(t, err, expectedErrMsg)
}
