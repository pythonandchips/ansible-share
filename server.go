package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

const directoryRoot = "/tmp/ansible_share"

func main() {
	createDirectory(directoryRoot)
	r := mux.NewRouter()
	r.HandleFunc("/roles/{role}/{tag}", UploadRole).Methods("POST")
	r.HandleFunc("/roles/{role}/{tag}", DownloadRole).Methods("GET")
	r.HandleFunc("/_ping", Ping).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}

func DownloadRole(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	role := vars["role"]
	tag := vars["tag"]
	file := filepath.Join(directoryRoot, role, tag)
	http.ServeFile(response, request, file)
}

func Ping(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "ping")
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createDirectory(path string) {
	fileExist, _ := exists(path)
	if !fileExist {
		err := os.MkdirAll(path, 0700)
		fmt.Println(err)
		if err != nil {
			panic("Cannot create directory")
		}
	}
}

func UploadRole(response http.ResponseWriter, request *http.Request) {
	fmt.Println("RECEIVED UPDATE")
	file, header, err := request.FormFile("role")
	vars := mux.Vars(request)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}
	defer file.Close()
	directory := filepath.Join(directoryRoot, vars["role"])
	createDirectory(directory)
	fmt.Println(filepath.Join(directory, vars["tag"]))
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
