package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/1partcarbon/ansible_share/file"
	"github.com/gorilla/mux"
)

func DownloadRoleHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	DownloadRole(response, request, vars)
}

func DownloadRole(response http.ResponseWriter, request *http.Request, vars map[string]string) {
	role := vars["role"]
	tag := vars["tag"]
	fmt.Println("Sending " + role + " at " + tag)
	if tag == "latest" {
		dirPath := filepath.Join(directoryRoot, role)
		fmt.Println(dirPath)
		fileWalker := file.FileWalker{}
		files := fileWalker.ListFiles(dirPath)
		latestFile := files[0]
		for _, file := range files {
			if latestFile.ModTime().Before(file.ModTime()) {
				latestFile = file
			}
		}
		fmt.Println(latestFile.Path)
		http.ServeFile(response, request, latestFile.Path)
	} else {
		file := filepath.Join(directoryRoot, role, tag)
		http.ServeFile(response, request, file)
	}
}
