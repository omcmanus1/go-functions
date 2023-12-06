package main

import (
	"fmt"
	"os"
)

func GetFile() (*os.File, error) {
	file, err := os.Open("01-input.txt")
	if err != nil {
		fmt.Println("error reading file")
		return nil, err
	}
	return file, err
}
