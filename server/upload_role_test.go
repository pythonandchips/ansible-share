package server

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	responseRecorder := httptest.NewRecorder()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("role", "testFile.tar.gz")
	part.Write([]byte("This is a test"))
	writer.Close()
	url := "https://localhost:8080/roles/postgres/12345678"
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	vars := map[string]string{"role": "postgres", "tag": "12345678"}
	UploadRole(responseRecorder, req, vars)
	here := exists("/tmp/ansible_share/postgres/12345678")
	if !here {
		t.Log("posted File does not exists")
		t.Fail()
	}
	os.Remove("/tmp/ansible_share/postgres/12345678")
}
