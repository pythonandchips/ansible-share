package server

import (
	"fmt"
	"os"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func createDirectory(path string) {
	fileExist := exists(path)
	if !fileExist {
		err := os.MkdirAll(path, 0700)
		fmt.Println(err)
		if err != nil {
			panic("Cannot create directory")
		}
	}
}
