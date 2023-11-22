package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// compare output with expected value
// return error if not equal
func CheckGotWant(t testing.TB, got, want interface{}) {
	t.Helper()
	if got != want {
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
