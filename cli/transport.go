package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
)

type Transport interface {
	UploadFile([]byte, string, string)
}

type HttpTransport struct {
	url string
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
