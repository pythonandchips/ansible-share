package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"flag"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/codegangsta/cli"
)

func CreateTestFiles() {
	os.MkdirAll("./nginx/tasks", 0700)
	fileContent := `
---
name: test task
command: do something good
`
	ioutil.WriteFile("./nginx/tasks/main.yml", []byte(fileContent), 0700)
}

func DestroyTestFiles() {
	os.RemoveAll("./nginx")
}

func TestPushToServer(t *testing.T) {
	CreateTestFiles()
	defer DestroyTestFiles()
	var file multipart.File
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, _, _ = r.FormFile("role")
	}))
	defer ts.Close()
	url, _ := url.Parse(ts.URL)

	set := flag.NewFlagSet("test", 0)
	tag := url.Host + "/postgres:v1.1"
	path := "./nginx"
	set.Parse([]string{path})
	set.String("tag", tag, "doc")
	context := cli.NewContext(nil, set, set)

	Push(context)

	role, err := ioutil.ReadAll(file)
	if err != nil {
		t.Logf("File not found")
		t.FailNow()
	}
	if len(role) == 0 {
		t.Logf("File Empty")
		t.FailNow()
	}
	file.Close()
	fileReader := bytes.NewReader(role)
	gzipReader, gerror := gzip.NewReader(fileReader)
	if gerror != nil {
		t.Logf("GZIP ERROR")
		t.FailNow()
	}
	tarBallReader := tar.NewReader(gzipReader)
	names := []string{}
	for {
		header, err := tarBallReader.Next()
		if err == io.EOF {
			break
		}
		names = append(names, header.Name)
	}
	if len(names) != 1 {
		t.Fail()
	}
	if names[0] != "tasks/main.yml" {
		t.Log("name not correct " + names[0])
		t.Fail()
	}
}
