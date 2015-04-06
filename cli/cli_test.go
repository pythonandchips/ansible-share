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
	"os"
	"testing"

	"github.com/codegangsta/cli"
)

func TestPushToServer(t *testing.T) {
	defer os.Remove("/tmp/ansible_share/postgres/v1.1")

	var file multipart.File
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, _, _ = r.FormFile("role")
	}))
	defer ts.Close()

	set := flag.NewFlagSet("test", 0)
	tag := ts.URL + "/postgres:v1.1"
	path := "/Users/colingemmell/1partcarbon/capasa/ansible/roles/nginx"
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
