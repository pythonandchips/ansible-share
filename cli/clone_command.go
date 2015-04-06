package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func NewCloneCommand(tag string) CloneCommand {
	return CloneCommand{tag: tag}
}

type CloneCommand struct {
	tag string
}

func (cloneCommand *CloneCommand) Execute() {
	cloneCommand.DownloadFile()
}

func (cloneCommand *CloneCommand) DownloadFile() {
	resp, getErr := http.Get("http://localhost:8080/roles/postgres/v1.1")
	if getErr != nil {
		fmt.Println(getErr)
	}
	role, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	defer resp.Body.Close()
	fileReader := bytes.NewReader(role)
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

		fmt.Println(header.Typeflag == tar.TypeReg)

		filename := "./role/postgres/" + header.Name
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
