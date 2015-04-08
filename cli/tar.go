package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/1partcarbon/ansible_share/file"
)

type Compressor interface {
	Compress(string, []file.File) []byte
}

type Decompressor interface {
	Uncompress([]byte, string)
}

type Tar struct {
}

func (t Tar) Uncompress(file []byte, basePath string) {
	fileReader := bytes.NewReader(file)
	gzipReader, gzipErr := gzip.NewReader(fileReader)
	if gzipErr != nil {
		fmt.Println(gzipErr)
	}
	tarBallReader := tar.NewReader(gzipReader)
	for {
		header, err := tarBallReader.Next()
		if err == io.EOF {
			break
		}

		filename := basePath + "/" + header.Name
		switch header.Typeflag {
		case tar.TypeDir:
			os.MkdirAll(filename, os.FileMode(header.Mode))
		case tar.TypeReg:
			os.MkdirAll(filepath.Dir(filename), 0700)
			fmt.Println("uncompressing " + filename)
			writer, writeErr := os.Create(filename)
			if writeErr != nil {
				fmt.Println("Write error")
				fmt.Println(writeErr)
			}
			io.Copy(writer, tarBallReader)
			os.Chmod(filename, os.FileMode(header.Mode))
			writer.Close()
		default:
			fmt.Printf("Unable to untar type: %c in file %s", header.Typeflag, filename)
		}
	}
}

func (t Tar) Compress(root string, files []file.File) []byte {
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
		file, _ := os.Open(fileInfo.Path)
		filePath, _ := filepath.Rel(root, fileInfo.Path)
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
