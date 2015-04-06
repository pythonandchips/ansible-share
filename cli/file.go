package main

import (
	"os"
	"path/filepath"
)

type File struct {
	os.FileInfo
	path string
}

type Walker interface {
	ListFiles(string) []File
}

type FileWalker struct {
}

func (file FileWalker) ListFiles(dir string) []File {
	files := []File{}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		files = append(files, File{info, path})
		return nil
	})
	return files
}
