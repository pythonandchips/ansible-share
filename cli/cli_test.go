package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/codegangsta/cli"
)

func TestPushToServer(t *testing.T) {
	defer os.Remove("/tmp/ansible_share/postgres/v1.1")
	set := flag.NewFlagSet("test", 0)
	tag := "localhost:8080/postgres:v1.1"
	path := "/Users/colingemmell/1partcarbon/capasa/ansible/roles/nginx"
	set.Parse([]string{path})
	set.String("tag", tag, "doc")
	context := cli.NewContext(nil, set, set)

	Push(context)

	if !exists("/tmp/ansible_share/postgres/v1.1") {
		t.Fail()
	}
	file, fileErr := os.Open("/tmp/ansible_share/postgres/v1.1")
	if fileErr != nil {
		t.Logf("file upload not found")
		t.FailNow()
	}
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
