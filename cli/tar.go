package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

type Compressor interface {
	Compress(string, []File) []byte
}

type Tar struct {
}

func (t Tar) Compress(root string, files []File) []byte {
	if len(files) == 0 {
		panic("I GOT NO FILES")
	}
	content := []byte{}
	tarfile := bytes.NewBuffer(content)
	fileWriter := gzip.NewWriter(tarfile)
	tarfileWriter := tar.NewWriter(fileWriter)
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}
		file, _ := os.Open(fileInfo.path)
		filePath, _ := filepath.Rel(root, fileInfo.path)
		defer file.Close()
		header := new(tar.Header)
		header.Name = filePath
		header.Size = fileInfo.Size()
		header.Mode = int64(fileInfo.Mode())
		header.ModTime = fileInfo.ModTime()
		header.Typeflag = tar.TypeReg

		tarfileWriter.WriteHeader(header)
		io.Copy(tarfileWriter, file)
	}
	tarfileWriter.Close()
	fileWriter.Close()
	return tarfile.Bytes()
}
