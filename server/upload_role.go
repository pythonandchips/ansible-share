package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func UploadRole(response http.ResponseWriter, request *http.Request, vars map[string]string) {
	file, header, err := request.FormFile("role")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	directory := filepath.Join(directoryRoot, vars["role"])
	createDirectory(directory)
	out, err := os.Create(filepath.Join(directory, vars["tag"]))
	if err != nil {
		fmt.Fprintf(response, "Unable to create the file for writing. Check your write access privilege")
	}
	defer out.Close()
	_, copyErr := io.Copy(out, file)
	if copyErr != nil {
		fmt.Fprintln(response, err)
	}
	fmt.Fprintf(response, "File uploaded successfully : ")
	fmt.Fprintf(response, header.Filename)
}

func UploadRoleHandler(response http.ResponseWriter, request *http.Request) {
	UploadRole(response, request, mux.Vars(request))
}
