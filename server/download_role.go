package main

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func DownloadRoleHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	DownloadRole(response, request, vars)
}

func DownloadRole(response http.ResponseWriter, request *http.Request, vars map[string]string) {
	role := vars["role"]
	tag := vars["tag"]
	file := filepath.Join(directoryRoot, role, tag)
	http.ServeFile(response, request, file)
}
