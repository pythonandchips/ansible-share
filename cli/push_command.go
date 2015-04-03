package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

type PushCommand struct {
	tag, path, pushUuid string
}

func NewPushCommand(tag string, path string) PushCommand {
	pushUuid := strings.Replace(uuid.New(), "-", "", -1)
	return PushCommand{tag: tag, path: path, pushUuid: pushUuid}
}

func (pushCommand *PushCommand) Execute() {
	files := pushCommand.getFiles()
	tarfile := pushCommand.createTarFile()
	defer tarfile.Close()
	pushCommand.AddFilesToTar(tarfile, files)
	pushCommand.UploadFile()
}

func (pushCommand *PushCommand) UploadFile() {
	tmpfile := filepath.Join("/tmp", "ansible_share", pushCommand.pushUuid)
	file, _ := os.Open(tmpfile)
	fileContent, _ := ioutil.ReadAll(file)
	file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fmt.Println(filepath.Base(tmpfile))
	part, errPart := writer.CreateFormFile("role", filepath.Base(tmpfile))
	if errPart != nil {
		panic(errPart)
	}
	part.Write(fileContent)
	url := "http://localhost:8080/roles/postgres/" + pushCommand.pushUuid
	fmt.Println(url)
	writer.Close()
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	_, reqErr := client.Do(req)
	if reqErr != nil {
		panic(reqErr)
	}
}

func (pushCommand *PushCommand) AddFilesToTar(tarfile *os.File, files []File) {
	var fileWriter io.WriteCloser = tarfile
	fileWriter = gzip.NewWriter(tarfile)
	defer fileWriter.Close()
	tarfileWriter := tar.NewWriter(fileWriter)
	defer tarfileWriter.Close()
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			continue
		}
		fmt.Println(fileInfo.path)
		file, _ := os.Open(fileInfo.path)
		filePath, _ := filepath.Rel(pushCommand.path, fileInfo.path)
		fmt.Println(filePath)
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
}

type File struct {
	os.FileInfo
	path string
}

func (pushCommand *PushCommand) createTarFile() *os.File {
	file, _ := os.Create(filepath.Join("/tmp", "ansible_share", pushCommand.pushUuid))
	return file
}

func (pushCommand *PushCommand) getFiles() []File {
	files := []File{}
	filepath.Walk(pushCommand.path, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		files = append(files, File{info, path})
		return nil
	})
	return files
}
