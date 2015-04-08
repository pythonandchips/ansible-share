package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type Transport interface {
	UploadFile([]byte, string, string)
	DownloadFile(string) []byte
}

type HttpTransport struct {
	url string
}

func (httpTransport HttpTransport) DownloadFile(url string) []byte {
	fmt.Println(url)
	resp, getErr := http.Get(url)
	if getErr != nil {
		fmt.Println(getErr)
	}
	role, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr)
	}
	defer resp.Body.Close()
	return role
}

func (httpTransport HttpTransport) UploadFile(file []byte, fieldName string, fileName string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, errPart := writer.CreateFormFile(fieldName, fieldName)
	if errPart != nil {
		panic(errPart)
	}
	part.Write(file)
	writer.Close()
	req, err := http.NewRequest("POST", httpTransport.url, body)
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
