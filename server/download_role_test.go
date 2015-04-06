package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	file, _ := os.Create("/tmp/ansible_share/postgres/testFile")
	file.Write([]byte("This is a file"))
	file.Close()

	responseRecorder := httptest.NewRecorder()
	url := "https://localhost:8080/roles/postgres/12345678"
	body := &bytes.Buffer{}
	req, _ := http.NewRequest("GET", url, body)

	vars := map[string]string{"role": "postgres", "tag": "testFile"}
	DownloadRole(responseRecorder, req, vars)

	if string(responseRecorder.Body.Bytes()) != "This is a file" {
		t.Log("file not returned")
		t.Fail()
	}

	os.Remove("/tmp/ansible_share/postgres/testFile")
}
